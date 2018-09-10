package purple

import (
	"fmt"
	"reflect"
)

// ForEach executes a function on each index of an iterable
//
// Parameters:
//  islice: the slice to run the function on
//  ifunction: the function to run on each index
// 	 Parameters:
// 	 	index interface{}: the index the function is being run on
// 	 	slice interface{}: the entire slice the function is being run over
// 	 	indexNum int: the index number we are at
//
// 	 Returns:
// 	 	it can only return up to one datatype, otherwise it's considered an invalid function
//
// Returns:
// 	a list of all of the data that was returned by each run of the function
func ForEach(islice, ifunction interface{}) []interface{} {

	vslice := reflect.ValueOf(islice)
	vfunction := reflect.ValueOf(ifunction)

	if vfunction.Kind() != reflect.Func {

		panic(fmt.Errorf("ifunction should be a function but is instead of type %s", vfunction.Kind().String()))

	}

	if vfunction.Type().NumIn() != 3 {

		panic(fmt.Errorf("ifunction should only have three input parameters but it instead has %d", vfunction.Type().NumIn()))

	}

	if vfunction.Type().NumOut() > 1 {

		panic(fmt.Errorf("ifunction can only have up to one return value but it instead has %d", vfunction.Type().NumOut()))

	}

	if vfunction.Type().In(0).Kind() != reflect.Interface {

		panic(fmt.Errorf("parameter 1 of ifunction should be an interface, not %s", vfunction.Type().In(0).Kind().String()))

	}

	if vfunction.Type().In(1).Kind() != reflect.Interface {

		panic(fmt.Errorf("parameter 2 of ifunction should be an interface, not %s", vfunction.Type().In(1).Kind().String()))

	}

	if vfunction.Type().In(2).Kind() != reflect.Int {

		panic(fmt.Errorf("parameter 3 of ifunction should be an interface, not %s", vfunction.Type().In(2).Kind().String()))

	}

	returnValues := []interface{}{}

	for i := 0; i < vslice.Len(); i++ {

		vindex := vslice.Index(i)

		returnData := vfunction.Call([]reflect.Value{vindex, vslice, reflect.ValueOf(i)})

		if len(returnData) != 0 {

			returnValues = append(returnValues, returnData[0].Interface())

		}

	}

	return returnValues

}

// IndexesOf gives you the indexes of every occurence of the given item in the given slice
//
// Parameters:
// 	islice: a slice or array of any type to search for iitem in
// 	iitem: the item to search islice in
//
// Returns:
// 	the index(es) of the element and a boolean if it was actually in there
func IndexesOf(islice, iitem interface{}) ([]int, bool) {

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

// IndexOf gives you the index of the first occurence of the given item in the given slice
//
// Parameters:
// 	islice: a slice or array of any type to search for iitem in
// 	iitem: the item to search islice in
//
// Returns:
// 	the index of the element and a boolean if it was actually in there
func IndexOf(islice, iitem interface{}) (int, bool) {

	vslice := reflect.ValueOf(islice)

	for i := 0; i < vslice.Len(); i++ {

		vindex := vslice.Index(i)

		if reflect.DeepEqual(vindex.Interface(), iitem) {

			return i, true

		}

	}

	return -1, false

}

// LastIndexOf gives you the index of the last occurence of the given item in the given slice
//
// Parameters:
// 	islice: a slice or array of any type to search for iitem in
// 	iitem: the item to search islice in
//
// Returns:
// 	the index of the element and a boolean if it was actually in there
func LastIndexOf(islice, iitem interface{}) (int, bool) {

	vslice := reflect.ValueOf(islice)

	for i := vslice.Len() - 1; i > -1; i-- {

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
