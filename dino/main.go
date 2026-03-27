package main

import (
	"image"
	_ "image/png"
	"log"
	"math"
	"math/rand"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type Apple struct {
	x, y          float64
	width, height float64
	apple_health  string
	image         *ebiten.Image
}

type Game struct {
	apple         []*Apple
	playerX       float64
	playerY       float64
	velocity_Y    float64
	is_jumping    bool
	frame         int
	tick          int
	drop_timer    int
	next_drop_in  int
	cameraX       float64
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
var apples = []*ebiten.Image{MustLodImage("assets/apple/apple_regular_30_30px.png"), MustLodImage("assets/apple/apple_golden_30_30px.png"), MustLodImage("assets/apple/apple_rotten_30_30px.png")}

func Collision(x1, y1, w1, h1, x2, y2, w2, h2 float64) bool {
	return x1 < x2+w2 &&
		x1+w1 > x2 &&
		y1 < y2+h2 &&
		y1+h1 > y2
}

func (g *Game) Draw(screen *ebiten.Image) {

	sw, sh := screen.Bounds().Dx(), screen.Bounds().Dy()

	layers := []*ebiten.Image{
		hill_layer_01,
		hill_layer_02,
		hill_layer_03,
		hill_layer_04,
		hill_layer_05,
	}

	for i, layer := range layers {
		bw, bl := layer.Bounds().Dx(), layer.Bounds().Dy()

		scaleX := float64(sw) / float64(bw)
		scaleY := float64(sh) / float64(bl)

		speed := float64(i+1) * 0.5

		scale_width := float64(bw) * scaleX

		offset := fmod(-g.cameraX*speed, scale_width)

		opition_before := &ebiten.DrawImageOptions{}

		opition_before.GeoM.Scale(scaleX, scaleY)
		opition_before.GeoM.Translate(offset, 0)

		screen.DrawImage(layer, opition_before)

		opition_after := &ebiten.DrawImageOptions{}

		opition_after.GeoM.Scale(scaleX, scaleY)

		opition_after.GeoM.Translate(offset+scale_width, 0)

		screen.DrawImage(layer, opition_after)

	}

	for _, apple := range g.apple {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(apple.x-g.cameraX, apple.y)
		screen.DrawImage(apple.image, op)
	}

	player_opition := &ebiten.DrawImageOptions{}
	player_opition.GeoM.Scale(2, 2)
	player_opition.GeoM.Translate(g.playerX-g.cameraX, g.playerY)

	current_sprite := g.player_sprite[g.frame]
	screen.DrawImage(current_sprite, player_opition)

}

func (g *Game) Layout(out_width, out_height int) (width, height int) {
	return 640, 320
}

func (g *Game) Update() error {
	const gravity = 0.6
	const jump_strength = -12.0
	const ground_y = 250

	if ebiten.IsKeyPressed(ebiten.KeySpace) && !g.is_jumping {
		g.velocity_Y = jump_strength
		g.is_jumping = true
	}

	g.playerY += g.velocity_Y
	g.velocity_Y += gravity

	if g.playerY >= ground_y {
		g.playerY = ground_y
		g.velocity_Y = 0
		g.is_jumping = false
	}

	if ebiten.IsKeyPressed(ebiten.KeyF11) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}

	g.drop_timer++
	pW, pH := 32.0, 32.0
	if g.drop_timer > g.next_drop_in {
		index := rand.Intn(3)
		var health string
		if index == 0 {
			health = "Normal"
		} else if index == 1 {
			health = "Cured"
		} else {
			health = "Sick"
		}
		new_apple := &Apple{
			x:            g.cameraX + 640,
			y:            280,
			width:        20,
			height:       20,
			image:        apples[index],
			apple_health: health,
		}
		g.apple = append(g.apple, new_apple)
		g.drop_timer = 0
		g.next_drop_in = rand.Intn(120) + 60
	}
	for i := len(g.apple) - 1; i >= 0; i-- {

		apple := g.apple[i]

		apple_width, apple_height := 20.0, 20.0

		if g.apple[i].x < g.cameraX-50 {
			g.apple = append(g.apple[:i], g.apple[i+1:]...)
		}

		if Collision(g.playerX, g.playerY, pW, pH, apple.x, apple.y, apple_width, apple_height) {

		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.playerX += 2
		g.tick++
		if g.tick%8 == 0 {
			g.frame++
		}
	}

	if g.frame >= len(g.player_sprite) {
		g.frame = 0
	}
	g.cameraX = g.playerX - 320
	return nil
}

func fmod(x, y float64) float64 {
	res := math.Mod(x, y)
	if res > 0 {
		res -= y
	}
	return res
}

func main() {
	game := &Game{
		playerX: 30,
		playerY: 250,
		player_sprite: []*ebiten.Image{
			MustLodImage("assets/player/RedDinosaur1.png"),
			MustLodImage("assets/player/RedDinosaur9.png"),
			MustLodImage("assets/player/RedDinosaur10.png"),
			MustLodImage("assets/player/RedDinosaur11.png"),
			MustLodImage("assets/player/RedDinosaur12.png"),
			MustLodImage("assets/player/RedDinosaur13.png"),
			MustLodImage("assets/player/RedDinosaur14.png"),
			MustLodImage("assets/player/RedDinosaur15.png"),
			MustLodImage("assets/player/RedDinosaur16.png"),
		},
	}
	ebiten.SetWindowSize(640, 320)
	ebiten.SetWindowTitle("Dino")
	ebiten.SetFullscreen(true)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatalln(err)
	}
}
