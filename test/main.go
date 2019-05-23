package main

import (
	"fmt"
	GoListFor "prange"
)

func main()  {
	//EleList := ListFor.RangeInt(0,10,1)
	//EleList = ListFor.ListFor(EleList, func(Item ListFor.ListEle) ListFor.ListEle {
	//	return ListFor.GetEle(Item).(int) + 1
	//})
	//fmt.Println(EleList)
	i := []int32{2,3,4}
	EleList := GoListFor.ToNoType(i, func(v interface{}) interface{} {
		return v.(int32)
	})
	EleList = GoListFor.ListFor(EleList, func(Item GoListFor.ListEle) GoListFor.ListEle {
		return complex(float32(GoListFor.GetEle(Item).(int32)), 2)
	})
	fmt.Println(EleList)
}
