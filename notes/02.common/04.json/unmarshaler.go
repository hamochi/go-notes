package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

var j = `{
		"id": "1",
		"created_at": 1704112360
	}`

type Data struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

func (d *Data) UnmarshalJSON(b []byte) error {
	type Temp struct {
		ID        string `json:"id"`
		CreatedAt int    `json:"created_at"`
	}

	var temp Temp
	err := json.Unmarshal(b, &temp)
	if err != nil {
		return err
	}

	// convert temp.ID to int and assign it to d.ID
	id, err := strconv.Atoi(temp.ID)
	if err != nil {
		return err
	}
	d.ID = id

	// convert temp.CreatedAt to time.Time and assign it to d.CreatedAt
	d.CreatedAt = time.Unix(int64(temp.CreatedAt), 0)

	return nil
}

func UnamarshalerExample() {
	// The json data id is a string and created_at is an int (unix timestamp)
	// but we want to unmarshal them into a struct with an int and a time.Time

	var d Data
	err := json.Unmarshal([]byte(j), &d)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", d)
}
