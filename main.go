package main

import (
	"doge/objects"
	"doge/offsets"
	"doge/win"
	"fmt"
)

func main() {
	process, _ := win.GetProcess("League of Legends.exe")
	mem := win.NewMemory(process)

	/*var (
		localPlayer uint32
		health      float32
	)
	_ = mem.Read(&localPlayer, uint32(process.BaseAddress+0x310ed68))
	_ = mem.Read(&health, localPlayer+0x0E74)
	fmt.Println(health)*/

	var (
		localPlayer uint32
	)

	err := mem.Read(&localPlayer, uint32(process.BaseAddress+offsets.LocalPlayer))
	if err != nil {
		panic(err)
	}

	obj, err := objects.ReadGameObject(mem, localPlayer)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", obj)
}
