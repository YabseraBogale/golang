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

var player = MustLodImage("assets/player/RedDinosaur1.png")
var hill_layer_01 = MustLodImage("assets/background/Hills Layer 01.png")
var hill_layer_02 = MustLodImage("assets/background/Hills Layer 02.png")
var hill_layer_03 = MustLodImage("assets/background/Hills Layer 03.png")
var hill_layer_04 = MustLodImage("assets/background/Hills Layer 04.png")
var hill_layer_05 = MustLodImage("assets/background/Hills Layer 05.png")

func (g *Game) Draw(screen *ebiten.Image) {

	sw, sh := screen.Bounds().Dx(), screen.Bounds().Dy()

	layers := []*ebiten.Image{
		hill_layer_01,
		hill_layer_02,
		hill_layer_03,
		hill_layer_04,
		hill_layer_05,
	}
	for _, layer := range layers {
		lw, lh := layer.Bounds().Dx(), layer.Bounds().Dy()

		scaleX, scaleY := float64(sw)/float64(lw), float64(sh)/float64(lh)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(scaleX, scaleY)
		screen.DrawImage(layer, op)
	}
	player_opition := &ebiten.DrawImageOptions{}
	player_opition.GeoM.Scale(2, 2)
	player_opition.GeoM.Translate(30, 250)
	screen.DrawImage(player, player_opition)

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
