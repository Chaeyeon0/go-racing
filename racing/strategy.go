package racing

import "math/rand"

// 이동 전략
type MovementStrategy interface {
	Movable() bool
}

// 랜덤 전략
type RandomMovementStrategy struct{}

func (r RandomMovementStrategy) Movable() bool {
	return rand.Intn(10) >= 4
}

// 테스트용 전략
type AlwaysMoveStrategy struct{}

func (a AlwaysMoveStrategy) Movable() bool {
	return true
}

type NeverMoveStrategy struct{}

func (n NeverMoveStrategy) Movable() bool {
	return false
}
