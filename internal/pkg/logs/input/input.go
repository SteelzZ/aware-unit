package input

type (
	Reader interface {
		Read() error
	}

	ReaderFunc func() error
)
