package objects

import (
	"doge/math"
	"doge/offsets"
	"doge/windows"
)

type Game struct {
	GameTime             float32
	ViewProjectionMatrix [16]float32
	WindowWidth          int32
	WindowHeight         int32
}

func ReadGame(mem windows.Memory) (game Game, err error) {
	baseAddress := mem.Process().BaseAddress

	if err != nil {
		return
	}

	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()

	readFromMemPanic(&game.GameTime, baseAddress+offsets.GameTime, mem)

	buff := readBuffFromMemPanic(132, baseAddress+offsets.ViewProjMatrices, mem)

	var viewMatrix, projectionMatrix [16]float32

	for i := 0; i < 16; i++ {
		readFromBuffPanic(&viewMatrix[i], uint32(i*4), buff)
	}
	for i := 0; i < 16; i++ {
		readFromBuffPanic(&projectionMatrix[i], 64+uint32(i*4), buff)
	}

	viewProjectionMatrix, err := math.Multiple4x4Matrices(viewMatrix, projectionMatrix)
	if err != nil {
		return
	}
	game.ViewProjectionMatrix = viewProjectionMatrix

	var (
		renderAddress uint32
	)
	readFromMemPanic(&renderAddress, baseAddress+offsets.Renderer, mem)

	buff = readBuffFromMemPanic(offsets.GameWindowHeight+4, renderAddress, mem)
	readFromBuffPanic(&game.WindowWidth, offsets.GameWindowWidth, buff)
	readFromBuffPanic(&game.WindowHeight, offsets.GameWindowHeight, buff)

	return
}
