package pipe_filter

import (
	"errors"
	"strings"
)

var SplitFilterWrongFormater = errors.New("input data must be string")

type SplitFilter struct {
	delimiter string
}

func NewSplitFilter(delimiter string) *SplitFilter {
	return &SplitFilter{delimiter}
}

func (sf *SplitFilter) Process(data Request) (Response, error) {
	str, ok := data.(string)
	if !ok {
		return nil, SplitFilterWrongFormater
	}
	parts := strings.Split(str, sf.delimiter)
	return parts, nil
}
