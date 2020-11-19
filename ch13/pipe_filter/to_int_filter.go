package pipe_filter

import (
	"errors"
	"strconv"
)

var ToIntFilterWrongFormatError = errors.New("input data should be a []string")

type ToIntFilter struct{}

func NewToIntFilter() *ToIntFilter {
	return new(ToIntFilter)
}

func (tif *ToIntFilter) Process(data Request) (Response, error) {
	parts, ok := data.([]string)
	if !ok {
		return nil, ToIntFilterWrongFormatError
	}
	ret := make([]int, 0, len(parts))
	for _, part := range parts {
		s, err := strconv.Atoi(part)
		if err != nil {
			continue
		}
		ret = append(ret, s)
	}
	return ret, nil

}
