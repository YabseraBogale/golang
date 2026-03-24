package main

import (
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	playerX       float64
	playerY       float64
	frame         int
	tick          int
	player_sprite []*ebiten.Image
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
	player_opition.GeoM.Translate(g.playerX, g.playerY)

	current_sprite := g.player_sprite[g.frame]
	screen.DrawImage(current_sprite, player_opition)

}

func (g *Game) Layout(out_width, out_height int) (width, height int) {
	return out_width, out_height
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.playerX += 2
		g.tick++
		if g.tick%8 == 0 {
			g.frame++
		}
	}
	if g.frame > len(g.player_sprite) {
		g.frame = 0
	}
	return nil
}

func main() {
	game := &Game{
		playerX: 30,
		playerY: 250,
		player_sprite: []*ebiten.Image{
			MustLodImage("assets/player/RedDinosaur1.png"),
			MustLodImage("assets/player/RedDinosaur2.png"),
			MustLodImage("assets/player/RedDinosaur3.png"),
			MustLodImage("assets/player/RedDinosaur4.png"),
			MustLodImage("assets/player/RedDinosaur5.png"),
			MustLodImage("assets/player/RedDinosaur6.png"),
			MustLodImage("assets/player/RedDinosaur7.png"),
			MustLodImage("assets/player/RedDinosaur8.png"),
			MustLodImage("assets/player/RedDinosaur9.png"),
			MustLodImage("assets/player/RedDinosaur10.png"),
			MustLodImage("assets/player/RedDinosaur11.png"),
			MustLodImage("assets/player/RedDinosaur12.png"),
			MustLodImage("assets/player/RedDinosaur13.png"),
			MustLodImage("assets/player/RedDinosaur14.png"),
			MustLodImage("assets/player/RedDinosaur15.png"),
			MustLodImage("assets/player/RedDinosaur16.png"),
			MustLodImage("assets/player/RedDinosaur17.png"),
			MustLodImage("assets/player/RedDinosaur18.png"),
		},
	}
	ebiten.SetWindowSize(640, 320)
	ebiten.SetWindowTitle("Dino")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatalln(err)
	}
}
