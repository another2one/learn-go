package funcs

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func assertEqual(actual any, want any, message string) {
	if !reflect.DeepEqual(actual, want) {
		log.Fatalf("%s want %+v but %+v", message, want, actual)
	}
}

func TestInArray1(t *testing.T) {
	assertEqual(InArray(162, []int{0, 0xa2, 4}), true, `test 16进制`)
	assertEqual(InArray(0.0, []float32{0, 2, 4}), true, `test '0.0'`)
}

func TestArrayChunk1(t *testing.T) {
	if _, err := ArrayChunk([]int{1, 2, 4, 8, 16}, -1); err == nil {
		log.Fatalf("%s", "-1 没有报错")
	}
	res, err := ArrayChunk([]int{1, 2, 4, 8, 16}, 3)
	assertEqual(res, [][]int{{1, 2, 4}, {8, 16}}, `test `)
	assertEqual(err, nil, `test `)
}

func TestArrayColumnMap(t *testing.T) {
	res, err := ArrayColumnMap([]map[string]string{{"name": "lizhi", "age": "12"}, {"name": "lizhi1", "age": "121"}}, "age", "name")
	assertEqual(res, map[string]string{"lizhi": "12", "lizhi1": "121"}, `test `)
	assertEqual(err, nil, `test `)
}

func TestArrayDiffAndIntersect(t *testing.T) {
	assertEqual(ArrayDiff([]string{"a", "b", "c", "a"}, []string{"b", "c", "d", "d"}), []string{"a"}, `test `)
	assertEqual(ArrayIntersect([]string{"a", "b", "c", "a"}, []string{"b", "c", "d", "c"}), []string{"b", "c"}, `test `)
}

func TestSlice(t *testing.T) {
	if _, err := InsertValue([]int{1, 2, 4}, 9, 9); err == nil {
		log.Fatalf("%s", "-1 没有报错")
	}
	res, err := InsertValue([]int{1, 2, 4}, 3, 2)
	assertEqual(res, []int{1, 2, 3, 4}, `test`)
	assertEqual(err, nil, `test `)

	assertEqual(DeleteValue([]int{1, 2, 3}, 2), []int{1, 3}, `test `)
	res, _ = DeleteIndex([]int{1, 2, 4}, 2)
	assertEqual(res, []int{1, 2}, `test `)
}

func TestSlice1(t *testing.T) {
	s1 := []int{1, 2, 3, 4}
	s2 := append(s1[:1], 6)
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
}
