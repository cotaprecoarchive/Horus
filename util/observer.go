package util

type Observer interface {
	Update(args ...interface{})
}
