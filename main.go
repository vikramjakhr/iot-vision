package main

import (
	_ "gitlab.intelligrape.net/tothenew/vision/routers"
	"github.com/astaxie/beego"
	"github.com/howeyc/fsnotify"
	"log"
	"os"
	"os/signal"
	"gitlab.intelligrape.net/tothenew/vision/services"
	"time"
	"math/rand"
)

var r *rand.Rand // Rand for this package.

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomString(strlen int) string {
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)
	for i := range result {
		result[i] = chars[r.Intn(len(chars))]
	}
	return string(result)
}

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
				go func() {
					//if ev.IsCreate() || ev.IsModify() {
					if ev.IsCreate() {
						log.Println("Event received : ", ev.String())
						name := RandomString(50)
						services.Process(ev.Name, name)
					}
				}()
			case err := <-watcher.Error:
				log.Println("Error while listening to event ", err)
			}
		}
	}()

	//err = watcher.Watch("/home/infra/ftp/20170818/images")
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
