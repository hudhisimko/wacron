package messenger

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/rs/zerolog/log"
)

type HTTPProvider struct {
	URL         string
	Method      string
	ContentType string
	// MessageTmpl is template that will be parsed and sent as request body to URL.
	// Accessible data is as defined in msgTmplData.
	MessageTmpl string
	Headers     map[string]string

	cli *http.Client
}

// NewHTTPProvider creates a new instance of HTTPProvider.
func NewHTTPProvider() *HTTPProvider {
	return &HTTPProvider{
		cli: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (h *HTTPProvider) buildRequest(msg Message, additionalData map[string]string) (*http.Request, error) {
	// build headers
	header := http.Header{}
	for k, v := range h.Headers {
		header.Add(k, v)
	}

	tmplData := msgTmplData{
		To:      msg.To,
		ToRaw:   msg.To,
		From:    msg.From,
		Message: strings.TrimSuffix(strings.TrimPrefix(strconv.Quote(msg.Body), `"`), `"`),
		// Why did I do that?
		// Because the h.MessageTmpl is raw json.
		// Lets say the template looks like
		// 	```
		// 	{"chat_id":"783030276","text":"{{- .Message -}}"}
		//	```
		//
		// and the message looks like
		// 	```
		//	Welcome\nHello World
		//	```
		//
		// The executed template will become
		//
		// {"chat_id":"783030276","text":"Welcome\nHello World"}
		//
		// which will error if we try to marshal to json (invalid character '\\n' in string literal)
		// like syntax error things, which is true.
		//
		// But if we modify to the json to be
		//	..."text":"Welcome\\nHello World"...
		// It works, because \\n means we want to have newline inside 'text' key.
	}

	if strings.HasPrefix(tmplData.To, "0") {
		tmplData.To = strings.TrimPrefix("0", "62")
	}

	tmpl, err := template.New("tmpl").Parse(h.MessageTmpl)
	if err != nil {
		return nil, fmt.Errorf("buildRequest tmpl: %w", err)
	}

	var jsonTmplRes bytes.Buffer
	if err := tmpl.Execute(&jsonTmplRes, tmplData); err != nil {
		return nil, fmt.Errorf("buildRequest tmplX: %w", err)
	}

	var rawJSON map[string]string
	if err := json.Unmarshal(jsonTmplRes.Bytes(), &rawJSON); err != nil {
		return nil, fmt.Errorf("buildRequest json: %w", err)
	}

	// append additionalData
	for k, v := range additionalData {
		rawJSON[k] = v
	}

	var reqBody bytes.Buffer

	switch h.ContentType {
	default:
		return nil, errors.New("httpprovider: unknown content-type")

	case "application/x-www-form-urlencoded":
		form := url.Values{}
		for k, v := range rawJSON {
			form.Add(k, v)
		}

		if _, err := reqBody.WriteString(form.Encode()); err != nil {
			return nil, err
		}

	case "application/json":
		if err := json.NewEncoder(&reqBody).Encode(rawJSON); err != nil {
			return nil, fmt.Errorf("buildRequest: %w", err)
		}
	}

	req, err := http.NewRequest(h.Method, h.URL, &reqBody)
	if err != nil {
		return nil, fmt.Errorf("buildRequest newReq: %w", err)
	}

	req.Header = header
	req.Header.Set("Content-Type", h.ContentType)

	log.Debug().Caller().
		Str("tmpl", jsonTmplRes.String()).
		Str("content-type", h.ContentType).
		Str("method", h.Method).
		Any("body_json", rawJSON).
		Send()

	return req, nil
}

// SendMessage implements messenger.MessagingProvider.
func (h *HTTPProvider) SendMessage(msg Message, additionalData map[string]string) error {
	req, err := h.buildRequest(msg, additionalData)
	if err != nil {
		return err
	}

	// Send the request
	resp, err := h.cli.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	resBody, _ := io.ReadAll(resp.Body)
	log.Debug().Caller().
		Int("status", resp.StatusCode).
		Bytes("res", resBody).
		Send()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		errB := resBody
		if len(errB) > 255 {
			errB = errB[:254]
		}
		return fmt.Errorf("unexpected status code: %d. res: %v", resp.StatusCode, string(errB))
	}

	return nil
}
