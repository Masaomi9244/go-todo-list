package domain

type Status int

const (
	NotStarted Status = iota
	InProgress
	Done
)

func (s Status) String() string {
	switch s {
	case NotStarted:
		return "未着手"
	case InProgress:
		return "進行中"
	case Done:
		return "完了"
	default:
		return "不明"
	}
}
