package items

// Vegetable Object
type Vegetable struct {
	name string
	amount int
}

func (vegetable Vegetable) getName() string {
	return vegetable.name
}
