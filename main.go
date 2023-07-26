package main

import (
	"reflect"
	"sportsman/funcs"
	"sportsman/models"
)

func main() {

	PersonStruct := []models.Person{
		{Name: "John Doe", Age: 30, Email: "john.doe@example.com"},
		{Name: "Jane Smith", Age: 25, Email: "jane.smith@example.com"},
		{Name: "Bob Johnson", Age: 40, Email: "bob.johnson@example.com"},
	}
	v1 := reflect.ValueOf(PersonStruct)
	length1 := len(PersonStruct)
	funcs.CreateExelFile(models.Person{}, v1, length1)

	//

	ComputerStruct := []models.Computers{
		{Name: "MacBook Pro", Price: 1999.99, Memory: "512GB", Model: "A2141", Processor: "Intel Core i7", RAM: "16GB"},
		{Name: "MacBook Pro", Price: 1999.99, Memory: "512GB", Model: "A2141", Processor: "Intel Core i7", RAM: "16GB"},
		{Name: "MacBook Pro", Price: 1999.99, Memory: "512GB", Model: "A2141", Processor: "Intel Core i7", RAM: "16GB"},
		{Name: "MacBook Pro", Price: 1999.99, Memory: "512GB", Model: "A2141", Processor: "Intel Core i7", RAM: "16GB"},
		{Name: "MacBook Pro", Price: 1999.99, Memory: "512GB", Model: "A2141", Processor: "Intel Core i7", RAM: "16GB"},
		{Name: "MacBook Pro", Price: 1999.99, Memory: "512GB", Model: "A2141", Processor: "Intel Core i7", RAM: "16GB"},
		{Name: "MacBook Pro", Price: 1999.99, Memory: "512GB", Model: "A2141", Processor: "Intel Core i7", RAM: "16GB"},
		{Name: "MacBook Pro", Price: 1999.99, Memory: "512GB", Model: "A2141", Processor: "Intel Core i7", RAM: "16GB"},
		{Name: "MacBook Pro", Price: 1999.99, Memory: "512GB", Model: "A2141", Processor: "Intel Core i7", RAM: "16GB"},
		{Name: "MacBook Pro", Price: 1999.99, Memory: "512GB", Model: "A2141", Processor: "Intel Core i7", RAM: "16GB"},
		{Name: "MacBook Pro", Price: 1999.99, Memory: "512GB", Model: "A2141", Processor: "Intel Core i7", RAM: "16GB"},
	}
	v2 := reflect.ValueOf(ComputerStruct)
	length2 := len(ComputerStruct)
	funcs.CreateExelFile(models.Computers{}, v2, length2)

}
