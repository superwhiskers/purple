package purple

import (
	"fmt"
	"reflect"
)

// ContainsItem checks if the given slice or array has the given item
//
// Parameters:
// 	islice: a slice or array of any type to search for iitem in
// 	iitem: the item to search islice in
//
// Returns:
// 	the index(es) of the element and a boolean if it was actually in there
func ContainsItem(islice, iitem interface{}) ([]int, bool) {

	vslice := reflect.ValueOf(islice)

	indexes := []int{}
	exists := false

	for i := 0; i < vslice.Len(); i++ {

		vindex := vslice.Index(i)

		if reflect.DeepEqual(vindex.Interface(), iitem) {

			indexes = append(indexes, i)
			exists = true

		}

	}

	return indexes, exists

}

// ContainsItemOnce checks if the given slice or array has the given item at least once
//
// Parameters:
// 	islice: a slice or array of any type to search for iitem in
// 	iitem: the item to search islice in
//
// Returns:
// 	the index of the element and a boolean if it was actually in there
func ContainsItemOnce(islice, iitem interface{}) (int, bool) {

	vslice := reflect.ValueOf(islice)

	for i := 0; i < vslice.Len(); i++ {

		vindex := vslice.Index(i)

		if reflect.DeepEqual(vindex.Interface(), iitem) {

			return i, true

		}

	}

	return -1, false

}

// RemoveUnordered removes an index from a slice without regards to order of the slice
//
// Parameters:
// 	islice: a slice of any type to remove the index from
// 	iindex: a valid index of islice
//
// Returns:
// 	an interface that is the same slice without the chosen index
func RemoveUnordered(islice, iindex interface{}) interface{} {

	var index int

	vindex := reflect.ValueOf(iindex)
	vslice := reflect.ValueOf(islice)

	indexType := vindex.Type()
	switch indexType.Kind() {

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		index = int(vindex.Int())

	default:
		panic(fmt.Errorf("iindex should be an int but is instead of type %s", indexType.Kind().String()))

	}

	lastInd := vslice.Index(vslice.Len() - 1)
	delInd := vslice.Index(index)
	savedLastInd := reflect.ValueOf(reflect.Indirect(lastInd).Interface())
	savedDelInd := reflect.ValueOf(reflect.Indirect(delInd).Interface())

	lastInd.Set(savedDelInd)
	delInd.Set(savedLastInd)

	return vslice.Slice(0, vslice.Len()-1).Interface()

}

// RemoveOrdered removes an index from a slice with regards to the order of the slice
//
// Parameters:
// 	islice: a slice of any type to remove the index from
// 	iindex: a valid index of islice
//
// Returns:
// 	an interface that is the same slice without the chosen index
func RemoveOrdered(islice, iindex interface{}) interface{} {

	var index int

	vindex := reflect.ValueOf(iindex)
	vslice := reflect.ValueOf(islice)
	indexType := vindex.Type()

	switch indexType.Kind() {

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		index = int(vindex.Int())

	default:
		panic(fmt.Errorf("iindex should be an int but is instead of type %s", indexType.Kind().String()))

	}

	latterHalf := vslice.Slice(index+1, vslice.Len())
	vslice = vslice.Slice(0, index)

	return reflect.AppendSlice(vslice, latterHalf).Interface()

}
