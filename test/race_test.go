package test

import (
	"goracing/domain"
	"goracing/tui"
	"runtime"
	"testing"
	"time"
)

func TestRace_ConcurrentExecution(t *testing.T) {
	cars, _ := domain.NewCars([]string{"pobi", "woni", "jun", "chae", "go"})
	race := tui.NewRace(cars.List)

	startGoroutines := runtime.NumGoroutine()

	updateCh, resultCh := race.Start()

	go func() {
		for range updateCh {
			// 단순히 채널을 소비해서 에러 방지
		}
	}()

	select {
	case <-resultCh:
		endGoroutines := runtime.NumGoroutine()
		t.Logf("시작 시 고루틴 수: %d, 종료 후: %d", startGoroutines, endGoroutines)

		if endGoroutines <= startGoroutines {
			t.Error("고루틴이 병렬로 실행되지 않았습니다.")
		}
	case <-time.After(15 * time.Second):
		t.Error("타임아웃 — 병렬 실행 중 문제 발생")
	}
}
