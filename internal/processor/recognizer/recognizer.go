package recognizer

type Result struct {
	Text string
	IsFinal bool
	Confidance float32
}

type Recognizer interface {
	Recognize([]byte) (*Result, error)
}

type recognizer struct {}

func (r *recognizer) Recognize([]byte) (*Result, error) {
	return nil, nil
}