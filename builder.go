package main

import (
	"log"
	"os/exec"
)

var CMD *exec.Cmd

type Build struct {
	TempApp    string
	TempOutDir string
	Path       string
}

func (b *Build) checkDIr() error {
	path := b.TempOutDir
	app := b.TempApp
	if err := ensureFolderExists(path); err != nil {
		return err
	}
	if err := ensureFileExists(path + "/" + app); err != nil {
		return err
	}
	return nil
}

func (b *Build) Build() {
	// Ensure Dir
	err := b.checkDIr()
	if err != nil {
		log.Fatalf("ERROR CREATE TEMP DIR: %v", err.Error())
	}

	// Build the project
	CMD = exec.Command("go", "build", "-o", b.TempOutDir+"/"+b.TempApp, b.Path)
	CMD.Stdout = log.Writer()
	CMD.Stderr = log.Writer()

	if err := CMD.Run(); err != nil {
		log.Fatal(err)
	}

	// Start App
	b.Reload()
}

func (b *Build) Reload() {
	// Ensure Dir
	err := b.checkDIr()
	if err != nil {
		log.Fatalf("ERROR CREATE TEMP DIR: %v", err.Error())
	}

	// Rebuild Project
	b.Build()

	// Kill Previous Process if running
	b.Cleanup()

	// Run The Project
	CMD = exec.Command("./" + b.TempOutDir + "/" + b.TempApp)
	CMD.Stdout = log.Writer()
	CMD.Stderr = log.Writer()
	if err := CMD.Run(); err != nil {
		log.Fatal(err)
	}
}

func (b *Build) Cleanup() {
	if CMD != nil && CMD.Process != nil {
		CMD.Process.Kill()
	}
}
