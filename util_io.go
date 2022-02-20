package main

import (
	"log"
	"os"
)

func mustMakeDirs(paths ...string) {
	for _, path := range paths {
		if err := os.MkdirAll(path, 0644); err != nil {
			log.Fatal(err)
		}
	}
}
