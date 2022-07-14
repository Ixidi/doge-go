package win

import (
	"bytes"
	"encoding/binary"
	"errors"
	"reflect"
	"unsafe"
)

type MemoryBuff interface {
	Read(v any, offset uint32) error
}

type Memory interface {
	Process() *Process
	ReadBuff(bytesCount uint, address uint32) (MemoryBuff, error)
	Read(v any, address uint32) error
}

type memory struct {
	process *Process
}

type memoryBuff struct {
	b   []byte
	mem Memory
}

func (m memory) Process() *Process {
	return m.process
}

func (m memory) Read(v any, address uint32) error {
	if reflect.TypeOf(v).Kind() != reflect.Pointer {
		return errors.New("v was expected to be a pointer")
	}

	switch t := v.(type) {
	case *string:
		var (
			strAddress uint32
			strSize    uint32
		)
		if err := m.Read(&strAddress, address); err != nil {
			return err
		}
		if err := m.Read(&strSize, address+4); err != nil {
			return err
		}

		buff := make([]byte, strSize)
		if err := m.process.ReadMemory(&buff, strAddress); err != nil {
			return err
		}

		*t = string(buff)
		return nil
	default:
		typeSize := unsafe.Sizeof(&v)
		buff := make([]byte, typeSize)
		err := m.process.ReadMemory(&buff, address)
		if err != nil {
			return err
		}

		return binary.Read(bytes.NewReader(buff), binary.LittleEndian, v)
	}
}

func (m memory) ReadBuff(bytesCount uint, address uint32) (MemoryBuff, error) {
	buff := make([]byte, bytesCount)
	if err := m.process.ReadMemory(&buff, address); err != nil {
		return nil, err
	}
	return memoryBuff{buff, &m}, nil
}

func (m memoryBuff) Read(v any, offset uint32) error {
	if reflect.TypeOf(v).Kind() != reflect.Pointer {
		return errors.New("v was expected to be a pointer")
	}

	switch t := v.(type) {
	case *string:
		var (
			strAddress uint32
			strSize    uint32
		)
		if err := m.Read(&strAddress, offset); err != nil {
			return err
		}
		if err := m.Read(&strSize, offset+4); err != nil {
			return err
		}

		buff := make([]byte, strSize)
		if err := m.mem.Process().ReadMemory(&buff, strAddress); err != nil {
			return err
		}

		*t = string(buff)
		return nil
	default:
		typeSize := uint32(unsafe.Sizeof(&v))
		return binary.Read(bytes.NewReader(m.b[offset:offset+typeSize]), binary.LittleEndian, v)
	}
}

func NewMemory(process Process) Memory {
	return memory{&process}
}
