package tui

import (
	"goracing/domain"
	"math/rand"
	"sort"
	"time"
)

type Race struct {
	Cars []*domain.Car
}

type RaceResult struct {
	Name     string
	Distance int
	Finish   time.Duration
}

func NewRace(cars []*domain.Car) *Race {
	return &Race{Cars: cars}
}

func (r *Race) Start() (<-chan []*domain.Car, <-chan []RaceResult) {
	rand.Seed(time.Now().UnixNano())
	updateCh := make(chan []*domain.Car)
	resultCh := make(chan []RaceResult, 1)
	start := time.Now()

	go func() {
		defer close(updateCh)
		defer close(resultCh)

		var results []RaceResult
		done := false

		for !done {
			for _, car := range r.Cars {
				if car.Distance < 30 && rand.Intn(10) > 3 {
					car.Distance++
					if car.Distance == 30 {
						results = append(results, RaceResult{
							Name:     car.Name,
							Distance: 30,
							Finish:   time.Since(start),
						})
					}
				}
			}

			temp := make([]*domain.Car, len(r.Cars))
			copy(temp, r.Cars)
			updateCh <- temp

			done = true
			for _, car := range r.Cars {
				if car.Distance < 30 {
					done = false
					break
				}
			}
			time.Sleep(100 * time.Millisecond)
		}

		sort.Slice(results, func(i, j int) bool {
			return results[i].Finish < results[j].Finish
		})
		resultCh <- results
	}()

	return updateCh, resultCh
}
