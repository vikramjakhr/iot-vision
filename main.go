package main

import (
	_ "gitlab.intelligrape.net/tothenew/vision/routers"
	"github.com/astaxie/beego"
	"github.com/howeyc/fsnotify"
	"log"
	"os"
	"os/signal"
	"gitlab.intelligrape.net/tothenew/vision/services"
)

func main() {
	watcher()
	beego.Run()
}

func watcher() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	// Process events
	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				log.Println("Event received : ", ev)
				go func() {
					if ev.IsCreate() {
						services.DetectText(ev.Name)
					}
				}()
			case err := <-watcher.Error:
				log.Println("Error while listening to event ", err)
			}
		}
	}()

	err = watcher.Watch("/home/vikram/Desktop/img")
	if err != nil {
		log.Fatal(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		watcher.Close()
		os.Exit(0)
	}()
}
