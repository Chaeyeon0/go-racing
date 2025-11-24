package main

import (
	"flag"
	"fmt"
	"goracing/cli"
	"goracing/domain"
	"goracing/tui"

	"github.com/rivo/tview"
)

func main() {
	mode := flag.String("mode", "cli", "실행 모드: cli 또는 tui")
	flag.Parse()

	switch *mode {
	case "cli":
		cli.StartGame()
	case "tui":
		app := tview.NewApplication()
		names := []string{"pobi", "woni", "jun", "chae", "go"}
		cars, _ := domain.NewCars(names)

		ui := tui.NewRaceUI(app, cars.List)
		go ui.Start()

		if err := app.SetRoot(ui.TextView, true).Run(); err != nil {
			panic(err)
		}
	default:
		fmt.Println("지원하지 않는 모드입니다. --mode=cli 또는 --mode=tui 사용하세요.")
	}
}
