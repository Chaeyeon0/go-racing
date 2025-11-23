package domain

import "fmt"

type Cars struct {
	List []*Car
}

type MoveResult struct {
	Name      string
	Distance  int
	MovedThis bool
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

func (c *Cars) MoveAll(strategy MovementStrategy) []MoveResult {
	results := make([]MoveResult, len(c.List))
	for i, car := range c.List {
		before := car.Distance
		car.Move(strategy)
		results[i] = MoveResult{
			Name:      car.Name,
			Distance:  car.Distance,
			MovedThis: car.Distance > before,
		}
	}
	return results
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

func (c *Cars) StatusList() []string {
	statuses := make([]string, len(c.List))
	for i, car := range c.List {
		statuses[i] = car.Status()
	}
	return statuses
}
