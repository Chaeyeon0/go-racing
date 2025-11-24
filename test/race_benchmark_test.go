package test

import (
	"goracing/domain"
	"goracing/tui"
	"runtime"
	"testing"
	"time"
)

func BenchmarkRacePerformance(b *testing.B) {
	testCases := []int{10, 100, 1000, 10000}

	for _, carCount := range testCases {
		b.Run(
			// 서브테스트 이름
			string(rune(carCount)),
			func(b *testing.B) {
				names := make([]string, carCount)
				for i := range names {
					names[i] = "car" + string(rune('A'+i%26))
				}

				cars, _ := domain.NewCars(names)
				race := tui.NewRace(cars.List)

				start := time.Now()
				updateCh, resultCh := race.Start()

				// 실시간 업데이트 소비
				go func() {
					for range updateCh {
					}
				}()

				<-resultCh
				elapsed := time.Since(start)

				b.ReportMetric(float64(runtime.NumGoroutine()), "goroutines")
				b.ReportMetric(float64(elapsed.Milliseconds()), "ms")
			})
	}
}
