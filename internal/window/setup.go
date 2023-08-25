package window

import (
	"golang.org/x/image/colornames"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/lxn/win"
)

func Setup() (*pixelgl.Window, error) {
	cfg := pixelgl.WindowConfig{
		Title:  windowTitle,
		Bounds: pixel.R(0, 0, windowWidth, windowHeight),
		VSync:  true,
	}

	w, h := getScreenDimensions()
	cfg.Bounds = pixel.R(0, 0, float64(w), float64(h))

	win, err := pixelgl.NewWindow(cfg)
	win.Clear(colornames.Lightskyblue)
	return win, err
}

func getScreenDimensions() (int, int) {
	hDC := win.GetDC(0)
	defer win.ReleaseDC(0, hDC)
	width := int(win.GetDeviceCaps(hDC, win.HORZRES))
	height := int(win.GetDeviceCaps(hDC, win.VERTRES))
	return width, height
}
