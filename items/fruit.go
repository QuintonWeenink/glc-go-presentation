package items

// Fruit Object
type Fruit struct {
	name   string
	amount int
}

func (fruit Fruit) getName() string {
	return fruit.name
}
