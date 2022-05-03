
# Enums
There are no enumns types in Go, but we can use constants and variables to get something similar. A good way to insure the type safety of enums is to define them as their own types.
This way if you have Month = 12 and Bananas = 12, then Month == Bananas will be false.
It's also useful to create a String() method for that type so packages like fmt and log
can print out the textual values of the enum.


```go
package enums

type Month int

const (
	January Month = iota + 1
	February
	March
	May
	June
	July
	August
	September
	October
	November
	December
)

func (m Month) String() string {
	switch m {
	case January:
		return "january"
	case February:
		return "february"
	case March:
		return "march"
	case May:
		return "may"
	case June:
		return "june"
	case July:
		return "july"
	case August:
		return "august"
	case September:
		return "september"
	case October:
		return "october"
	case November:
		return "november"
	case December:
		return "december"
	}
	return "unknown"
}

func MyBirthMonth() Month {
	return December
}
```

```go
package enums

import (
	"fmt"
	"testing"
)

func TestMyBirthMonth(t *testing.T) {
	bMonth := MyBirthMonth()
	if bMonth != December {
		t.Fatalf("unexpected month")
	}

	const Bananas = 12

	if bMonth == Bananas {
		t.Fatalf("type missmatch")
	}

	fmt.Println(bMonth)
}
```

Output:
```
december
```

## Using variables and structs
In this examples we create a type called Color that is a Struct and another type called NamedColor based on Color. 
We then create our enumns based on NamedColor.

```go
package enums

type Color struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

type NamedColor Color

func (n NamedColor) String() string {
	switch n {
	case Yellow:
		return "yellow"
	case Cyan:
		return "cyan"
	case Magenta:
		return "magenta"
	case Silver:
		return "silver"
	}
	return "unknown color"
}

var (
	Yellow  = NamedColor{255, 255, 0}
	Cyan    = NamedColor{0, 255, 255}
	Magenta = NamedColor{255, 0, 255}
	Silver  = NamedColor{192, 192, 192}
)

func ReturnMagenta() NamedColor {
	return Magenta
}
```

Here we can also see that we can type cast NamedColor back to Color to check equality.

```go
package enums

import (
	"fmt"
	"testing"
)

func TestReturnMagenta(t *testing.T) {
	myColor := ReturnMagenta()
	if myColor != Magenta {
		t.Fatal("Unexpected color")
	}

	newColor := Color{255, 0, 255}
	if Color(Magenta) != newColor {
		t.Fatalf("Our newColor is not magenta, newColor: %v, Magenta: %v", newColor, Color(Magenta))
	}

	fmt.Println(myColor)
}
```

output:
```
magenta
```