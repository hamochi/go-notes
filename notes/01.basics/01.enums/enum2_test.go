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
