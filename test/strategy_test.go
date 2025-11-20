package test

import (
	"goracing/racing"
	"math/rand"
	"testing"
)

func TestAlwaysMoveStrategy(t *testing.T) {
	strategy := racing.AlwaysMoveStrategy{}
	if !strategy.Movable() {
		t.Error("AlwaysMoveStrategy는 항상 true여야 합니다.")
	}
}

func TestNeverMoveStrategy(t *testing.T) {
	strategy := racing.NeverMoveStrategy{}
	if strategy.Movable() {
		t.Error("NeverMoveStrategy는 항상 false여야 합니다.")
	}
}

func TestRandomMovementStrategy(t *testing.T) {
	rand.Seed(1) // 결과 예측 가능하게 고정
	strategy := racing.RandomMovementStrategy{}
	result := strategy.Movable()
	if result != (rand.Intn(10) >= 4) {
		t.Error("RandomMovementStrategy의 이동 조건이 올바르지 않습니다.")
	}
}
