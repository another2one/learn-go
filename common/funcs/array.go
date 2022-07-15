package funcs

import (
	"errors"
	"math"
)

// https://segmentfault.com/a/1190000041634906
type Integer interface {
	Signed | Unsigned
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Float interface {
	~float32 | ~float64
}

type IntUintFloatString interface {
	Integer | Float | ~string
}

// InArray 判断 item 是否在 s1 里面
//  true 表示 item 在 s1 里面;  false 表示不在
func InArray[T comparable](item T, s1 []T) bool {
	for _, v := range s1 {
		if v == item {
			return true
		}
	}
	return false
}

// ArrayChunk 把一个数组分割为新的数组块
//  true 表示 item 在 s1 里面;  false 表示不在
func ArrayChunk[T any](s1 []T, length int) ([][]T, error) {
	if length <= 0 {
		return [][]T{s1}, errors.New("长度错误")
	}
	arrLen := len(s1)
	if length >= arrLen {
		return [][]T{s1}, nil
	}
	cap := int(math.Ceil(float64(arrLen) / float64(length)))
	s2 := make([][]T, cap)
	for i := 0; i < cap; i++ {
		right := (i + 1) * length
		if right >= arrLen {
			right = arrLen
		}
		s2[i] = s1[i*length : right]
	}
	return s2, nil
}

// 多数情况下是 []struct
func ArrayColumnMap[T IntUintFloatString](s1 []map[string]T, column_key, index_key string) (ret map[T]T, err error) {
	if column_key == "" {
		return ret, errors.New("column_key 不能为空")
	}
	arrLen := len(s1)
	if arrLen == 0 {
		return ret, errors.New("数组不能为空")
	}
	ret = make(map[T]T, arrLen)
	for _, val := range s1 {
		v, ok := val[column_key]
		k, ok2 := val[index_key]
		if ok && ok2 {
			ret[k] = v
		}
	}
	return
}

// 数组差集
func ArrayDiff[T comparable](s1 []T, s2 []T) (s []T) {
	m1 := make(map[T]bool, len(s1))
	for _, v := range s1 {
		m1[v] = false
	}
	for _, v := range s2 {
		if _, ok := m1[v]; ok {
			m1[v] = true
		}
	}
	for k, v := range m1 {
		if !v {
			s = append(s, k)
		}
	}
	return
}

// 数组交集
func ArrayIntersect[T comparable](s1 []T, s2 []T) (s []T) {
	m1 := make(map[T]bool, len(s1))
	for _, v := range s1 {
		m1[v] = false
	}
	for _, v := range s2 {
		if _, ok := m1[v]; ok {
			m1[v] = true
		}
	}
	for k, v := range m1 {
		if v {
			s = append(s, k)
		}
	}
	return
}

func InsertValue[T any](s1 []T, v T, index int) ([]T, error) {
	if index > len(s1)-1 || index < 0 {
		return s1, errors.New("index error, out of slice ... ")
	}
	s2 := make([]T, len(s1)-index)
	copy(s2, s1[index:])
	return append(append(s1[:index], v), s2...), nil
}

func DeleteValue[T comparable](s1 []T, item T) []T {
	s2 := make([]T, 0)
	for _, v := range s1 {
		if v != item {
			s2 = append(s2, v)
		}
	}
	return s2
}

func DeleteIndex[T any](s1 []T, index int) ([]T, error) {
	if index > len(s1)-1 || index < 0 {
		return s1, errors.New("index error, out of slice ... ")
	}
	return append(s1[:index], s1[index+1:]...), nil
}
