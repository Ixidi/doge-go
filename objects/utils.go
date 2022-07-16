package objects

import (
	"doge/math"
	"doge/windows"
)

func readFromBuffPanic(v any, offset uint32, buff windows.MemoryBuff) {
	if err := buff.Read(v, offset); err != nil {
		panic(err)
	}
}

func readVectorFromBuffPanic(v *math.Vector3, offset uint32, buff windows.MemoryBuff) {
	var x, y, z float32
	readFromBuffPanic(&x, offset, buff)
	readFromBuffPanic(&y, offset+4, buff)
	readFromBuffPanic(&z, offset+8, buff)

	*v = math.Vector3{X: float64(x), Y: float64(y), Z: float64(z)}
}

func readFromMemPanic(v any, offset uint32, mem windows.Memory) {
	if err := mem.Read(v, offset); err != nil {
		panic(err)
	}
}

func readBuffFromMemPanic(size uint, address uint32, mem windows.Memory) windows.MemoryBuff {
	buff, err := mem.ReadBuff(size, address)
	if err != nil {
		panic(err)
	}
	return buff
}
