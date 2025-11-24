package tui

import (
	"fmt"
	"goracing/domain"
	"sort"
	"strings"
	"time"

	"github.com/rivo/tview"
)

type RaceUI struct {
	App      *tview.Application
	TextView *tview.TextView
	Race     *Race
}

func NewRaceUI(app *tview.Application, cars []*domain.Car) *RaceUI {
	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetChangedFunc(func() { app.Draw() })

	return &RaceUI{
		App:      app,
		TextView: textView,
		Race:     NewRace(cars),
	}
}

func (ui *RaceUI) Start() {
	fmt.Fprintf(ui.TextView, "[yellow::b]🚗 고루틴 자동차 경주 시작!\n\n")
	updateCh, resultCh := ui.Race.Start()

	go func() {
		lastUpdate := time.Now()
		for update := range updateCh {

			if time.Since(lastUpdate) < 100*time.Millisecond {
				continue
			}
			lastUpdate = time.Now()

			ui.App.QueueUpdateDraw(func() {
				ui.TextView.Clear()

				// 진행 상황 순으로 정렬 (거리 기준)
				sorted := make([]*domain.Car, len(update))
				copy(sorted, update)
				sort.Slice(sorted, func(i, j int) bool {
					return sorted[i].Distance > sorted[j].Distance
				})

				for i, car := range sorted {
					rank := i + 1
					bar := strings.Repeat(">", car.Distance)

					// 색상: 진행중 초록, 완주 노랑
					color := "[green]"
					status := ""
					if car.Distance >= 30 {
						color = "[yellow]"
						status = " 🏁 완료!"
					}

					fmt.Fprintf(ui.TextView, "[white]%2d위: %-5s %s%s%s\n",
						rank, car.Name, color, bar, status)
				}
			})
		}

		// 결과 채널에서 최종 완주 순위 받기
		results := <-resultCh
		ui.App.QueueUpdateDraw(func() {
			fmt.Fprintf(ui.TextView, "\n🏆 [green::b]모든 차량 완주!\n\n")
			for i, r := range results {
				fmt.Fprintf(ui.TextView, "[white]%d위: 🚗 %s — %.2fs\n",
					i+1, r.Name, r.Finish.Seconds())
			}
		})

		time.Sleep(3 * time.Second)
		ui.App.Stop()
	}()

	if err := ui.App.SetRoot(ui.TextView, true).Run(); err != nil {
		panic(err)
	}
}
