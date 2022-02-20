package main

import "os"

func main() {

	lit := NewLit()

	// TODO:
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	switch os.Args[1] {
	case "init":
		lit.Init()
	default:
		os.Exit(1)
	}

}
