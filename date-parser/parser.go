package parser

import (
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
)

type DateParser interface {
	Parse(text string) (*ParseResult, error)
}

type Parser struct {
	w *when.Parser
}

type ParseResult struct {
	Time        time.Time
	MatchedText string
}

func New() DateParser {
	w := when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)

	return Parser{
		w: w,
	}
}

func (p Parser) Parse(text string) (*ParseResult, error) {
	r, err := p.w.Parse(text, time.Now())
	if err != nil {
		return nil, err
	}
	if r != nil {
		return &ParseResult{Time: r.Time, MatchedText: r.Text}, nil
	}
	return nil, err
}
