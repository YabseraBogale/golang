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

func NewCar(speed, batteryDrain int) Car {
	return Car{
        speed: speed,
        batteryDrain: batteryDrain,
        battery:100,
        distance: 0,
    }
}

func Drive(car Car) Car {
	if car.battery==0{
        return car
    } else{
    	return Car{
            speed: 5,
            batteryDrain: 2,
            battery: car.battery-car.batteryDrain,
            distance: car.distance + car.speed,
        }
    }
}

func main() {
    speed := 5
    batteryDrain := 2
    car:= NewCar(speed, batteryDrain)
   for i:=0;i<100;i++{
       car=Drive(car)
       if car.battery==0{
           fmt.Println(i)
           break
       }
       fmt.Println(car)
   }
}





















