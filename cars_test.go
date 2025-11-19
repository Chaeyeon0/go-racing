package goracing

import (
	"reflect"
	"testing"
)

func TestCars_MoveAll(t *testing.T) {
	cars, _ := NewCars([]string{"pobi", "crong"})
	cars.MoveAll(AlwaysMoveStrategy{})

	if cars.Winners()[0] != "pobi" {
		t.Errorf("이동 실패 — 기대값: pobi, 결과값: %v", cars.Winners())
	}
}

func TestCars_Winners(t *testing.T) {
	cars, _ := NewCars([]string{"pobi", "crong", "honux"})
	cars.MoveAll(AlwaysMoveStrategy{})
	cars.MoveAll(AlwaysMoveStrategy{}) // 2번 이동

	expected := []string{"pobi", "crong", "honux"}
	if !reflect.DeepEqual(cars.Winners(), expected) {
		t.Errorf("우승자 계산 오류 — 기대값: %v, 결과값: %v", expected, cars.Winners())
	}
}
