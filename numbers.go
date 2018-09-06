package purple

import (
	"fmt"
	"reflect"
)

// Sum calculates the sum of the input slice and returns the float64 sum of it
//
// Parameters:
// 	islice: a slice with any type of number in it
//
// Returns:
// 	a sum of all of the numbers in the slice with the same type as what you put in
func Sum(islice interface{}) interface{} {

	var (
		vindex reflect.Value

		vslice = reflect.ValueOf(islice)
	)

	if vslice.Kind() != reflect.Slice {

		panic(fmt.Errorf("a %s is not a slice", vslice.Kind().String()))

	}

	if vslice.Len() == 0 {

		return reflect.Zero(vslice.Type().Elem()).Interface()

	}

	rvalue := vslice.Index(0)

	for i := 1; i < vslice.Len(); i++ {

		vindex = vslice.Index(i)

		switch vindex.Kind() {

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rvalue.SetInt(rvalue.Int() + vindex.Int())

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			rvalue.SetUint(rvalue.Uint() + vindex.Uint())

		case reflect.Float32, reflect.Float64:
			rvalue.SetFloat(rvalue.Float() + vindex.Float())

		default:
			panic(fmt.Errorf("a %s is not a summable type", vindex.Kind().String()))

		}

	}

	return rvalue.Interface()

}
