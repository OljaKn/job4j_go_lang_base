package main

import (
	"job4j.ru/go-lang-base/internal/tracker"
)

func main() {
	ui := tracker.UI{
		In:      tracker.ConsoleInput{},
		Out:     tracker.ConsoleOutput{},
		Tracker: tracker.NewTracker(),
	}
	ui.Run()
	/*first := 100
	second := 10
	res := base.Max(first, second)
	fmt.Println(fmt.Sprintf("%d + %d = %d", first, second, res))*/

}
