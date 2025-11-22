package test

import (
	"goracing/domain"
	"testing"
)

func TestCar_MoveWithAlwaysMoveStrategy(t *testing.T) {
	car, _ := domain.NewCar("pobi")
	car.Move(domain.AlwaysMoveStrategy{})

	if car.Distance != 1 {
		t.Errorf("전진 실패: Distance = %d, 기대값 = 1", car.Distance)
	}
}

func TestCar_MoveWithNeverMoveStrategy(t *testing.T) {
	car, _ := domain.NewCar("crong")
	car.Move(domain.NeverMoveStrategy{})

	if car.Distance != 0 {
		t.Errorf("멈춤 실패: Distance = %d, 기대값 = 0", car.Distance)
	}
}

func TestCar_NameTooLong(t *testing.T) {
	_, err := domain.NewCar("abcdef")
	if err == nil {
		t.Error("5자 초과 이름에 대해 에러가 발생해야 합니다.")
	}
}
