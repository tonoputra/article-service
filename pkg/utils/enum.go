package utils

import "strconv"

func EnumString(m map[int8]string, v int8) string {
	s, ok := m[v]
	if ok {
		return s
	}
	return strconv.Itoa(int(v))
}
