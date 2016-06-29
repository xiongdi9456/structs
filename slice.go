package structs

import (
	"reflect"
)

func SliceMap(s interface{}) []map[string]interface{} {
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Slice {
		panic("not slice")
	}
	sliceMap := []map[string]interface{}{}

	for i := 0; i < v.Len(); i++ {
		subv := v.Index(i)
		if subv.Kind() == reflect.Ptr {
			subv = subv.Elem()
		}
		s := Struct{
			raw:     subv,
			value:   subv,
			TagName: DefaultTagName,
		}
		sliceMap = append(sliceMap, s.Map())
	}
	return sliceMap
}
