package packets1

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWillMsgRespStruct(t *testing.T) {
	msg := NewWillMsgResp(RC_ACCEPTED)

	if assert.NotNil(t, msg, "New packet should not be nil") {
		assert.Equal(t, "*packets1.WillMsgResp", reflect.TypeOf(msg).String(), "Type should be WillMsgResp")
		assert.Equal(t, RC_ACCEPTED, msg.ReturnCode, "Default ReturnCode should be RC_ACCEPTED")
	}
}

func TestWillMsgRespMarshal(t *testing.T) {
	assert := assert.New(t)
	buf := bytes.NewBuffer(nil)

	msg1 := NewWillMsgResp(RC_CONGESTION)
	if err := msg1.Write(buf); err != nil {
		t.Fatal(err)
	}

	r := bytes.NewReader(buf.Bytes())
	msg2, err := ReadPacket(r)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(msg1, msg2.(*WillMsgResp))
}