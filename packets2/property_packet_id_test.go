package packets2

// TODO: Uncomment as soon as Pubrec and Pubrel types are defined.
/*
func TestPacketID(t *testing.T) {
	msgID := uint16(1234)

	pkt1 := NewPubrec()
	pkt1.SetPacketID(msgID)
	assert.Equal(t, msgID, pkt1.PacketID())

	pkt2 := NewPubrel()
	pkt2.CopyPacketID(pkt1)
	assert.Equal(t, msgID, pkt2.PacketID())
}
*/
