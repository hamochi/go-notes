
# JSON
1. [Encoding](#encoding)
  1. [OmitEmpty](#omitempty)
2. [Decoding](#decoding)
3. [Marshaler interface](#marshaler-interface)
4. [Unmarshaler interface](#unmarshaler-interface)


## Encoding
```go
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
	Password   string            `json:"-"`          // this field is never encoded
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
		Password: "secret",
	}

	j, err := json.Marshal(p)
	// to pretty print:
	// j, err := json.MarshalIndent(p, "", " ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(j))

}
```
Output:
```JSON
{
 "IsRobot": false,
 "name": "John Doe",
 "age": 42,
 "weight": 75.5,
 "nicknames": [
  "Johnny",
  "J"
 ],
 "jobs": [
  {
   "company": "ACME",
   "position": "CEO",
   "started": "2015-01-01T00:00:00Z",
   "ended": null
  },
  {
   "company": "ACME",
   "position": "CTO",
   "started": "2011-01-01T00:00:00Z",
   "ended": "2015-01-01T00:00:00Z",
   "comments": "I was promoted"
  }
 ],
 "attributes": {
  "eyes": "blue",
  "hair": "brown"
 }
}
```

### OmitEmpty
The omitempty option in Go's JSON encoding works by omitting the field from the JSON output if the field has an empty value. An "empty" value means the zero value for the field's type, which includes:

- false for booleans,
 -0 for numeric types,
- "" (the empty string) for strings,
- nil for pointers, interfaces, maps, slices, and channels,
- the zero value for arrays, structs, and any other types.

However, it's important to note that for structs, omitempty will not omit the field if the struct is empty but not nil. This is because an empty struct still has a valid, non-nil zero value. If you want to omit a struct, the best solution would be to add it as a pointer in the struct.

## Decoding
As long as we know how the JSON data is structured, we can decode it into a struct. And we can use the parts of the data we need.

The struct must have exported fields, otherwise the decoder will not be able to set them.

The struct can have more fields than the JSON data, but the decoder will only set the fields that are present in the JSON data.

The struct must tag the fields with the name of the JSON field, otherwise the decoder will not be able to set them.
```go
package main

import (
	"encoding/json"
	"fmt"
)

var data = `{
		"coord": {
			"lon": 10.99,
			"lat": 44.34
		},
		"weather": [
			{
				"id": 501,
				"main": "Rain",
				"description": "moderate rain",
				"icon": "10d"
			}
		],
		"base": "stations",
		"main": {
			"temp": 298.48,
			"feels_like": 298.74,
			"temp_min": 297.56,
			"temp_max": 300.05,
			"pressure": 1015,
			"humidity": 64,
			"sea_level": 1015,
			"grnd_level": 933
		},
		"visibility": 10000,
		"wind": {
			"speed": 0.62,
			"deg": 349,
			"gust": 1.18
		},
		"rain": {
			"1h": 3.16
		},
		"clouds": {
			"all": 100
		},
		"dt": 1661870592,
		"sys": {
			"type": 2,
			"id": 2075663,
			"country": "IT",
			"sunrise": 1661834187,
			"sunset": 1661882248
		},
		"timezone": 7200,
		"id": 3163858,
		"name": "Zocca",
		"cod": 200
	}`

func SimpleDecoding() {

	type WeatherData struct {
		Coord struct {
			Lon float64 `json:"lon"`
			Lat float64 `json:"lat"`
		} `json:"coord"`
		Weather []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
	}

	var wd WeatherData
	err := json.Unmarshal([]byte(data), &wd)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", wd)
}
```
Output:
```bash
{Coord:{Lon:10.99 Lat:44.34} Weather:[{ID:501 Main:Rain Description:moderate rain Icon:10d}]}
```

## Marshaler interface
```go
type Marshaler interface {
	MarshalJSON() ([]byte, error)
}
```
Everytime we use the json.Marshal() function, the Go runtime will check if the type of the object we are trying to marshal implements the Marshaler interface. If it does, then the MarshalJSON() method will be called to marshal the object. If it doesn't, then the default marshaling will be used.

This is useful when we want to customize the way an object is marshaled. For example, we can use it to change the name of a field in the JSON output, or to omit a field from the JSON output, or marshal unexported fields.

```go
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
```

Output:
```JSON
{"name":"John Doe","age_in_dog_years":294}
```

## Unmarshaler interface
```go
type Unmarshaler interface {
	UnmarshalJSON([]byte) error
}
```
Everytime we use the json.Unmarshal() function, the Go runtime will check if the type of the object we are trying to unmarshal implements the Unmarshaler interface. If it does, then the UnmarshalJSON() method will be called to unmarshal the object. If it doesn't, then the default unmarshaling will be used.

This is useful when we want to customize the way an object is unmarshaled. For example, we can use it to change the time format of a field, or to unmarshal a field into a different type.

```go
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
```

Output:
```JSON
{ID:1 CreatedAt:2024-01-01 13:32:40 +0100 CET}
```