// contract
package processor

import "context"

type Result struct {
	Text string
	IsFinal bool
	Confidance float32
}

type Recognizer interface {
	Process(ctx context.Context, chunk []byte) (*Result, error)
}
