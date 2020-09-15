package main

import (
	"fmt"

	"github.com/tjdalsvndn9/logger2"
)

func main() {

	logger := logger2.NewLogger()
	logger.Hook = func(level logger2.Level, msg string) {
		fmt.Println(level.String())
		fmt.Println(msg)
	}

	logger.WithField("it sent out data to s3", "completed").Info("dsagads")

}
