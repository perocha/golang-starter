package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	ID        int
	FirstName string `json:"name"`
	LastName  string
	Address   string `json:"address,omitempty"`
}

type Employee struct {
	Person
	ManagerID int
}

type Contractor struct {
	Person
	CompanyID int
}

func main() {
	employeeList := []Employee{
		Employee{
			Person: Person{
				LastName: "Doe", FirstName: "John",
			},
		},
		Employee{
			Person: Person{
				LastName: "Campbell", FirstName: "David",
			},
		},
	}

	jsonData, _ := json.Marshal(employeeList)
	fmt.Printf("%s\n", jsonData)

	var decoded []Employee
	json.Unmarshal(jsonData, &decoded)
	fmt.Printf("%v", decoded)
}
