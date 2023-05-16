package types

import jsoniter "github.com/json-iterator/go"

var json = jsoniter.ConfigFastest

type Message struct {
	Id        int64  `json:"id"`
	IsGroup   bool   `json:"is_group`
	From      int64  `json:"from"`
	To        int64  `json:"to"`
	Index     int64  `json:"idx"`
	Content   string `json:"content"`
	EOF       bool   `json:"eof"`
	CreatedAt int64  `json:"time"`
}

func (m *Message) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

func (m *Message) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}
