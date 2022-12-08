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
