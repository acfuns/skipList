package main

import "fmt"

func main() {
	sl := NewSkipList()
	sl.Insert(10)
	sl.Insert(20)
	sl.Insert(30)
	sl.Insert(40)
	sl.Insert(50)
	sl.Insert(60)
	sl.Insert(70)
	sl.Insert(80)
	sl.Insert(90)
	sl.Insert(100)
	sl.Insert(110)
	sl.Insert(120)
	res := sl.Search(60)
	fmt.Println(res) // true

	sl.Remove(60)

	res = sl.Search(60)
	fmt.Println(res) // false
}
