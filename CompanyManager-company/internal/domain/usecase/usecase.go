package usecase

import (
	"strconv"
)

func  MessageService(message string) int64 {
	id, err := strconv.Atoi(message)
	if err != nil {
		panic(err)
	}
	return int64(id)
}
