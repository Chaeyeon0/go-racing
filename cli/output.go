package cli

import (
	"fmt"
	"goracing/domain"
	"strings"
)

func PrintRoundResult(cars *domain.Cars) {
	for _, car := range cars.List {
		fmt.Println(car.Status())
	}
	fmt.Println()
}

func PrintFinalWinners(winners []string) {
	fmt.Printf("🏁 최종 우승자: %s\n", strings.Join(winners, ", "))
}
