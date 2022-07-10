package helper

import (
	"strconv"
	"strings"
)

type StringUtils interface {
	ArrToSliceString(str string) []string
	ArrToSliceInt(str string) []int
}

func NewString() StringUtils {
	return &stringUtils{}
}

type stringUtils struct {
}

func (*stringUtils) ArrToSliceString(str string) []string {
	sr := strings.ReplaceAll(str, "[", "")
	sr = strings.ReplaceAll(sr, "]", "")
	sr = strings.ReplaceAll(sr, "'", "")
	ss := strings.Split(sr, ",")

	var r []string

	for _, s := range ss {
		tr := strings.TrimSpace(s)
		r = append(r, tr)
	}

	return r
}

func (*stringUtils) ArrToSliceInt(str string) []int {
	sr := strings.ReplaceAll(str, "[", "")
	sr = strings.ReplaceAll(sr, "]", "")
	sr = strings.ReplaceAll(sr, "'", "")
	ss := strings.Split(sr, ",")

	var r []int

	for _, s := range ss {
		tr := strings.TrimSpace(s)
		in, _ := strconv.Atoi(tr)
		r = append(r, in)
	}

	return r
}
