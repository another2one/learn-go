package demo02

import "errors"

type StructMap struct {
	Data map[queryKey][]*Profile
}

type queryKey struct {
	Name string
	Age  int
}

func NewStructMap() *StructMap {
	return &StructMap{Data: make(map[queryKey][]*Profile)}
}

func (h *StructMap) getKey(name string, age int) queryKey {
	return queryKey{name, age}
}

func (h *StructMap) BuildIndex(pros []*Profile) {
	for _, val := range pros {
		key := h.getKey(val.Name, val.Age)
		h.Data[key] = append(h.Data[key], val)
	}
}

func (h *StructMap) QueryData(name string, age int) (data []*Profile, err error) {
	res, ok := h.Data[h.getKey(name, age)]
	if ok {
		return res, nil
	}
	return nil, errors.New("not found")
}
