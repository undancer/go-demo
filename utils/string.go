package utils

import "bytes"

func BytesToString(bs *[]byte) *string {
	s := bytes.NewBuffer(*bs)
	r := s.String()
	return &r
}
