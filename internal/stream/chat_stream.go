package stream

import (
	"encoding/json"
	"io"
	"cursor-api/internal/domain/model"
)

type ChatStream struct {
	reader io.Reader
}

func NewChatStream(reader io.Reader) *ChatStream {
	return &ChatStream{
		reader: reader,
	}
}

func (s *ChatStream) Read() (*model.ChatResponse, error) {
	decoder := json.NewDecoder(s.reader)
	var response model.ChatResponse
	if err := decoder.Decode(&response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (s *ChatStream) Process(ch chan<- model.ChatResponse) error {
	defer close(ch)

	for {
		response, err := s.Read()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		ch <- *response
	}
} 