package utils

import (
	"fmt"
	"reflect"
)

func ArrayToStrings(input interface{}) []string {
	s := reflect.ValueOf(input)
	if s.Kind() != reflect.Slice {
		return nil
	}

	ret := make([]string, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = fmt.Sprint(s.Index(i).Interface())
	}
	return ret
}
