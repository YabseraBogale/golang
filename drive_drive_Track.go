// Online Go compiler to run Golang program online
// Print "Hello World!" message

package main
import "fmt"


type Car struct{
    battery int
    batteryDrain int
    speed int
    distance int
}


func Drive(car Car) Car {
	if car.battery==0{
        return car
    } else{
    	return Car{
            speed: car.speed,
            batteryDrain: car.batteryDrain,
            battery: car.battery-car.batteryDrain,
            distance: car.distance + car.speed,
        }
    }
}

func main() {

    car:= Car{battery:10 ,batteryDrain:10 ,speed:18, distance:0}
    //car=Drive(car)
    fmt.Println(car)
   for i:=0;i<100;i++{
       car=Drive(car)
       fmt.Println(i)
       if car.battery==0{
           fmt.Println(i)
           break
       }
       fmt.Println(car)
   }
}





















