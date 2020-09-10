package simple

import (
	"testing"
)

func TestSimple(t *testing.T) {
	petFactory := PetFactory{}

	petFactory.CreatePet("cat").call()
}
