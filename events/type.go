package events

type Fetcher interface {
	Fetch(limit int, offset int) ([]Events, error)
}

type Processor interface {
	Process(e Events) error
}

type Type int

const (
	Unknown Type = iota
	Message
)

type Events struct {
	Type Type
	Text string
}
