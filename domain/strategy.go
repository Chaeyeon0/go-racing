package domain

import "math/rand"

// 이동 전략
type MovementStrategy interface {
	Movable() bool
}

type RandomMovementStrategy struct {
	r *rand.Rand
}

func NewRandomMovementStrategy(seed int64) RandomMovementStrategy {
	source := rand.NewSource(seed)
	return RandomMovementStrategy{r: rand.New(source)}
}

func (r RandomMovementStrategy) Movable() bool {
	return r.r.Intn(10) >= 4
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
