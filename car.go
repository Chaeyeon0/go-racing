package goracing

import "fmt"

type Car struct {
	Name     string
	Distance int
}

func NewCar(name string) (*Car, error) {
	if len(name) > 5 {
		return nil, fmt.Errorf("자동차 이름은 5자 이하만 가능합니다: %s", name)
	}
	return &Car{Name: name, Distance: 0}, nil
}

func (c *Car) Move(strategy MovementStrategy) {
	if strategy.Movable() {
		c.Distance++
	}
}

func (c Car) Status() string {
	return fmt.Sprintf("%s : %s", c.Name, string(make([]rune, c.Distance, c.Distance)))
}
