# 🚗 Go 고루틴 자동차 경주

> Go 언어의 고루틴(Goroutine) 과 객체지향 설계(Struct + Interface) 를 결합한
> 
> 
> 병렬 자동차 경주 게임입니다.
> 
> 각 자동차는 랜덤한 속도로 이동하며, 모든 차량이 완주할 때까지 순위를 실시간으로 갱신합니다.
> 

---

## 🗂️ 패키지 구조

```bash
goracing/
├── go.mod
├── main.go                  # 프로그램 진입점 (CLI / TUI 모드 선택)
├── cli/                     # CLI 모드 관련
│   ├── startgame.go         # 게임 컨트롤러 (입력 → 실행 → 출력)
│   ├── input.go             # 사용자 입력 처리 (자동차 이름, 시도 횟수)
│   └── output.go            # 라운드별 출력, 최종 우승자 출력
├── domain/                  # 도메인 모델
│   ├── car.go               # 자동차 구조체 및 상태 출력
│   ├── cars.go              # 자동차 그룹 관리, MoveAll, Winners
│   └── strategy.go          # 이동 전략 인터페이스 및 구현 (Random, AlwaysMove, NeverMove)
├── tui/                     # TUI (실시간 UI) 관련
│   ├── race.go              # 병렬 레이스 구현 (고루틴, 채널)
│   └── race_ui.go           # tview 기반 UI 구현
└── test/                    # 단위 및 성능 테스트
    ├── car_test.go          # Car 단위 테스트
    ├── cars_test.go         # Cars 단위 테스트
    ├── strategy_test.go     # 이동 전략 테스트
    ├── race_test.go         # Race 동시성 테스트
    └── benchmark_race.go    # 대규모 레이스 성능 테스트
```

---

## 🧩 파일별 역할 요약

| 파일명 | 패키지 | 주요 역할 |
| --- | --- | --- |
| `car.go` | `domain` | 자동차 이름, 이동 거리, 전진 로직 구현 |
| `cars.go` | `domain` | 자동차 리스트 관리, 우승자 판별 |
| `strategy.go` | `domain` | 이동 조건 전략 인터페이스 및 구현체 정의 |
| `startgame.go` | `cli` | CLI 기반 게임 로직 제어 (입력 → 시도 → 결과) |
| `input.go` | `cli` | 사용자 입력 처리 (자동차 이름, 시도 횟수 등) |
| `output.go` | `cli` | 각 라운드별 상태 및 최종 우승자 출력 |
| `race.go` | `tui` | 고루틴 병렬 레이스, 실시간 상태 채널 관리 |
| `race_ui.go` | `tui` | tview 기반 실시간 UI 시각화 (자동차 이동, 순위 표시) |
| `main.go` | root | CLI / TUI 실행 분기 및 진입점 |
| `test/` | test | 단위, 동시성, 성능 테스트 |

---

## ✅ 기능 요구 사항 체크리스트

### 1️⃣ 자동차 기본 동작 (`car.go`)

- [x]  자동차는 이름(name)과 이동 거리(distance)를 가진다.
- [x]  이름은 5자 이하만 허용하며 초과 시 `error` 발생
- [x]  0~9 사이 랜덤 값 중 **4 이상이면 전진**, 아니면 정지
- [x]  전진 여부를 `MovementStrategy` 인터페이스로 분리
- [x]  자동차 상태 문자열 반환(`String()` or `Status()`)

---

### 2️⃣ 자동차 그룹 관리 (`cars.go`)

- [x]  여러 자동차(`[]Car`)를 관리하는 `Cars` 구조체 정의
- [x]  모든 자동차 전진 (`MoveAll(strategy)`) 기능
- [x]  각 자동차의 이름과 이동 거리 상태 출력
- [x]  최대 거리 기준 **우승자 판별 (복수 가능)**

---

### 3️⃣ 이동 전략 (`strategy.go`)

