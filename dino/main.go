package main

type Game struct {
}

func (g *Game) Draw() error {
	return nil
}

func (g *Game) Layout(out_width, out_height int) (width, height int) {
	return out_width, out_height
}

func (g *Game) Update() {

}

func main() {

}
