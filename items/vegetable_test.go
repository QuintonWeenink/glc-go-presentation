package items

import (
	"reflect"
	"testing"
)

func TestVegetableGetName(t *testing.T) {
	var vegetable Item = Vegetable{"Carrot", 20}

	expected := "Carrot"

	if !reflect.DeepEqual(vegetable.getName(), expected) {
		t.Errorf("Name() = %v, want %v", vegetable.getName(), expected)
	}
}
