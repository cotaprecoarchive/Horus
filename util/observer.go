package util

type Observer interface {
	Update(subject interface{})
}
