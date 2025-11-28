package processor

type Config struct {
	Encoding string
	SampleRate int32
	AudioChannelCount int32
	Language string
}

func NewProcess(cfg Config) Recognizer {
	return nil
}