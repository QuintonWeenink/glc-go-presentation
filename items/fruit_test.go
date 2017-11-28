package items

import (
	"reflect"
	"testing"
)

func TestFruitGetName(t *testing.T) {
	var fruit Item = Fruit{"Apple", 20}

	expected := "Apple"

	if !reflect.DeepEqual(fruit.getName(), expected) {
		t.Errorf("Name() = %v, want %v", fruit.getName(), expected)
	}
}
