// Online Go compiler to run Golang program online
// Print "Hello World!" message

package main

import (
    "strconv"
	"fmt"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	return "Welcome to my party, "+name+"!"
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.


// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {

    w:=Welcome(name)
    n:="You will be sitting next to "+neighbor+"."
    t:=strconv.Itoa(table)
    d:=". Your table is "+direction+", exactly "+fmt.Sprintf("%.1f",distance)+" meters from here.\n"
    switch{
        case len(t)==1: return w+"\nYou have been assigned to table "+"00"+t+d+n
    	case len(t)==2: return w+"\nYou have been assigned to table "+"0"+t+d+n
    	default: return w+"\nYou have been assigned to table "+t+d+n
    }

}




func main(){

    fmt.Println(AssignTable("Chihiro", 127, "Akachi Chikondi", "on the left", 23.7834298))
}












