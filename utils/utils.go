package utils

import (
	"fmt"
	"reflect"
)

const (
	InjectTagKey = "inject"
)

func GetFullUniqueName(instance interface{}) string {

	t := reflect.TypeOf(instance)
	if t == nil {
		panic("GetFullUniqueName parameter can't be nil")
	}
	//	not support slice map inject
	if kind := t.Kind(); kind == reflect.Slice || kind == reflect.Map {
		panic("type can't be slice or map")
	}

	if pkgPath := t.PkgPath(); pkgPath != "" {
		return fmt.Sprintf("%v/%v", pkgPath, t.String())
	}
	return t.String()
}

func FieldNeedToInject(f reflect.StructField) bool {
	_, ok := f.Tag.Lookup(InjectTagKey)
	return ok
}

func CanRegeiste(instance interface{}) bool {
	t := reflect.TypeOf(instance)
	return t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct
}
