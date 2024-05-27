package other

import "fmt"

type Vehicle struct {
	Name string
}

func (v Vehicle) Accelerate() string {
	return (v.Name + " accelerate")
}

func (v Vehicle) Brake() {
	fmt.Println(v.Name + " brake")
}

type VehicleInterface interface {
	Accelerate() string
	Brake()
}

func NewVehicle(name string) VehicleInterface {
	return Vehicle{Name: name}
}

// train
type TrainInterface interface {
	Accelerate() string
}
type Train struct {
	TrainVec TrainInterface
}

func (v Train) Gotodest(dest string) {
	res := v.TrainVec.Accelerate()
	fmt.Println(res + " to " + dest)
}

func ExampleInterface() {
	t := Train{
		TrainVec: NewVehicle("jayabaya"),
	}
	t.Gotodest("malang")
}
