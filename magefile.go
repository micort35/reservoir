// +build mage

package main

import (
	"log"

	"github.com/magefile/mage/sh"
)

// Builds the binary for the Reservoir CLI application.
func BuildCli() {
	if err := sh.Run("go", "build", "-o", "./build/reservoir_cli.exe", "./cmd/reservoir_cli.go"); err != nil {
		log.Fatal(err)
	}
}

// Removes build artifacts.
func Clean() {
	if err := sh.Rm("./build"); err != nil {
		log.Fatal(err)
	}
}
