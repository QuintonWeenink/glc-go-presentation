package items

type Fruit struct {
	name string
	amount int
}

func (fruit Fruit) getName() string {
	return fruit.name
}
