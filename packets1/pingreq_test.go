package packets1

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingreqStruct(t *testing.T) {
	clientID := []byte("test-client")
	msg := NewPingreq(clientID)

	if assert.NotNil(t, msg, "New packet should not be nil") {
		assert.Equal(t, "*packets1.Pingreq", reflect.TypeOf(msg).String(), "Type should be Pingreq")
		assert.Equal(t, clientID, msg.ClientID, "Bad ClientID value")
	}
}

func TestPingreqMarshal(t *testing.T) {
	assert := assert.New(t)
	buf := bytes.NewBuffer(nil)

	msg1 := NewPingreq([]byte("test-client"))
	if err := msg1.Write(buf); err != nil {
		t.Fatal(err)
	}

	r := bytes.NewReader(buf.Bytes())
	msg2, err := ReadPacket(r)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(msg1, msg2.(*Pingreq))
}