package samplestructs

import "time"

type SampleStructA struct {
	Id   string    `json:"id"`
	Name string    `json:"name"`
	Date time.Time `json:"date"`
}

func (s *SampleStructA) SampleFuncA() string {
	return "Hello"
}
