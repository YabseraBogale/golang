package main

import "fmt"

type person struct{
    name string
    age int
}

func (p *person)setter() {
    var age_input int
    var name_input string
    fmt.Print("Enter Person Name:")
    fmt.Scanf("%s",&name_input)
    fmt.Print("Enter Person Age:")
    fmt.Scanf("%d",&age_input)
    p.age,p.name=age_input,name_input

}

func main(){
    input:=person{}
    input.setter()
    fmt.Println("Person Age is",input.age)
    fmt.Println("Person Name is",input.name)


}
