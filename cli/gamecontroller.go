package cli

import (
	"fmt"
	"goracing/domain"
)

func StartGame() {
	names, attempts, err := ReadInputs()
	if err != nil {
		fmt.Println("입력 에러:", err)
		return
	}

	cars, err := domain.NewCars(names)
	if err != nil {
		fmt.Println("자동차 생성 에러:", err)
		return
	}

	RunGame(cars, attempts)
}

// ReadInputs 입력 처리 통합
func ReadInputs() ([]string, int, error) {
	names, err := ReadCarNames()
	if err != nil {
		return nil, 0, err
	}

	attempts, err := ReadAttemptCount()
	if err != nil {
		return nil, 0, err
	}

	return names, attempts, nil
}

// RunGame 게임 실행 및 결과 출력
func RunGame(cars *domain.Cars, attempts int) {
	fmt.Println("\n실행 결과")

	for i := 0; i < attempts; i++ {
		cars.MoveAll(domain.RandomMovementStrategy{})
		PrintRound(cars)
	}

	winners := cars.Winners()
	PrintFinalWinners(winners)
}

// PrintRound 라운드별 상태 출력
func PrintRound(cars *domain.Cars) {
	statuses := cars.StatusList() // Cars에서 문자열 리스트 반환
	for _, s := range statuses {
		fmt.Println(s)
	}
	fmt.Println()
}
