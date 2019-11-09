package handler

import (
	"github.com/valyala/fasthttp"
	"reflect"
	"strconv"
	"unsafe"
)

func FillFieldValue(data map[string]interface{}, in reflect.Type) (rValue reflect.Value, ok bool) {
	if in == nil {
		ok = false
		return
	}
	if in.Kind() == reflect.Ptr {
		in = in.Elem()
		rValue = reflect.New(in)
	}
	numField := in.NumField()
	for i := 0; i < numField; i++ {
		field := in.Field(i)
		fieldPtr := uintptr(unsafe.Pointer(rValue.Pointer())) + field.Offset
		fN := field.Tag.Get("json")
		if fV, ok := data[fN]; ok {
			switch fV.(type) {
			case string:
				*((*string)(unsafe.Pointer(fieldPtr))) = fV.(string)
			case int:
				*((*int)(unsafe.Pointer(fieldPtr))) = fV.(int)
			}
		}
	}
	ok = true
	return
}

func FillFieldValueByQueryArgs(args *fasthttp.Args, in reflect.Type) (rValue reflect.Value, ok bool) {
	if in == nil {
		ok = false
		return
	}
	if in.Kind() == reflect.Ptr {
		in = in.Elem()
		rValue = reflect.New(in)
	}
	numField := in.NumField()
	for i := 0; i < numField; i++ {
		field := in.Field(i)
		fieldPtr := uintptr(unsafe.Pointer(rValue.Pointer())) + field.Offset
		fN := field.Tag.Get("json")
		if fV := args.Peek(fN); fV != nil {
			switch field.Type.Kind() {
			case reflect.String:
				*((*string)(unsafe.Pointer(fieldPtr))) = string(fV)
			case reflect.Int64:
				atoi, err := strconv.ParseInt(string(fV), 10, 64)
				if err != nil {
					ok = false
					return
				}
				*((*int64)(unsafe.Pointer(fieldPtr))) = atoi
			case reflect.Int:
				atoi, err := strconv.Atoi(string(fV))
				if err != nil {
					ok = false
					return
				}
				*((*int)(unsafe.Pointer(fieldPtr))) = atoi
			default:
				ok = false
			}
		}
	}
	ok = true
	return
}
