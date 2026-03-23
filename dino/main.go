package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
}

func (g *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Layout(out_width, out_height int) (width, height int) {
	return out_width, out_height
}

func (g *Game) Update() error {
	return nil
}

func main() {
	game := &Game{}
	ebiten.SetWindowSize(640, 320)
	ebiten.SetWindowTitle("Dino")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatalln(err)
	}
}
