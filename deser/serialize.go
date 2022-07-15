package deser

import (
	"fmt"
	"reflect"
)

func (self deserialized) serialize() string {
	for key, element := range self {
		fmt.Println(key, reflect.TypeOf(element))
	}
	return ""
}
