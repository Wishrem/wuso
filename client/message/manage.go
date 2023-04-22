package message

var m *manage

type manage struct {
}

func init() {
	m = &manage{}
}

func (m *manage) SendMsg(msg ...Message) error {
	return nil
}

func (m *manage) RecvMsg(bytes []byte) {

}
