package main

import (
	"flag"
	"os"

	"gitlab.com/nlulic/lit/logger"
)

func main() {

	commitCmd := flag.NewFlagSet("commit", flag.ExitOnError)
	commitMsg := commitCmd.String("m", "", "commit message")

	branchCmd := flag.NewFlagSet("branch", flag.ExitOnError)

	checkoutCmd := flag.NewFlagSet("checkout", flag.ExitOnError)

	lit := NewLit(
		logger.New(logger.LevelInfo),
	)

	if len(os.Args) < 2 {
		lit.logger.Fatal("fatal: subcommand expected")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "init":
		lit.Init()
	case "commit":
		commitCmd.Parse(os.Args[2:])
		lit.Commit(*commitMsg)
	case "branch":
		branchCmd.Parse(os.Args[2:])
		lit.Branch(branchCmd.Arg(0))
	case "checkout":
		checkoutCmd.Parse(os.Args[2:])
		lit.Checkout(checkoutCmd.Arg(0))
	case "log":
		lit.Log()
	default:
		os.Exit(1)
	}
}
