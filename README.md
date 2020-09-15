


### Instruction

	logger := logger2.NewLogger()
	logger.Hook = func(level logger2.Level, msg string) {
        //works like JS callback
		fmt.Println(level.String())
		fmt.Println(msg)
	}

	logger.WithField("logging from my S3 bucket", "file upload complete").Info("done")


### Result
{"level":"info","logging from our S3 bucket":"completed","msg":"done","time":"2020-09-15 11:44:34.590928 -0700 PDT m=+0.343550616"}info
call backdone