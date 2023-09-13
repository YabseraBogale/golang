package main
import (
    "strings"
    "fmt"
)
// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.


// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1 string,option2 string) string{
	optional_1:=strings.Split(option1," ")[1]
    optional_2:=strings.Split(option2," ")[1]
    if optional_1>optional_2{
        return option1 + " is clearly the better choice."
    }
	return option2 +" is clearly the better choice."
}

func main(){
    o1:=ChooseVehicle("Bugatti Veyron","ford")
    fmt.Println(o1)
    fmt.Println(len(o1))
}
