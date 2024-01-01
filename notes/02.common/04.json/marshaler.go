package main

import (
	"encoding/json"
	"fmt"
)

type Human struct {
	firstname string
	lastname  string
	age       int
}

func (p Human) FirstName() string {
	return p.firstname
}

func (p Human) LastName() string {
	return p.lastname
}

func (p Human) Age() int {
	return p.age
}

func (p Human) MarshalJSON() ([]byte, error) {
	// We could return a JSON string here, but it's not a good idea.
	// return []byte(fmt.Sprintf(`{"fullname": "%s %s", "age_in_dog_years": %d}`, p.firstname, p.lastname, p.age*7)), nil

	// Instead we create a temporary struct and marshal that.
	type TempHuman struct {
		Name          string `json:"name"`
		AgeInDogYears int    `json:"age_in_dog_years"`
	}

	temp := TempHuman{
		Name:          p.FirstName() + " " + p.LastName(),
		AgeInDogYears: p.Age() * 7,
	}

	return json.Marshal(temp)
}

func NewHuman(firstname, lastname string, age int) Human {
	return Human{
		firstname: firstname,
		lastname:  lastname,
		age:       age,
	}
}

func MarshalerExample() {
	p := NewHuman("John", "Doe", 42)

	data, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
