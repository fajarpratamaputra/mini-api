package common

import (
	"strconv"
)

type MessageEnum int32

const (
	Success      MessageEnum = 0
	ServerError  MessageEnum = 1
	InvalidInput MessageEnum = 2
	Unathorized  MessageEnum = 3
	NotFound     MessageEnum = 4
)

var MessageEnumMessageName = map[int32]string{
	0: "success",
	1: "Something went Wrong",
	2: "Invalid Input, check your parameter",
	3: "Unauthorized",
	4: "Category Not Found",
}

// String : return message string from enum
func (x MessageEnum) String() string {
	return enumToStr(MessageEnumMessageName, int32(x))
}

func (x MessageEnum) Int() int {
	return int(x)
}

func enumToStr(m map[int32]string, v int32) string {
	s, ok := m[v]
	if ok {
		return s
	}
	return strconv.Itoa(int(v))
}
