package task

import (
	"bytes"
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/adrg/xdg"
)

const fileName = "data"

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

func (t Task) IsDue() bool {
	return t.DueDate.Before(time.Now())
}

func (t Task) IsOld() bool {
	return t.DueDate.Before(time.Now().AddDate(0, 0, -1))
}

func (t Task) WriteToFile() error {
	encoded, err := t.Encode()
	if err != nil {
		return err
	}

	err = createFolder()
	if err != nil {
		return err
	}

	f, err := os.OpenFile(GetFileLocation()+fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err = f.WriteString(encoded); err != nil {
		return err
	}
	return nil
}

func LoadTasks() ([]Task, error) {
	createFolder()
	f, err := os.Open(GetFileLocation() + fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)

	records, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	tasks := make([]Task, len(records))
	for i, record := range records {
		notified, err := strconv.ParseBool(record[1])
		if err != nil {
			return nil, err
		}

		dueDate, err := time.Parse("2006-01-02 15:04:05", record[2])
		if err != nil {
			return nil, err
		}

		task := Task{
			Text:     record[0],
			Notified: notified,
			DueDate:  dueDate,
		}
		tasks[i] = task
	}

	return tasks, nil
}

func ClearSavedTasks() error {
	err := createFolder()
	if err != nil {
		return err
	}

	_, err = os.OpenFile(GetFileLocation()+fileName, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0600)
	return err
}

func FromString(text string) (*Task, error) {
	r := strings.NewReader(text)
	reader := csv.NewReader(r)

	record, err := reader.Read()
	if err != nil {
		return nil, err
	}

	notified, err := strconv.ParseBool(record[1])
	if err != nil {
		return nil, err
	}

	dueDate, err := time.Parse("2006-01-02 15:04:05", record[2])
	if err != nil {
		return nil, err
	}

	task := Task{
		Text:     record[0],
		Notified: notified,
		DueDate:  dueDate,
	}

	return &task, nil
}

func GetFileLocation() string {
	dataHome := xdg.DataHome
	return dataHome + "/godo/"
}

func createFolder() error {
	err := os.MkdirAll(GetFileLocation(), 0755)
	if err != nil {
		return err
	}
	return nil
}
