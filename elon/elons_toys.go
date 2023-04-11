package elon

import "fmt"

func (car *Car) Drive() {
	if car.battery-car.batteryDrain < 0 {
		return
	}
	car.battery = car.battery - car.batteryDrain
	car.distance += car.speed
}

func (car Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

func (car Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

func (car Car) CanFinish(trackDistance int) bool {
	batteryRequired := trackDistance * car.batteryDrain / car.speed
	return batteryRequired <= car.battery
}
