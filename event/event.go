package event

type Event[T any] struct {
	funcs []T
}

func (e *Event[T]) Invoke() {
	for _, item := range e.funcs {
		item()
	}
}

func (e *Event[T]) Add(event T) {
	e.funcs = append(e.funcs, event)
}
