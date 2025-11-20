package racing

import (
	"fmt"
	"strings"
)

func PrintRoundResult(cars *Cars) {
	for _, car := range cars.list {
		fmt.Println(car.Status())
	}
	fmt.Println()
}

func PrintFinalWinners(winners []string) {
	fmt.Printf("🏁 최종 우승자: %s\n", strings.Join(winners, ", "))
}
