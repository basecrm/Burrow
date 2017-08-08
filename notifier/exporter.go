package notifier

type Exporter interface {
	Export(msg Message) error
}
