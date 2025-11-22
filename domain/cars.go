package domain

import "fmt"

type Cars struct {
	List []*Car
}

func NewCars(names []string) (*Cars, error) {
	cars := make([]*Car, 0)
	for _, name := range names {
		car, err := NewCar(name)
		if err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}
	return &Cars{List: cars}, nil
}

func (c *Cars) MoveAll(strategy MovementStrategy) {
	for _, car := range c.List {
		car.Move(strategy)
	}
}

func (c *Cars) PrintStatus() {
	for _, car := range c.List {
		fmt.Println(car.Status())
	}
	fmt.Println()
}

func (c *Cars) Winners() []string {
	maxDist := 0
	for _, car := range c.List {
		if car.Distance > maxDist {
			maxDist = car.Distance
		}
	}

	winners := make([]string, 0)
	for _, car := range c.List {
		if car.Distance == maxDist {
			winners = append(winners, car.Name)
		}
	}
	return winners
}
