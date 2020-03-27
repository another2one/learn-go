package demo02

import "errors"

// 人员档案
type Profile struct {
	Name    string // 名字
	Age     int    // 年龄
	Married bool   // 已婚
}

type HashMap struct {
	Data map[int][]*Profile
}

func (h *HashMap) getKey(name string, age int) int {
	sum := 0
	for _, val := range []byte(name) {
		sum += int(val)
	}
	return sum + age*1000000
}

func (h *HashMap) BuildIndex(pros []*Profile) {
	for _, val := range pros {
		key := h.getKey(val.Name, val.Age)
		h.Data[key] = append(h.Data[key], val)
	}
}

func (h *HashMap) QueryData(name string, age int) (data *Profile, err error) {
	res, ok := h.Data[h.getKey(name, age)]
	if ok {
		for _, val := range res {
			if val.Name == name && val.Age == age {
				return val, nil
			}
		}
	}
	return nil, errors.New("not found")
}
