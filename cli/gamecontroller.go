package cli

import (
	"fmt"
	"goracing/domain"
)

func StartGame() {
	names, err := ReadCarNames()
	if err != nil {
		fmt.Println("입력 에러:", err)
		return
	}

	attempts, err := ReadAttemptCount()
	if err != nil {
		fmt.Println("입력 에러:", err)
		return
	}

	cars, _ := domain.NewCars(names)
	fmt.Println("\n실행 결과")

	for i := 0; i < attempts; i++ {
		cars.MoveAll(domain.RandomMovementStrategy{})
		PrintRoundResult(cars)
	}

	winners := cars.Winners()
	PrintFinalWinners(winners)
}
