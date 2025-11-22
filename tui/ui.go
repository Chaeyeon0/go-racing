package tui

import (
	"fmt"
	"goracing/domain"
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
		for update := range updateCh {
			ui.App.QueueUpdateDraw(func() {
				ui.TextView.Clear()
				for _, car := range update {
					bar := strings.Repeat(">", car.Distance)
					fmt.Fprintf(ui.TextView, "[white]%-5s: [green]%s>\n", car.Name, bar)
				}
			})
		}

		results := <-resultCh
		ui.App.QueueUpdateDraw(func() {
			fmt.Fprintf(ui.TextView, "\n🏁 [green::b]모든 차량 완주!\n\n")
			for i, r := range results {
				fmt.Fprintf(ui.TextView, "[white]%d위: 🚗 %s — %.2fs\n", i+1, r.Name, r.Finish.Seconds())
			}
		})

		time.Sleep(3 * time.Second)
		ui.App.Stop()
	}()

	if err := ui.App.SetRoot(ui.TextView, true).Run(); err != nil {
		panic(err)
	}
}
