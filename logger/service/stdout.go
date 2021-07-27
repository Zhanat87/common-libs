package service

import "fmt"

type stdout struct{}

func NewStdout() Logger {
	return &stdout{}
}

func (s *stdout) Error(err error) {
	fmt.Printf("error: %#v\r\n", err)
}

func (s *stdout) Warning(warning string) {
	fmt.Printf("warning: %s\r\n", warning)
}

func (s *stdout) Info(info string) {
	fmt.Printf("info: %s\r\n", info)
}

func (s *stdout) Debug(debug string) {
	fmt.Printf("debug: %s\r\n", debug)
}
