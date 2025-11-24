package test

import (
	"goracing/domain"
	"goracing/tui"
	"testing"
)

func TestRaceIntegration(t *testing.T) {
	names := []string{"pobi", "crong", "honux", "woni", "jun"}
	cars, _ := domain.NewCars(names)

	strategy := domain.NewRandomMovementStrategy(1)

	// MoveAll과 Race 연동 테스트
	for i := 0; i < 5; i++ {
		results := cars.MoveAll(strategy)
		for _, r := range results {
			t.Logf("Round %d: %s 이동 %v, 거리 %d", i+1, r.Name, r.MovedThis, r.Distance)
		}
	}

	race := tui.NewRace(cars.List)
	updateCh, resultCh := race.Start()

	// 업데이트 채널 소비
	go func() {
		for range updateCh {
		}
	}()

	results := <-resultCh
	if len(results) != len(names) {
		t.Errorf("완주 차량 수 불일치: 기대 %d, 실제 %d", len(names), len(results))
	}

	for i, r := range results {
		t.Logf("%d위: %s, 거리: %d, 완료 시간: %.2fs", i+1, r.Name, r.Distance, r.Finish.Seconds())
	}
}
