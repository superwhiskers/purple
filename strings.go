package purple

import (
	"fmt"
)

// Pad pads a byte slice or string with the specified character to the specified length
//
// Parameters:
// 	ipad: the byte slice or string to pad
// 	ichar: the byte or character to pad with
// 	length: the length to pad it to
//
// Returns:
// 	the padded string
func Pad(ipad, ichar interface{}, length int) interface{} {

	var (
		char    uint8
		pad     []byte
		padding []byte

		isString = false
	)

	switch ichar.(type) {

	case uint8:
		char = ichar.(uint8)

	case string:
		if len(ichar.(string)) != 1 {

			panic(fmt.Errorf("ichar should be only one character long but is instead %d characters long", len(ichar.(string))))

		}
		char = []byte(ichar.(string))[0]

	default:
		panic(fmt.Errorf("ichar should be either a byte or a character but is instead a different type"))

	}

	switch ipad.(type) {

	case string:
		pad = []byte(ipad.(string))
		isString = true

	case []byte:
		pad = ipad.([]byte)

	default:
		panic(fmt.Errorf("ipad should be either a byte slice or a string but is instead of a different type"))

	}

	lengthToPad := length - len(pad)
	if lengthToPad < 1 {

		return pad

	}

	for x := 0; x < lengthToPad; x++ {

		padding = append(padding, char)

	}

	pad = append(pad, padding...)

	if isString == true {

		return string(pad[:])

	}

	return pad

}

// LeftPad pads a byte slice or string with the specified character to the specified length but on the left side instead of the right
//
// Parameters:
// 	ipad: the byte slice or string to pad
// 	ichar: the byte or character to pad with
// 	length: the length to pad it to
//
// Returns:
// 	the padded string
func LeftPad(ipad, ichar interface{}, length int) interface{} {

	var (
		char    uint8
		pad     []byte
		padding []byte

		isString = false
	)

	switch ichar.(type) {

	case uint8:
		char = ichar.(uint8)

	case string:
		if len(ichar.(string)) != 1 {

			panic(fmt.Errorf("ichar should be only one character long but is instead %d characters long", len(ichar.(string))))

		}
		char = []byte(ichar.(string))[0]

	default:
		panic(fmt.Errorf("ichar should be either a byte or a character but is instead a different type"))

	}

	switch ipad.(type) {

	case string:
		pad = []byte(ipad.(string))
		isString = true

	case []byte:
		pad = ipad.([]byte)

	default:
		panic(fmt.Errorf("ipad should be either a byte slice or a string but is instead of a different type"))

	}

	lengthToPad := length - len(pad)
	if lengthToPad < 1 {

		return pad

	}

	for x := 0; x < lengthToPad; x++ {

		padding = append(padding, char)

	}

	pad = append(padding, pad...)

	if isString == true {

		return string(pad[:])

	}

	return pad

}