- [x]  `MovementStrategy` 인터페이스 정의 (`Movable() bool`)
- [x]  `RandomMovementStrategy` 구현 (랜덤 값 ≥ 4 → 전진)
- [x]  테스트용 전략 (`AlwaysMoveStrategy`, `NeverMoveStrategy`) 추가

---

### 4️⃣ 입력 / 출력 (`input.go`, `output.go`)

### InputView

- [x]  자동차 이름 입력 (`쉼표(,)`로 구분)
- [x]  시도 횟수 입력
- [x]  잘못된 입력(이름 초과, 음수, 비숫자) → `error`

### OutputView

- [x]  각 라운드별 결과 출력
- [x]  최종 우승자 출력
- [x]  `tview`로 시각화 (색상, 순위 표시)

---

### 5️⃣ 게임 컨트롤러 (`controller.go`)

- [x]  전체 흐름 제어 (입력 → 실행 → 출력)
- [x]  라운드별 전진 및 상태 출력
- [x]  모든 시도 후 우승자 계산 및 출력

---

### 6️⃣ 병렬 경주 (`race.go`)

- [x]  `Race` 구조체에서 자동차별 **고루틴 실행**
- [x]  각 자동차의 **완주 시간(time.Duration)** 기록
- [x]  완료 순서 정렬 후 **순위 계산**
- [x]  실시간 진행 상태 `update` 채널로 송신
- [x]  100~1000대 확장 성능 테스트

---

### 7️⃣ 애플리케이션 진입점 (`main.go`)

- [x]  프로그램 실행 및 예외 처리
- [x]  `GameController` 초기화
- [x]  완주 후 순위 출력

---

### 8️⃣ 테스트 (`_test.go`)

- [x]  `Car` 이름 검증
- [x]  이동 조건 테스트 (랜덤 / AlwaysMove)
- [x]  `Cars` 우승자 계산
- [x]  `Race` 동시성 테스트
- [x]  대규모 레이스 성능 테스트

---

## ⚙️ 구현 순서 (TDD 기반)

| 단계 | 구현 대상 | 테스트 포인트 |
| --- | --- | --- |
| 1️⃣ | `car.go` | 이름 길이, 이동 조건 테스트 |
| 2️⃣ | `strategy.go` | 전략 패턴 동작 테스트 |
| 3️⃣ | `cars.go` | MoveAll + 우승자 계산 |
| 4️⃣ | `controller.go` | 전체 게임 흐름 테스트 |
| 5️⃣ | `race.go` | 고루틴 병렬 처리 결과 테스트 |
| 6️⃣ | `ui.go` | tview 시각화 및 실시간 순위 확인 |

---

## 🎯 목표

- Go의 **객체지향 설계에 대해서 고민하고 익히기**
- **TDD 기반 개발 프로세스** 연습
- **병렬성(Goroutine, Channel)** 이해 및 실습하기

---

## 🏁 실행 예시

### 1️⃣ CLI 모드

```bash
$ go run cmd/main.go --mode=cli
경주할 자동차 이름을 입력하세요 (쉼표로 구분): pobi, crong, honux
시도할 횟수는 몇 회인가요?: 5

실행 결과
pobi   : >>>>
crong  : >>>
honux  : >>>>

pobi   : >>>>>>
crong  : >>>>>
honux  : >>>>>>

🏁 최종 우승자: pobi, honux
```

### 2️⃣ TUI 모드

```bash
$ go run cmd/main.go --mode=tui
🚗 고루틴 자동차 경주 시작!

 1위: pobi    >>>>>>> 🏁 완료!
 2위: honux   >>>>>>>
 3위: crong   >>>>>>>

🏆 모든 차량 완주!
1위: 🚗 pobi — 3.57s
2위: 🚗 honux — 3.66s
3위: 🚗 crong — 3.83s
```

- TUI에서는 **실시간 순위 반영**
- **색상**으로 진행 중/완주 구분
- 완료한 자동차는 **🏁 표시**
