package windows

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type Process struct {
	Handle      windows.Handle
	BaseAddress uintptr
}

func (p *Process) ReadMemory(buff *[]byte, address uint32) error {
	return windows.ReadProcessMemory(p.Handle, uintptr(address), &(*buff)[0], uintptr(len(*buff)), nil)
}

func GetProcess(name string) (Process, error) {
	id, err := getProcessId(name)
	if err != nil {
		return Process{}, err
	}

	handle, err := getProcessHandle(id)
	if err != nil {
		return Process{}, err
	}

	var module windows.Handle
	var cbNeeded uint32

	err = windows.EnumProcessModules(handle, &module, uint32(unsafe.Sizeof(module)), &cbNeeded)
	if err != nil {
		return Process{}, err
	}

	var modInfo windows.ModuleInfo

	err = windows.GetModuleInformation(handle, module, &modInfo, uint32(unsafe.Sizeof(modInfo)))
	if err != nil {
		return Process{}, err
	}

	return Process{handle, modInfo.BaseOfDll}, nil
}

func getProcessHandle(id uint32) (windows.Handle, error) {
	return windows.OpenProcess(uint32(0x1F0FFF), true, id)
}

func getProcessId(name string) (uint32, error) {
	h, e := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPPROCESS, 0)
	if e != nil {
		return 0, e
	}
	p := windows.ProcessEntry32{Size: 568}
	for {
		e := windows.Process32Next(h, &p)
		if e != nil {
			return 0, e
		}
		if windows.UTF16ToString(p.ExeFile[:]) == name {
			return p.ProcessID, nil
		}
	}
}
