package tui

import (
	"goracing/domain"
	"math/rand"
	"sort"
	"sync"
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

	var wg sync.WaitGroup
	results := make([]RaceResult, 0)
	startTime := time.Now()

	// 각 자동차를 병렬로 실행
	for _, car := range r.Cars {
		wg.Add(1)
		go func(c *domain.Car) {
			defer wg.Done()

			for c.Distance < 30 {
				// 랜덤 전진 (고루틴별 독립 동작)
				if rand.Intn(10) >= 4 {
					c.Distance++
				}
				// 각 상태를 브로드캐스트용 임시 복사로 전송
				snapshot := make([]*domain.Car, len(r.Cars))
				copy(snapshot, r.Cars)
				updateCh <- snapshot
				time.Sleep(100 * time.Millisecond)
			}

			// 완주 시 기록 저장
			results = append(results, RaceResult{
				Name:     c.Name,
				Distance: c.Distance,
				Finish:   time.Since(startTime),
			})
		}(car)
	}

	// 모든 자동차 완주 후 결과 전송
	go func() {
		wg.Wait()
		close(updateCh)

		// 완주 시간 순으로 정렬
		sort.Slice(results, func(i, j int) bool {
			return results[i].Finish < results[j].Finish
		})
		resultCh <- results
		close(resultCh)
	}()

	return updateCh, resultCh
}
