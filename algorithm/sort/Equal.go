package sort

import "reflect"

func Equals(x, y interface{}) bool {
	if x == nil || y == nil {
		return x == y
	}
	return reflect.DeepEqual(x, y)
}
