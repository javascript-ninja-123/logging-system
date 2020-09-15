package main

import (
	"fmt"

	"github.com/tjdalsvndn9/logger2"
)

func main() {

	logger := logger2.NewLogger()
	logger.Hook = func(level logger2.Level, msg string) {
		fmt.Println(level.String())
		fmt.Println("call back" + msg)
	}

	logger.WithField("logging from our S3 bucket", "completed").Info("done")

}
