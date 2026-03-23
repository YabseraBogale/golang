package main

import (
	"embed"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
}

var assests embed.FS
var player = must_load_image("RedDino/RedDinosaur1.png")

func must_load_image(name string) *ebiten.Image {
	img, err := assests.Open(name)

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

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(player, nil)
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
