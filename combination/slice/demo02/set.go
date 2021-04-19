package demo02

import "sort"

// 去重 利用map[string]struct{} 可实现set
func SliceToSet1(s1 []int) []int {
	if len(s1) <= 1 {
		return s1
	}
	m := make(map[int]struct{})
	for _, v := range s1 {
		m[v] = struct{}{}
	}
	s2 := make([]int, 0)
	for index, _ := range m {
		s2 = append(s2, index)
	}
	sort.Ints(s2) // map无序，测试方便
	return s2
}

// 去重 利用map[string]struct{} 可实现set
func SliceToSet3(s1 []int) []int {
	if len(s1) <= 1 {
		return s1
	}
	m := make(map[int]struct{})
	s2 := make([]int, 0)
	for _, v := range s1 {
		if _, ok := m[v]; !ok {
			s2 = append(s2, v)
			m[v] = struct{}{}
		}

	}
	return s2
}

// 去重 直接对比
func SliceToSet2(s1 []int) []int {
	if len(s1) <= 1 {
		return s1
	}
	s2 := s1[:1]
	for _, v := range s1[1:] {
		tag := true
		for _, item := range s2 {
			if item == v {
				tag = false
				break
			}
		}
		if tag {
			s2 = append(s2, v)
		}
	}
	return s2
}

// 去重 直接对比
func StringSliceToSet(s1 []string) []string {
	if len(s1) <= 1 {
		return s1
	}
	s2 := s1[:1]
	for _, v := range s1[1:] {
		tag := true
		for _, item := range s2 {
			if item == v {
				tag = false
				break
			}
		}
		if tag {
			s2 = append(s2, v)
		}
	}
	return s2
}
