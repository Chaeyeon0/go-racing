package tui

import (
	"fmt"
	"goracing/domain"
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

	for update := range updateCh {
		ui.TextView.Clear()
		for _, car := range update {
			bar := string(make([]rune, car.Distance))
			fmt.Fprintf(ui.TextView, "[white]%-5s: [green]%s>\n", car.Name, bar)
		}
		time.Sleep(100 * time.Millisecond)
	}

	results := <-resultCh
	fmt.Fprintf(ui.TextView, "\n🏁 [green::b]모든 차량 완주!\n\n")
	for i, r := range results {
		fmt.Fprintf(ui.TextView, "[white]%d위: 🚗 %s — %.2fs\n", i+1, r.Name, r.Finish.Seconds())
	}

	ui.App.Stop()
}
