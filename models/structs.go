package models

import "reflect"

type Computers struct {
	Name      string
	Price     float64
	Memory    string
	Model     string
	Processor string
	RAM       string
}

func (c Computers) ValueOf() reflect.Value {

	return reflect.ValueOf(c)
}

func (c Computers) StructFields() reflect.Type {
	//TODO implement me
	return c.ValueOf().Type()
}

func (c Computers) NumOfFields() int {
	//TODO implement me
	return c.ValueOf().NumField()
}

type Person struct {
	Name  string
	Age   int
	Email string
}

func (p Person) ValueOf() reflect.Value {
	return reflect.ValueOf(p)
}

func (p Person) StructFields() reflect.Type {
	return p.ValueOf().Type()
}

func (p Person) NumOfFields() int {
	return p.ValueOf().NumField()
}

type Phones struct {
	ID    int
	Name  string
	Model string
	Price float64
}

func (p Phones) ValueOf() reflect.Value {
	return reflect.ValueOf(p)
}

func (p Phones) StructFields() reflect.Type {
	return p.ValueOf().Type()
}

func (p Phones) NumOfFields() int {
	return p.ValueOf().NumField()
}

type Config struct {
	DB   string `json:"db"`
	Port int64  `json:"portRun"`
}
