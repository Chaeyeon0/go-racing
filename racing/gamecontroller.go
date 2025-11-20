package racing

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func StartGame() {
	namesInput := getInput("경주할 자동차 이름을 입력하세요 (쉼표(,)로 구분): ")
	names := strings.Split(namesInput, ",")

	attemptInput := getInput("시도할 횟수는 몇 회인가요?: ")
	var attempts int
	fmt.Sscan(attemptInput, &attempts)

	cars, err := NewCars(names)
	if err != nil {
		fmt.Println("에러:", err)
		return
	}

	fmt.Println("\n실행 결과")
	for i := 0; i < attempts; i++ {
		cars.MoveAll(RandomMovementStrategy{})
		cars.PrintStatus()
	}

	winners := cars.Winners()
	fmt.Printf("🏁 최종 우승자: %s\n", strings.Join(winners, ", "))
}
