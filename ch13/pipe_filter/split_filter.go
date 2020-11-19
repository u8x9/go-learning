package pipe_filter

import (
	"errors"
	"strings"
)

var SplitFilterWrongFormatError = errors.New("input data should be a string")

type SplitFilter struct {
	Delimiter string
}

func NewSplitFilter(delimiter string) *SplitFilter {
	return &SplitFilter{
		Delimiter: delimiter,
	}
}
func (sf *SplitFilter) Process(data Request) (Response, error) {
	str, ok := data.(string)
	if !ok {
		return nil, SplitFilterWrongFormatError
	}
	parts := strings.Split(str, sf.Delimiter)
	return parts, nil
}
