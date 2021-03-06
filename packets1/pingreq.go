package packets1

import (
	"fmt"
	"io"
)

type Pingreq struct {
	Header
	ClientID []byte
}

// NOTE: Packet length is initialized in this constructor and recomputed in m.Write().
func NewPingreq(clientID []byte) *Pingreq {
	p := &Pingreq{
		Header:   *NewHeader(PINGREQ, 0),
		ClientID: clientID,
	}
	p.computeLength()
	return p
}

func (p *Pingreq) computeLength() {
	length := len(p.ClientID)
	p.Header.SetVarPartLength(uint16(length))
}

func (p *Pingreq) Write(w io.Writer) error {
	p.computeLength()

	buf := p.Header.pack()
	if len(p.ClientID) > 0 {
		buf.Write(p.ClientID)
	}

	_, err := buf.WriteTo(w)
	return err
}

func (p *Pingreq) Unpack(r io.Reader) (err error) {
	if p.VarPartLength() > 0 {
		p.ClientID = make([]byte, p.VarPartLength())
		_, err = io.ReadFull(r, p.ClientID)
	} else {
		p.ClientID = nil
	}
	return
}

func (p Pingreq) String() string {
	return fmt.Sprintf("PINGREQ(ClientID=%#v)", string(p.ClientID))
}
