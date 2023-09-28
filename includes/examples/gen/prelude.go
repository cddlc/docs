package foo

import "time"

type Planet int

const (
	Mercury = iota
	Venus
	Earth
	Mars
)

type Location struct {
	Planet  Planet
	Country string

	FirstOpened time.Time
}
