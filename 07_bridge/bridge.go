package bridge

import "fmt"

// IMsgSender
type IMsgSender interface {
	Send(msg string) error
}

type EmailMsgSender struct {
	emails []string
}

type WechatMsgSender struct {
	groups []string
}

func NewEmailMsgSender(emails []string) *EmailMsgSender {
	return &EmailMsgSender{emails: emails}
}

func NewWechatMsgSender(groups []string) *WechatMsgSender {
	return &WechatMsgSender{groups: groups}
}

func (s *EmailMsgSender) Send(msg string) error {
	fmt.Printf("EmailMsgSender send: %s\n", msg)
	return nil
}

func (s *WechatMsgSender) Send(msg string) error {
	fmt.Printf("WechatMsgSender send: %s\n", msg)
	return nil
}

type INotification interface {
	Notify(msg string) error
}

type ErrorNotification struct {
	sender IMsgSender
}

type WarningNotification struct {
	sender IMsgSender
}

func (n *ErrorNotification) Notify(msg string) error {
	return n.sender.Send(msg)
}

func (n *WarningNotification) Notify(msg string) error {
	return n.sender.Send(msg)
}
