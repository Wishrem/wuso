package screen

type Code int

const (
	NoticeAttention = iota
	NoticeWrong
)

func Notice(title string, msg string, code Code) {
}
