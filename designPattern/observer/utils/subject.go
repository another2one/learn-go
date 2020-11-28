package utils

type Subject struct {
	status   int
	Observer *Observer
}

func (subject *Subject) SetStatus(i int) {
	subject.status = i
	subject.Observer.Notify(i)
}
