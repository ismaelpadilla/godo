package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/adrg/xdg"
	"github.com/gen2brain/beeep"
	parser "github.com/ismaelpadilla/godo/date-parser"
	"github.com/ismaelpadilla/godo/task"
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

	encoded, err := t.Encode()
	decoded, err := task.FromString(encoded)
	reencoded, err := decoded.Encode()
	fmt.Printf("encoded: %v\n", encoded)
	fmt.Printf("decoded: %v\n", decoded)
	fmt.Printf("reencoded: %v\n", reencoded)
}

func matchesExactly(text, matchedResult string) bool {
	return text == matchedResult
}

func notify(text string) error {
	return beeep.Notify("godo reminder", text, "")
}
