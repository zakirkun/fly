package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

var targetFolder *string

func init() {
	targetFolder = flag.String("target", ".", "Target folder to watch for changes (default: current directory)")
	flag.Parse()
}

func main() {

	// Resolve the full path of the target folder
	fullPath, err := filepath.Abs(*targetFolder)
	if err != nil {
		log.Fatal("Error resolving target folder:", err)
	}

	// Setup Project Builder
	build := Build{
		TempApp:    "app" + randomString(12),
		TempOutDir: "temp_" + randomString(5),
		Path:       fullPath,
	}

	// Setup Flying Watcher
	params := FlyParams{
		Path:  build.Path,
		Build: build,
	}

	flying := NewEvent(params)
	go flying.Watch()

	// Signal
	sigChan := make(chan os.Signal, 1)
	// Listen for interrupt signals (Ctrl+C, SIGINT, SIGTERM)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal is received
	go func() {
		sig := <-sigChan
		fmt.Println("\nReceived signal:", sig)
		build.Cleanup()
		os.Exit(0)
	}()

	// Block main goroutine forever.
	<-make(chan struct{})
}
