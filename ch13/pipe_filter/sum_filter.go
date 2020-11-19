package pipe_filter

import (
	"errors"
)

var SumFilterWrongFormatError = errors.New("input data should be a []int")

type SumFilter struct{}

func NewSumFilter() *SumFilter {
	return &SumFilter{}
}

func (sf *SumFilter) Process(data Request) (Response, error) {
	elems, ok := data.([]int)
	if !ok {
		return nil, SumFilterWrongFormatError
	}
	sum := 0
	for _, i := range elems {
		sum += i
	}
	return sum, nil
}
