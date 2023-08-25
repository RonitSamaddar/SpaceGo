package game

import (
	"fmt"

	"github.com/RonitSamaddar/SpaceGo/internal/window"

	"github.com/faiface/pixel/pixelgl"
)

func Start() {
	fmt.Println("SpaceGo started!")
	pixelgl.Run(setup)
}

func setup() {
	fmt.Println("We are inside run!!")
	gameWindow, err := window.Setup()
	if err != nil {
		panic(fmt.Sprintf("Game window setup failed : %+v", err))
	}

	for !gameWindow.Closed() {
		gameWindow.Update()
	}
}
