package main

import (
	"fmt"
	"sort"
	"math/rand"
	"time"
)

// func main() {

// 	rand.Seed(time.Now().Unix())
// 	var s1 []hero
// 	for i := 0; i < 10; i++ {
// 		s1 = append(s1, hero{
// 			"英雄" + fmt.Sprintf("%d", rand.Intn(100)) + "号",
// 			rand.Intn(100),
// 		})
// 	}
// 	var heros St = s1
// 	sort.Sort(heros)
// 	fmt.Println(s1)
// }

// type hero struct {
// 	Name string
// 	Age int
// }

// type St []hero

// func (s St) Len() int {
// 	return len(s)
// }

// func (s St) Less(i, j int) bool {
// 	return s[i].Age < s[j].Age
// }

// func (s St) Swap(i, j int) {
// 	s[i], s[j] = s[j], s[i]
// }

type Student struct {
	Name string
	Age int
	Score float64
}

type stuSlice []Student

func (stus stuSlice) Len() int {
	return len(stus)
}

func (stus stuSlice) Less(i, j int) bool {
	return stus[i].Score < stus[j].Score
}

func (stus stuSlice) Swap(i, j int) {
	stus[i], stus[j] = stus[j], stus[i]
}

func main(){
	
	rand.Seed(time.Now().Unix())
	stus := make([]Student, 10, 10)
	for index, _ := range stus {
		stus[index] = Student{
			"李" + fmt.Sprintf("%d", rand.Intn(100)+1) + "盼",
			rand.Intn(100),
			float64(rand.Intn(100))/2,
		}
	}
	var stuSlice stuSlice = stus
	fmt.Println(stuSlice)
	sort.Sort(stuSlice)
	fmt.Println(stuSlice)
}