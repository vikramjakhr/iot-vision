package main

import (
	_ "gitlab.intelligrape.net/tothenew/vision/routers"
	"github.com/astaxie/beego"
	"github.com/howeyc/fsnotify"
	"log"
	"os"
	"os/signal"
	"gitlab.intelligrape.net/tothenew/vision/services"
	"strings"
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
				log.Println("Event received : ", ev.String())
				go func() {
					if ev.IsCreate() {
						name := ev.Name[strings.LastIndex(ev.Name, "/")+1:]
						services.DetectText(ev.Name, name)
					}
				}()
			case err := <-watcher.Error:
				log.Println("Error while listening to event ", err)
			}
		}
	}()

	err = watcher.Watch("/home/infra/ftp/20170818/images")
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
