package messenger

type Message struct {
	From string
	To   string
	Body string
}

type MessagingProvider interface {
	SendMessage(msg Message) error
}

type NoOpMsgProvider struct {
}

// SendMessage implements MessagingProvider.
func (n *NoOpMsgProvider) SendMessage(msg Message) error {
	return nil
}

var _ MessagingProvider = (*NoOpMsgProvider)(nil)
