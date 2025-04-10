package samplestructs

import "time"

type SampleStructA struct {
	Id   string
	Name string
	Date time.Time
}

func (s *SampleStructA) SampleFuncA() string {
	return "Hello"
}
