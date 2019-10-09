package main

import (
	"fmt"
	"github.com/ilijamt/vht/internal/cmd"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)
	//log.SetLevel(log.DebugLevel)
}

func main() {
	var err error
	if err = cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
