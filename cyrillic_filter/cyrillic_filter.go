package cyrillic_filter

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
)

func Filter(i interface{}) error{
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("non-pointer %v",v.Type())
	}
	// get the value that the pointer v points to.
	v = v.Elem()
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("can't fill non-struct value %v", v.Type())
	}
	for i:=0;i<v.NumField();i++ {
		f:=v.Field(i)
		if f.Kind() == reflect.Ptr && f.Elem().Kind()== reflect.String{
			result := Filter(*f.Interface().(*string))
			f.Set(reflect.ValueOf(&result))
		}else if f.Kind() == reflect.String{
			result := Filter(f.Interface().(string))
			f.Set(reflect.ValueOf(result))
		}
	}
	return nil
}

func deleteCyrillic(s string) string{
	result := ""
	for _, v := range s{
		reg, err := regexp.Compile("[А-яЁё]")
		if err != nil {
			log.Fatal(err)
		}
		newStr := reg.ReplaceAllString(string(v), "")
			result +=newStr
	}
	return result
}
