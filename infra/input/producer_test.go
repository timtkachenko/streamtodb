package input

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPanicInvalidJson(t *testing.T) {
	a := assert.New(t)
	p := NewProducer()
	buf := bytes.NewBuffer([]byte(`"test":{}"`))
	a.Panics(func() {
		p.parse(buf)
	})
}
func TestParsed(t *testing.T) {
	a := assert.New(t)
	p := NewProducer()
	buf := bytes.NewBuffer([]byte(`{"test":{}}`))
	go p.parse(buf)
	expected := Item{
		Body: []byte("{}"),
		Key:  "test",
	}
	a.Equal(expected, <-p.Output())
}
