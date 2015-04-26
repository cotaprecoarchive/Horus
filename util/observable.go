package util

type Observable []Observer

func (observers *Observable) Attach(observer Observer) {
	*observers = append(*observers, observer)
}

func (observers Observable) NotifyAll(subject interface{}) {
	for _, observer := range observers {
		observer.Update(subject)
	}
}
