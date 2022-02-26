package main

import (
	"os"

	"gitlab.com/nlulic/lit/logger"
)

func main() {

	lit := NewLit(
		logger.New(logger.LevelDebug),
	)

	// TODO:
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	switch os.Args[1] {
	case "init":
		lit.Init()
	case "commit":
		// TODO: commit message should come from args
		lit.Commit("default message")
	default:
		os.Exit(1)
	}

}
