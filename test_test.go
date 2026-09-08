package gosygr

import (
	"testing"
	"strconv"

	"fmt"
)

type testCar struct {
	Engine testEngine
	Seats int
	Model string
	Passengers []string
}

type testEngine struct {
	cylinders int
}

func (t *testEngine) FromIr(pI *Struct) error {
	cylindersStr, ok := pI.stringFields["cylinders"]
	if !ok {
		return fmt.Errorf("Cylinders field not found in ir")
	}
	cylinders, err := strconv.Atoi(cylindersStr)
	if err != nil {
		return fmt.Errorf("Converting cylinders from string to integer, got: %w", err)
	}
	t.cylinders = cylinders
	return nil
}

func (t *testEngine) ToIr(pI *Struct) {
	*pI = Struct{}
	pI.stringFields = map[string]string{"cylinders": strconv.Itoa(t.cylinders)}
}

func (t *testCar) FromIr(pI *Struct) error {
	engineI, ok := pI.structFields["Engine"]
	if !ok {
		return fmt.Errorf("Engine field not found in IR.")
	}
	err := (t.Engine).FromIr(engineI)
	if err != nil {
		return fmt.Errorf("Parsing ir struct to engine, got: %w", err)
	}
	seats, err := strconv.Atoi(pI.stringFields["Seats"])
	if err != nil {
		return fmt.Errorf("Parsing seats from string to int, got: %w", err)
	}
	t.Seats = seats 
	model, ok := pI.stringFields["Model"]
	if !ok {
		return fmt.Errorf("Model field not found in ir struct")
	}
	t.Model = model
	passengers, ok := pI.arrayFields["Passengers"]
	if !ok {
		return fmt.Errorf("Passengers field not found in ir struct")
	}
	t.Passengers = passengers.strings
	return nil
}

func (t *testCar) ToIr(ir *Struct) error {
	seatsK := "Seats"
	seatsS := strconv.Itoa(t.Seats)
	modelK := "Model"
	modelS := t.Model
	enginesK := "Engine"
	enginesS := Struct{
		stringFields: map[string]string{},
		structFields: map[string]*Struct{},
		arrayFields: map[string]*Array{},
		order: []string{},
	}
	(&t.Engine).ToIr(&enginesS) 
	passengersK := "Passengers"
	passengersI := Array{
		strings: t.Passengers,
		structs: []*Struct{},
		arrays: []*Array{},
	}
	*ir = Struct{
		stringFields: map[string]string{seatsK: seatsS, modelK: modelS,},
		structFields: map[string]*Struct{enginesK: &enginesS},
		arrayFields: map[string]*Array{passengersK: &passengersI},
		order: []string{enginesK, seatsK, modelK, passengersK},
	}
	return nil
}

func TestFromTestCarToIr(t *testing.T) {
	tcar1 := testCar{
		Engine: testEngine{
			cylinders: 78,
		},
		Seats: 4, 
		Model: "Some model",
		Passengers: []string{"John", "Bill", "Bob"},
	}
	ir := Struct{}
	err := tcar1.ToIr(&ir)
	if err != nil {
		t.Errorf("ERROR: parsing from car to ir struct, got: %v", err)
	}
	tcar2 := testCar{}
	err = tcar2.FromIr(&ir)
	if err != nil {
		t.Errorf("ERROR: parsing from ir to car struct, got: %v", err)
	}
	fmt.Printf("TEST: %+v \n %+v \n %+v \n", tcar1, ir, tcar2)
}

func TestOne(t *testing.T) {
	fmt.Println("test")
}


