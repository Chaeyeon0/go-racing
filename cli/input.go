package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadCarNames() ([]string, error) {
	fmt.Print("경주할 자동차 이름을 입력하세요 (쉼표(,)로 구분): ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	names := strings.Split(input, ",")
	for _, name := range names {
		if len(name) > 5 {
			return nil, fmt.Errorf("자동차 이름은 5자 이하만 가능합니다: %s", name)
		}
	}
	return names, nil
}

func ReadAttemptCount() (int, error) {
	fmt.Print("시도할 횟수는 몇 회인가요?: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	var attempts int
	_, err := fmt.Sscan(input, &attempts)
	if err != nil || attempts <= 0 {
		return 0, fmt.Errorf("유효한 숫자를 입력해주세요.")
	}
	return attempts, nil
}
