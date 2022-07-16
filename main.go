package main

import (
	"doge/objects"
	"doge/windows"
	"fmt"
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

	//app.Start()
	process, err := windows.GetProcess("League of Legends.exe")
	if err != nil {
		log.Panic(err)
	}
	mem := windows.NewMemory(process)
	objectManager, err := objects.NewObjectManager(mem)
	if err != nil {
		log.Panic(err)
	}
	for {
		fmt.Printf("%+v\n", objectManager.Game())
		time.Sleep(time.Second)
		objectManager.Reread()
	}

}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
