package tool

import (
	"time"
)

type timeSpend struct {
	lastDate time.Time
	tagSlice []any
}

func NewTimeSpend() *timeSpend {
	return &timeSpend{
		lastDate: time.Now(),
		tagSlice: make([]any, 0),
	}
}

func (t *timeSpend) Record(name string) {
	t.lastDate = time.Now()
	t.tagSlice = append(t.tagSlice, name, time.Now().Sub(t.lastDate))
}

func (t *timeSpend) End() []any {
	t.Record("end")
	return t.tagSlice
}

func (t *timeSpend) EndPrint() {
	t.Record("end")
	data := make([][]string, len(t.tagSlice)/2)
	for i := 0; i < len(t.tagSlice); i += 2 {
		data[i/2] = []string{t.tagSlice[i].(string), t.tagSlice[i+1].(time.Duration).String()}
	}
	PrettyPrint("time spend", data)
}
