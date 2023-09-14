// Online Go compiler to run Golang program online
// Print "Hello World!" message

package main
import "fmt"

type point struct{
    x float32
    y float32
    z float32
}

func (p *point) length(){
    p.x+=1
    p.y+=2
    p.z+=3
}

func (p *point) moving(x float32,y float32,z float32){
    p.x+=x
    p.y+=y
    p.z+=z
}

func main() {
    p:=point{}
    p.length()
  fmt.Println(p)
  p.moving(1.2,0.1,3.1)
   fmt.Println(p)
}
