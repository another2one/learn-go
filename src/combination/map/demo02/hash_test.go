package demo02

import "testing"

func getList() []*Profile {
	return []*Profile{
		{Name: "张三", Age: 30, Married: true},
		{Name: "李四", Age: 21},
		{Name: "李四", Age: 21},
		{Name: "李四", Age: 21},
		{Name: "王麻子", Age: 21},
	}
}

func TestHash(t *testing.T) {
	hash := &HashMap{
		Data: make(map[int][]*Profile),
	}
	hash.BuildIndex(getList())
	res, err := hash.QueryData("李四", 21)
	if err != nil {
		t.Error("TestHash query fail:", err)
	}
	t.Log(res)

	t.Logf("structMap: %+v \n", hash.Data)
}

func TestStructMap(t *testing.T) {

	hs := NewStructMap()
	hs.BuildIndex(getList())
	res1, err := hs.QueryData("李四", 21)
	if err != nil {
		t.Error("TestStructMap query fail:", err)
	}
	t.Log(res1)

	t.Logf("structMap: %+v \n", hs.Data)
}
