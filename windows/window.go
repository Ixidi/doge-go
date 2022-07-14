package windows

import (
	"github.com/lxn/win"
	"syscall"
)

func FindWindow(name string) (win.HWND, error) {
	lpName, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return 0, err
	}
	return win.FindWindow(nil, lpName), nil
}

func SetWindowLong(handle win.HWND) {
	win.SetWindowLong(handle, win.GWL_EXSTYLE, win.WS_EX_COMPOSITED|win.WS_EX_LAYERED|win.WS_EX_TRANSPARENT|win.WS_EX_TOOLWINDOW|win.WS_EX_TOPMOST)
}

func SetForegroundWindow(handle win.HWND) {
	win.SetForegroundWindow(handle)
}
