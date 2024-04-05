package main

import (
	"log"
	"time"
)

func HelloRepeat() {
	interval := 5
	ticker := time.NewTicker(time.Second * time.Duration(interval))
	for {
		<-ticker.C
		log.Println("Hello, World!")
	}

}

func main() {
	HelloRepeat()
}
