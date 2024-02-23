package task

import (
	"bytes"
	"encoding/csv"
	"io"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	Text     string
	Notified bool
	DueDate  time.Time
}

func (t Task) Encode() (string, error) {
	var buf bytes.Buffer
	w := io.MultiWriter(&buf)
	writer := csv.NewWriter(w)

	record := []string{t.Text, strconv.FormatBool(t.Notified), t.DueDate.Format("2006-01-02 15:04:05")}

	err := writer.Write(record)
	if err != nil {
		return "", err
	}

	writer.Flush()

	return buf.String(), nil
}

func FromString(text string) (*Task, error) {
	r := strings.NewReader(text)
	reader := csv.NewReader(r)

	record, err := reader.Read()
	if err != nil {
		return nil, err
	}

	notified, err := strconv.ParseBool(record[1])
	dueDate, err := time.Parse("2006-01-02 15:04:05", record[2])
	task := Task{
		Text:     record[0],
		Notified: notified,
		DueDate:  dueDate,
	}

	return &task, nil
}
