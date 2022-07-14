package app

import (
	"doge/windows"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
)

type Application struct {
	firstUpdate bool
	windowName  string
}

func (a *Application) Update() error {
	if a.firstUpdate {
		handle, err := windows.FindWindow(a.windowName)
		if err != nil {
			print(err)
			return err
		}
		go func() {
			windows.SetWindowLong(handle)
			windows.SetForegroundWindow(handle)
		}()
		a.firstUpdate = false
	}

	return nil
}

func (a *Application) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, 40, 40, 40, 40, colornames.Red)
}

func (a *Application) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1920, 1080
}

func Start() error {
	application := &Application{true, "eseklnses"}
	ebiten.SetWindowSize(1920, 1080)
	ebiten.SetWindowTitle(application.windowName)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeDisabled)
	ebiten.SetWindowDecorated(false)
	ebiten.SetScreenTransparent(true)
	ebiten.SetWindowFloating(true)
	ebiten.SetInitFocused(true)
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOn)
	ebiten.SetInitFocused(true)

	return ebiten.RunGame(application)
}
