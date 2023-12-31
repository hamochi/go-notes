package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Person struct {
	id         int               // unexported fields are not encoded
	IsRobot    bool              // if no tag is specified, the field name is used
	Name       string            `json:"name"`
	Age        int               `json:"age"`
	Weight     float64           `json:"weight"`
	Nicknames  []string          `json:"nicknames"` // slices are encoded as JSON arrays
	Jobs       []Job             `json:"jobs"`
	Attributes map[string]string `json:"attributes"` // maps are encoded as JSON objects
}

type Job struct {
	Company  string     `json:"company"`
	Position string     `json:"position"`
	Started  time.Time  `json:"started"`            // encode as RFC3339
	Ended    *time.Time `json:"ended"`              // encode as null if nil
	Comments string     `json:"comments,omitempty"` // omit if empty
}

func SimpleEncode() {

	jobChangeDate := time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC)

	p := Person{
		Name:   "John Doe",
		Age:    42,
		Weight: 75.5,
		Nicknames: []string{
			"Johnny",
			"J",
		},
		Jobs: []Job{
			{
				Company:  "ACME",
				Position: "CEO",
				Started:  jobChangeDate,
			},
			{
				Company:  "ACME",
				Position: "CTO",
				Started:  time.Date(2011, 1, 1, 0, 0, 0, 0, time.UTC),
				Ended:    &jobChangeDate,
				Comments: "I was promoted",
			},
		},
		Attributes: map[string]string{
			"hair": "brown",
			"eyes": "blue",
		},
	}

	j, err := json.Marshal(p)
	// to pretty print:
	// j, err := json.MarshalIndent(p, "", " ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(j))

}
