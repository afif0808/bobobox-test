package structs

import (
	"errors"
	"reflect"
)

// Utility package for stuffs related to struct
// inspired by : https://dave.cheney.net/practical-go/presentations/qcon-china.html#_avoid_package_names_like_base_common_or_util

// Merge merges any structs which rule determined by 'merge' tag
func Merge(dst interface{}, src interface{}) error {
	dstVal := reflect.ValueOf(dst)
	dstType := reflect.TypeOf(dst)
	if dstVal.Kind() != reflect.Ptr || dstVal.Elem().Kind() != reflect.Struct {
		return errors.New("destination should be a struct pointer")
	}
	srcVal := reflect.ValueOf(src)
	srcTyp := reflect.TypeOf(src)
	if srcVal.Kind() == reflect.Ptr {
		srcVal = srcVal.Elem()
	}

	if srcVal.Kind() != reflect.Struct {
		return errors.New("source should be a struct type")
	}
	matches := map[string]reflect.Value{}

	for i := 0; i < srcVal.NumField(); i++ {
		tag, ok := srcTyp.Field(i).Tag.Lookup("merge")
		if !ok {
			continue
		}
		matches[tag] = srcVal.Field(i)
	}

	for i := 0; i < dstVal.Elem().NumField(); i++ {
		tag, exist := dstType.Elem().Field(i).Tag.Lookup("merge")
		if !exist {
			continue
		}
		match, exist := matches[tag]
		if !exist {
			continue
		}
		dstVal.Elem().Field(i).Set(match)
	}

	return nil
}
