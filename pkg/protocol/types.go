package protocol

type data struct {
	Type DataType `json:"type"`
	Data []byte   `json:"data"`
}

func (d *data) Marshal() ([]byte, error) {
	return json.Marshal(d)
}

func (d *data) Unmarshal(data []byte) error {
	return json.Unmarshal(data, d)
}

type ErrCode int

const (
	// TODO
	ErrUnknown ErrCode = iota
)

type errs struct {
	Id   int64   `json:"id"`
	Code ErrCode `json:"error_type"`
}

func (e *errs) Marshal() ([]byte, error) {
	return json.Marshal(e)
}

func (e *errs) Unmarshal(data []byte) error {
	return json.Unmarshal(data, e)
}

type msg struct {
	Id   int64  `json:"id"`
	From int64  `json:"from"`
	To   int64  `json:"to"`
	Text string `json:"text"`
}

func (m *msg) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

func (m *msg) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}
