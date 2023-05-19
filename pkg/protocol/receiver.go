package protocol

type Receiver interface {
	ReceiveACK(id int64)
	ReceiveError(id int64, errCode ErrCode)
	ReceiveText(id, from, to int64, text string, needACK bool)
}
