package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gen2brain/beeep"
	parser "github.com/ismaelpadilla/godo/date-parser"
	"github.com/ismaelpadilla/godo/task"
)

func main() {
	input := strings.Join(os.Args[1:], " ")

	if input == "--check" {
		err := checkAndNotify()
		if err != nil {
			panic(err)
		}
		return
	}
	separated := strings.Split(input, "@@")

	if len(separated) != 2 {
		fmt.Println("Correct usage: \"godo Remind me to do something @@ tomorrow\"")
		return
	}

	reminderText := strings.TrimSpace(separated[0])
	reminderTime := strings.TrimSpace(separated[1])

	p := parser.New()
	r, err := p.Parse(reminderTime)
	if err != nil {
		panic(err)
	}
	if r == nil {
		fmt.Println("Couldn't parse date")
		return
	}

	if !matchesExactly(reminderTime, r.MatchedText) {
		fmt.Println("Text doesn't exactly match")
	}

	t := task.Task{
		Text:     reminderText,
		Notified: false,
		DueDate:  r.Time,
	}

	err = t.WriteToFile()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Task due on %s: %s", t.DueDate, t.Text)
}

func checkAndNotify() error {
	tasks, err := task.LoadTasks()
	if err != nil {
		panic(err)
	}

	err = task.ClearSavedTasks()
	if err != nil {
		panic(err)
	}
	for _, t := range tasks {
		if t.IsDue() && !t.Notified {
			err := notify(t.Text)
			if err != nil {
				panic(err)
			}
			t.Notified = true
		}

		if !t.IsOld() {
			t.WriteToFile()
		}
	}

	return nil
}

func matchesExactly(text, matchedResult string) bool {
	return text == matchedResult
}

func notify(text string) error {
	return beeep.Notify("godo reminder", text, "")
}
