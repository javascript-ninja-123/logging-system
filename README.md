


### Instruction

	logger := logger2.NewLogger()
	logger.Hook = func(level logger2.Level, msg string) {
        //works like JS callback
		fmt.Println(level.String())
		fmt.Println(msg)
	}

	logger.WithField("logging from my S3 bucket", "file upload complete").Info("done")