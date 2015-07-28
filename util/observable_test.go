package util_test

import (
	"testing"

	"github.com/CotaPreco/Horus/util"
	"github.com/stretchr/testify/mock"
)

type Observable struct {
	util.Observable
}

type Observer struct {
	mock.Mock
}

func (m *Observer) Update(args ...interface{}) {
	m.Called(args)
}

func TestObservableNotifyAll(t *testing.T) {
	observer := new(Observer)
	observer.On("Update", mock.Anything).Return(nil)

	observable := &Observable{}
	observable.Attach(observer)

	observable.NotifyAll(observable)
	observable.NotifyAll(observable)
	observable.NotifyAll(observable)
	observable.NotifyAll(observable)

	observer.AssertNumberOfCalls(t, "Update", 4)
}
