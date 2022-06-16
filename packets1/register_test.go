package packets1

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterStruct(t *testing.T) {
	topicID := uint16(123)
	topic := "test-topic"
	msg := NewRegister(topicID, topic)

	if assert.NotNil(t, msg, "New packet should not be nil") {
		assert.Equal(t, "*packets1.Register", reflect.TypeOf(msg).String(), "Type should be Register")
		assert.Equal(t, topicID, msg.TopicID, "Bad TopicID value")
		assert.Equal(t, uint16(0), msg.MessageID(), "Default MessageID should be 0")
		assert.Equal(t, topic, msg.TopicName, "Bad TopicName value")
	}
}

func TestRegisterMarshal(t *testing.T) {
	assert := assert.New(t)
	buf := bytes.NewBuffer(nil)

	msg1 := NewRegister(123, "test-topic")
	msg1.SetMessageID(12)
	if err := msg1.Write(buf); err != nil {
		t.Fatal(err)
	}

	r := bytes.NewReader(buf.Bytes())
	msg2, err := ReadPacket(r)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(msg1, msg2.(*Register))
}
