package main

import (
	"fmt"
	cmd "github.com/ilijamt/vht/cmd/vht/cmd"
	"os"
)

func main() {
	var err error
	if err = cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
