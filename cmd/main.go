package main

import (
	"encoding/json"
	"github.com/google/logger"
	"os"
	"streamtodb/domain/entity"
	"streamtodb/domain/service"
	"streamtodb/infra/input"
	"streamtodb/infra/persistence"
	"streamtodb/interfaces"
	"streamtodb/pkg/shutdown"
)

func main() {
	// init logger and config
	logger.Init("streamtodb", false, false, os.Stdout)
	InitConfig()

	if len(os.Args) <= 1 {
		logger.Fatalln("provide data file")
	}
	filePath := os.Args[1]

	handler := interfaces.NewPortHandler(service.NewPortService(persistence.NewPortRepository(persistence.ConnectDb())))
	// producer parses the file and sends to handler
	producer := input.NewProducer()
	go producer.Start(dataStream(filePath))
	connect(producer, handler)
	// gracefully shutdown
	shutdown.Graceful()
}

func dataStream(filePath string) *os.File {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	return f
}

func connect(producer *input.Producer, handler *interfaces.PortHandler) {
	// handle incoming messages from stream
	for item := range producer.Output() {
		// assert interface of parsed item
		item, ok := item.(input.Item)
		if !ok {
			logger.Info("unknown type")
			continue
		}
		var port entity.Port
		if err := json.Unmarshal(item.Body, &port); err != nil {
			logger.Error(err)
			continue
		}
		port.Codename = item.Key
		if err := handler.SavePort(port); err != nil {
			logger.Error(err)
			continue
		}
	}
}
