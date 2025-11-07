package main

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/samber/lo/mutable"
	"strconv"
)

func main() {
	even1 := lo.Filter([]int{1, 2, 3, 4}, func(x int, index int) bool {
		return x%2 == 0
	})
	fmt.Println("Filter", even1)
	even2 := lo.Map([]int{1, 2, 3, 4}, func(x int, index int) int {
		return x * x
	})
	fmt.Println("Map", even2)
	even3 := lo.UniqMap([]int{1, 2, 3, 2}, func(x int, index int) int {
		return x % 2
	})
	fmt.Println("UniqMap", even3)
	even4 := lo.FilterMap([]int{1, 2, 3, 4}, func(x int, index int) (int, bool) {
		return x * 2, x*2 < 7
	})
	fmt.Println("FilterMap", even4)
	even5 := lo.FlatMap([]int64{1, 2, 3, 4}, func(x int64, index int) []string {
		return []string{strconv.FormatInt(x, 10), strconv.FormatInt(x, 2)}
	})
	fmt.Println("FlatMap", even5)
	even6 := lo.Reduce([]int{1, 2, 3, 4}, func(agg string, x int, index int) string {
		return agg + "-" + strconv.Itoa(x)
	}, "pre")
	fmt.Println("Reduce", even6)
	even7 := lo.ReduceRight([]int{1, 2, 3, 4}, func(agg string, x int, index int) string {
		return agg + "-" + strconv.Itoa(x)
	}, "pre")
	fmt.Println("ReduceRight", even7)
	even8 := lo.Times(3, func(index int) int {
		return index + 1
	})
	fmt.Println("Times", even8)
	even9 := lo.Uniq([]int{1, 2, 3, 2})
	fmt.Println("Uniq", even9)
	even10 := lo.UniqBy([]int{1, 2, 3, 4}, func(item int) int {
		return item % 2
	})
	fmt.Println("UniqBy", even10)
	even11 := lo.GroupBy([]int{1, 2, 3, 4}, func(item int) string {
		if item%2 == 0 {
			return "偶数"
		} else {
			return "奇数"
		}
	})
	fmt.Println("GroupBy", even11)
	even12 := lo.Chunk([]int{1, 2, 3, 4, 5}, 2)
	fmt.Println("Chunk", even12)
	even13 := lo.PartitionBy([]int{1, 2, 3, 4}, func(item int) int {
		return item%2 + 1
	})
	fmt.Println("PartitionBy", even13)
	even14 := lo.Flatten([][]int{{1, 2}, {3, 4}})
	fmt.Println("Flatten", even14)
	even15 := lo.Interleave([]int{1, 2}, []int{3, 4})
	fmt.Println("Interleave", even15)
	mutable.Shuffle(even14)
	fmt.Println("Shuffle", even14)
	mutable.Reverse(even15)
	fmt.Println("Reverse", even15)
	even17 := lo.Fill([]CloneSlice{{1}, {2}}, CloneSlice{0})
	fmt.Println("Fill", even17)
	even18 := lo.Repeat(3, CloneSlice{1})
	fmt.Println("Repeat", even18)
	even19 := lo.RepeatBy(3, func(index int) int {
		return index * index
	})
	fmt.Println("RepeatBy", even19)
	even20 := lo.KeyBy([]map[string]string{{"name": "lizhi", "id": "1"}}, func(item map[string]string) string {
		return item["id"]
	})
	fmt.Println("KeyBy", even20)
	even21 := lo.Associate([]map[string]string{{"name": "lizhi", "id": "1"}}, func(item map[string]string) (string, string) {
		return item["id"], item["name"]
	})
	fmt.Println("Associate", even21)
	even22 := lo.FilterSliceToMap([]map[string]string{{"name": "lizhi", "id": "1"}, {"name": "lizhi", "id": "2"}}, func(item map[string]string) (string, string, bool) {
		return item["id"], item["name"], item["id"] == "1"
	})
	fmt.Println("FilterSliceToMap", even22)
	even23 := lo.Keyify([]int{1, 2, 3, 1})
	fmt.Println("Keyify", even23)
	even24 := lo.Drop([]int{1, 2, 3, 4}, 2)
	fmt.Println("Drop", even24)
	even25 := lo.DropRight([]int{1, 2, 3, 4}, 2)
	fmt.Println("DropRight", even25)
	even26 := lo.DropWhile([]int{1, 2, 3, 4}, func(item int) bool {
		return item != 3
	})
	fmt.Println("DropWhile", even26)
	even27 := lo.DropRightWhile([]int{1, 2, 3, 4}, func(item int) bool {
		return item != 3
	})
	fmt.Println("DropRightWhile", even27)
	even28 := lo.DropByIndex([]int{1, 2, 3, 4}, 1, 2)
	fmt.Println("DropByIndex", even28)
	even29 := lo.Count([]int{1, 2, 3, 1}, 1)
	fmt.Println("Count", even29)
	even30 := lo.CountBy([]int{1, 2, 3, 4}, func(item int) bool {
		return item != 3
	})
	fmt.Println("DropRightWhile", even30)
	even31 := lo.CountValues([]int{1, 2, 3, 1})
	fmt.Println("Count", even31)
	even32 := lo.CountValuesBy([]map[string]string{{"name": "lizhi", "id": "1"}, {"name": "lizhi", "id": "2"}}, func(item map[string]string) string {
		return item["name"]
	})
	fmt.Println("Count", even32)
	even33 := lo.Subset([]int{1, 2, 3, 4, 5}, 1, 3)
	fmt.Println("Subset", even33)
	even34 := lo.Slice([]int{1, 2, 3, 4, 5}, 1, 3)
	fmt.Println("Slice", even34)
	even35 := lo.Replace([]int{1, 2, 3, 1, 2, 3}, 1, 6, 1)
	fmt.Println("Replace", even35)
	even36 := lo.ReplaceAll([]int{1, 2, 3, 1, 2, 3}, 1, 6)
	fmt.Println("ReplaceAll", even36)
	even37 := lo.Compact([]int{1, 2, 3, 0})
	fmt.Println("Compact", even37)
	even38 := lo.IsSorted([]int{1, 2, 3, 4})
	fmt.Println("IsSorted", even38)
	even39 := lo.IsSortedByKey([]map[string]string{{"name": "lizhi", "id": "1"}, {"name": "lizhi", "id": "2"}}, func(item map[string]string) string {
		return item["id"]
	})
	fmt.Println("IsSortedBy", even39)
	even40 := lo.Splice([]int{1, 2, 3, 4}, -1, 1, 1)
	fmt.Println("Splice", even40)
}

type CloneSlice []int

func (c1 CloneSlice) Clone() CloneSlice {
	clone := make([]int, len(c1))
	copy(clone, c1)
	return clone
}
