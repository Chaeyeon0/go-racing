package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadCarNames() ([]string, error) {
	fmt.Print("경주할 자동차 이름을 입력하세요 (쉼표(,)로 구분): ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	names := strings.Split(input, ",")
	for _, name := range names {
		name = strings.TrimSpace(name)
		if len(name) == 0 {
			return nil, fmt.Errorf("빈 이름은 허용되지 않습니다.")
		}
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

	if input == "" {
		return 0, fmt.Errorf("시도 횟수를 입력해야 합니다.")
	}

	attempts, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("숫자만 입력해야 합니다.")
	}
	if attempts <= 0 {
		return 0, fmt.Errorf("시도 횟수는 1 이상이어야 합니다.")
	}
	return attempts, nil
}
