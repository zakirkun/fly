package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

type FlyParams struct {
	Path  string
	Build Build
}

type FlyEvent interface {
	Watch()
}

func NewEvent(opt FlyParams) FlyEvent {
	return &FlyParams{
		Path:  opt.Path,
		Build: opt.Build,
	}
}

// Watch implements FlyEvent.
func (f *FlyParams) Watch() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Goroutine For watcher
	done := make(chan bool)

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("EVENT DETECTED: ", event)
				if event.Has(fsnotify.Write) {
					log.Println("Modified Files:", event.Name)
				} else if event.Has(fsnotify.Create) {
					log.Println("Create Files:", event.Name)
				} else if event.Has(fsnotify.Remove) {
					log.Println("Delete Files:", event.Name)
				} else if event.Has(fsnotify.Rename) {
					log.Println("Rename Files:", event.Name)
				} else {
					log.Println("Event Init")
				}

				// Build Temp Project
				f.Build.Build()

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Fatal("error:", err)
				return
			}
		}
	}()

	err = watcher.Add(f.Path)
	if err != nil {
		log.Fatal(err)
	}

	// Stop
	<-done
}
