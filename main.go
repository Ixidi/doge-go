package main

import (
	"doge/app"
	"log"
	"time"
)

func main() {
	/*log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	process, err := windows.GetProcess("League of Legends.exe")
	if err != nil {
		log.Panic(err)
	}
	mem := windows.NewMemory(process)*/

	app.Start()
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
