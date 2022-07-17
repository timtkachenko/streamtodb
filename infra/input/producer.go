package input

import (
	"bufio"
	"encoding/json"
	"github.com/google/logger"
	"io"
)

type Item struct {
	Body json.RawMessage
	Key  string
}

// Producer creates stream of Items from json object
type Producer struct {
	output chan interface{}
}

func NewProducer() *Producer {
	ch := make(chan interface{})
	return &Producer{ch}
}

func (p Producer) Output() <-chan interface{} {
	return p.output
}
func (p Producer) Start(rd io.Reader) {
	jsonStream := bufio.NewReader(rd)
	p.parse(jsonStream)
}

func (p Producer) parse(jsonStream io.Reader) {
	dec := json.NewDecoder(jsonStream)
	for {
		if err := p.decode(dec); err == io.EOF {
			break
		}
		if !dec.More() {
			logger.Info("parsing finished")
		}
	}
	close(p.output)
}

// convert buffered json chunk into Item
func (p Producer) decode(dec *json.Decoder) error {
	t, err := dec.Token()
	if err == io.EOF {
		return io.EOF
	}
	if err != nil {
		panic(err)
	}
	if _, ok := t.(json.Delim); !ok {
		var entity Item
		err := dec.Decode(&entity.Body)
		if err != nil {
			panic(err)
		}
		entity.Key = t.(string)
		p.output <- entity
	}
	return nil
}
