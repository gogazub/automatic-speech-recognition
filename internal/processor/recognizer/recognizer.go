package recognizer

import (
	"context"
	"math/rand"
	"time"
)

type Result struct {
	Text string
	IsFinal bool
	Confidance float32
}

type Recognizer interface {
	Recognize(ctx context.Context, pcmData []byte) (*Result, error)
}

type MockRecognizer struct {
	buffer     []byte
	sampleRate int32
	words      []string
}

func NewRecognizer(sampleRate int32) *MockRecognizer {
	return &MockRecognizer{
		buffer:     make([]byte, 0, 32000),
		sampleRate: sampleRate,
		words:      []string{"здравствуйте", "я", "хочу", "оформить", "карту", "тинькофф", "спасибо", "до", "свидания"},
	}
}

func (r *MockRecognizer) Recognize(ctx context.Context, pcmData []byte) (*Result, error) {
	r.buffer = append(r.buffer, pcmData...)

	// Если samplerate = x Гц, то x * bitdepth * 1 сек = количество байт в секунде записи. Пусть будем обрабатывать чанками по 0.5сек записи.
	threshold := int(r.sampleRate * 2 / 2) // bitdepth == 2

	if len(r.buffer) >= threshold {
		
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(50*time.Millisecond) :
		}

		word := r.words[rand.Intn(len(r.words))]
		IsFinal := rand.Float32() > 0.7

		r.buffer = r.buffer[:0]

		return &Result{
			Text: word,
			IsFinal: IsFinal,
			Confidance: 0.9+rand.Float32()/10,
		}, nil
	}		
	return nil, nil
}
