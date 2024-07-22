package other

import "fmt"

type Vehicle struct {
	Name string
}

func (v Vehicle) Accelerate() string {
	return (v.Name + " accelerate")
}

func (v Vehicle) Brake() string {
	return (v.Name + " brake")
}

type VehicleInterface interface {
	Accelerate() string
	Brake() string
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

// becak
type BecakInterface interface {
	Brake() string
}
type Becak struct {
	BecakVec BecakInterface
}

func (v Becak) Ngerem(dest string) {
	res := v.BecakVec.Brake()
	fmt.Println(res + " karena " + dest)
}

func ExampleInterface() {
	t := Train{
		TrainVec: NewVehicle("jayabaya"),
	}
	t.Gotodest("malang")

	b := Becak{
		BecakVec: NewVehicle("Becak"),
	}
	b.Ngerem("Turunan")
}
