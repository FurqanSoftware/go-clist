package clist

import "strconv"

type Error struct {
	StatusCode int
}

func (e Error) Error() string {
	return "clist: status " + strconv.Itoa(e.StatusCode)
}
