package bridge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorNotification_Notify(t *testing.T) {

	errNotification := &ErrorNotification{sender: NewWechatMsgSender([]string{"group1"})}

	err := errNotification.Notify("error")
	assert.Nil(t, err)

	warnNotification := &WarningNotification{sender: NewEmailMsgSender([]string{"email1"})}
	err = warnNotification.Notify("warning")
	assert.Nil(t, err)

}
