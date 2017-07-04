// Package main provides ...
package main

import (
	"log"

	"golang.org/x/exp/inotify"
)

func main() {

	watcher, err := inotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	err = watcher.Watch("/Users/Jialin/tmp/go_test.sh")
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case ev := <-watcher.Event:
			log.Println("event:", ev)
		case err := <-watcher.Error:
			log.Println("error:", err)
		}
	}
}
