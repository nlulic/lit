package main

import (
	"flag"
	"os"

	"gitlab.com/nlulic/lit/logger"
)

func main() {

	lit := NewLit(
		logger.New(logger.LevelDebug),
	)

	if len(os.Args) < 2 {
		os.Exit(1)
	}

	flag.Parse()

	switch os.Args[1] {
	case "init":
		lit.Init()
	case "commit":
		lit.Commit(flag.Arg(1))
	case "branch":
		lit.Branch(flag.Arg(1))
	case "checkout":
		lit.Checkout(flag.Arg(1))
	case "log":
		lit.Log()
	default:
		os.Exit(1)
	}

}
