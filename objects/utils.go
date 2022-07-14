package objects

import (
	"doge/math"
	"doge/win"
)

func readPanic(v any, offset uint32, buff win.MemoryBuff) {
	if err := buff.Read(v, offset); err != nil {
		panic(err)
	}
}

func readVectorPanic(v *math.Vector3, offset uint32, buff win.MemoryBuff) {
	var x, y, z float32
	readPanic(&x, offset, buff)
	readPanic(&y, offset+4, buff)
	readPanic(&z, offset+8, buff)

	*v = math.Vector3{X: float64(x), Y: float64(y), Z: float64(z)}
}
