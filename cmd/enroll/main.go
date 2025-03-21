package main

import (
	"flag"
	"github.com/alhaos/enroll/internal/config"
	"github.com/alhaos/enroll/internal/handlers"
	"github.com/alhaos/enroll/internal/logging"
	"github.com/alhaos/enroll/internal/webServer"
)

func main() {

	configFilename := parserConfigFilename()

	// Init config
	conf, err := config.MustLoadFromFile(configFilename)
	if err != nil {
		panic(err)
	}

	// Init logger
	logger := logging.New(conf.Logging)

	// Init handlers
	handler := handlers.New(logger)

	// Init web server
	ws, err := webServer.New(conf.WebServer, handler)
	if err != nil {
		panic(err)
	}

	// Run web server
	err = ws.Run()
	if err != nil {
		panic(err)
	}
}

// parserConfigFilename take config filename from arguments
func parserConfigFilename() string {
	filenamePointer := flag.String("config", "config/config.yml", "enroll config filename")
	flag.Parse()
	return *filenamePointer
}
