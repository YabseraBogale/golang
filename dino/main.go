package main

import (
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
}

func MustLodImage(name string) *ebiten.Image {
	img, err := os.Open(name)

	if err != nil {
		panic(err)
	}
	defer img.Close()

	i, _, err := image.Decode(img)

	if err != nil {
		panic(err)
	}
	return ebiten.NewImageFromImage(i)
}

var player = MustLodImage("assets/RedDinosaur1.png")

func (g *Game) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(2, 2)
	op.GeoM.Translate(150, 200)

	screen.DrawImage(player, op)
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
