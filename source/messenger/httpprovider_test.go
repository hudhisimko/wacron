package messenger

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTTPProvider_buildRequest(t *testing.T) {
	type fields struct {
		URL         string
		Method      string
		ContentType string
		MessageTmpl string
		cli         *http.Client
	}
	type args struct {
		msg Message
		add map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		assert  func(h *HTTPProvider, gotReq *http.Request, gotErr error, t *testing.T)
		wantErr bool
	}{
		{
			name: "#1 parse template for JSON",
			fields: fields{
				MessageTmpl: `{"app":"appOrSomethings", "key": "KeyOrSecretData", "to": "{{ .To }}", "message": "{{ .Message }}", "from": "{{.From}}" }`,
				ContentType: "application/json",
			},
			args: args{
				msg: Message{
					From: "foo",
					To:   "bar",
					Body: "Hello World!",
				},
				add: map[string]string{
					"sender": "baz",
				},
			},
			wantErr: false,
			assert: func(h *HTTPProvider, gotReq *http.Request, gotErr error, t *testing.T) {
				b, err := io.ReadAll(gotReq.Body)
				if err != nil {
					t.Error(err)
				}

				assert.Equal(t, `{"app":"appOrSomethings","from":"foo","key":"KeyOrSecretData","message":"Hello World!","sender":"baz","to":"bar"}`, strings.TrimSpace(string(b)))
			},
		},
		{
			name: "#2 parse template for FORM",
			fields: fields{
				MessageTmpl: `{"to": "{{ .To }}"}`,
				ContentType: "application/x-www-form-urlencoded",
			},
			args: args{
				msg: Message{
					From: "foo",
					To:   "bar",
					Body: "Hello World!",
				},
			},
			wantErr: false,
			assert: func(h *HTTPProvider, gotReq *http.Request, gotErr error, t *testing.T) {
				b, err := io.ReadAll(gotReq.Body)
				if err != nil {
					t.Error(err)
				}

				assert.Equal(t, `to=bar`, strings.TrimSpace(string(b)))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HTTPProvider{
				URL:         tt.fields.URL,
				Method:      tt.fields.Method,
				ContentType: tt.fields.ContentType,
				MessageTmpl: tt.fields.MessageTmpl,
				cli:         tt.fields.cli,
			}
			got, err := h.buildRequest(tt.args.msg, tt.args.add)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPProvider.buildRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tt.assert(h, got, err, t)
		})
	}
}
