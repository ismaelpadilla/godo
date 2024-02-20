package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/adrg/xdg"
	"github.com/gen2brain/beeep"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
)

func main() {
	dataHome := xdg.DataHome
	fmt.Println(dataHome)
	input := strings.Join(os.Args[1:], " ")
	separated := strings.Split(input, "@@")

	fmt.Printf("input: %v\n", input)

	if len(separated) != 2 {
		fmt.Println("Correct usage: \"godo Remind me to do something @@ tomorrow\"")
		return
	}

	reminderText := separated[0]
	reminderTime := strings.TrimSpace(separated[1])

	r, err := parse(reminderTime)
	if err != nil {
		panic(err)
	}
	if r == nil {
		fmt.Println("Couldn't parse date")
		return
	}

	if !matchesExactly(reminderTime, r) {
		fmt.Println("Text doesn't exactly match")
	}

	fmt.Printf("reminderText: %v\n", reminderText)
	fmt.Printf("r.Time: %v\n", r.Time)
}

func matchesExactly(text string, result *when.Result) bool {
	return text == result.Text
}

func notify(text string) error {
	return beeep.Notify("godo reminder", text, "")
}

func parse(text string) (*when.Result, error) {
	w := when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)

	return w.Parse(text, time.Now())
}
