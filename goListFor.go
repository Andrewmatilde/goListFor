package GoListFor

import (
	"math"
	"reflect"
)

func RangeInt(start,end,step int) []ListEle {
	size := int(math.Ceil(float64(end-start)/float64(step)))
	arrange := make([]ListEle, size)
	for i := 0;i < size ;i++  {
		arrange[i] = start + i * step
	}
	return arrange
}

type ListEle interface {}

func GetEle(e ListEle) interface{} {
	return reflect.ValueOf(e).Interface()
}

func ListFor(iList []ListEle,f func(Item ListEle) ListEle) []ListEle {
	var outList []ListEle
	for _, n := range iList {
		outList = append(outList,f(n))
	}
	return outList
}

func ToNoType(List interface{},f func(v interface{}) interface{}) []ListEle {
	if reflect.TypeOf(List).Kind().String()=="slice" {
		iList :=  make([]ListEle,reflect.ValueOf(List).Len())
		for i := 0;i < reflect.ValueOf(List).Len() ;i++  {
			iList[i] = f(reflect.ValueOf(List).Index(i).Interface())
		}
		return iList
	} else {
		return nil
	}
}

