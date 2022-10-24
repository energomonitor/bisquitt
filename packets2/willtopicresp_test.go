package packets2

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	pkts "github.com/energomonitor/bisquitt/packets"
)

func TestWillTopicRespConstructor(t *testing.T) {
	assert := assert.New(t)

	pkt := NewWillTopicResp(RC_CONGESTION)

	if pkt == nil {
		t.Fatal("New packet should not be nil")
	}

	assert.Equal("*packets2.WillTopicResp", reflect.TypeOf(pkt).String(), "Type should be WillTopicResp")
	assert.Equal(RC_CONGESTION, pkt.ReasonCode, "ReasonCode should be RC_CONGESTION")
}

func TestWillTopicRespMarshal(t *testing.T) {
	pkt1 := NewWillTopicResp(RC_CONGESTION)
	pkt2 := testPacketMarshal(t, pkt1)
	assert.Equal(t, pkt1, pkt2.(*WillTopicResp))
}

func TestWillTopicRespUnmarshalInvalid(t *testing.T) {
	assert := assert.New(t)

	// Packet too short.
	buff := bytes.NewBuffer([]byte{
		2,                        // Length
		byte(pkts.WILLTOPICRESP), // Packet Type
		// Return Code missing
	})
	_, err := ReadPacket(buff)
	if assert.Error(err) {
		assert.Contains(err.Error(), "bad WILLTOPICRESP2 packet length")
	}

	// Packet too long.
	buff = bytes.NewBuffer([]byte{
		4,                        // Length
		byte(pkts.WILLTOPICRESP), // Packet Type
		0,                        // Reason Code
		0,                        // junk
	})
	_, err = ReadPacket(buff)
	if assert.Error(err) {
		assert.Contains(err.Error(), "bad WILLTOPICRESP2 packet length")
	}
}

func TestWillTopicRespStringer(t *testing.T) {
	pkt := NewWillTopicResp(RC_CONGESTION)
	assert.Equal(t, "WILLTOPICRESP2(ReasonCode=congestion)", pkt.String())
}
