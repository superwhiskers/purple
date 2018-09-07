package purple

import (
	"fmt"
	"math/rand"
	"reflect"
	"sync"
	"time"
)

// RandomGenerator holds the minimum and maximum values allowed from number generators
//
// Fields:
// 	Min: the minimum value allowed from the generator
// 	Max: the maximum value allowed from the generator
type RandomGenerator struct {
	Min, Max int
	Source   *rand.Rand

	randMutex sync.Mutex
}

// NewRandomGenerator makes a new RandomGenerator
func NewRandomGenerator() *RandomGenerator {

	return &RandomGenerator{
		Min:    0,
		Max:    10,
		Source: rand.New(rand.NewSource(time.Now().Unix())),

		randMutex: sync.Mutex{},
	}

}

func (rg *RandomGenerator) random() int {

	rg.randMutex.Lock()
	defer rg.randMutex.Unlock()
	return rg.Source.Intn(rg.Max-rg.Min+1) + rg.Min

}

func (rg *RandomGenerator) seed(seed int64) {

	rg.randMutex.Lock()
	defer rg.randMutex.Unlock()
	rg.Source.Seed(seed)

}

// Reseed resets the seed to the current unix time
func (rg *RandomGenerator) Reseed() {

	rg.seed(time.Now().Unix())

}

// Seed sets the seed to the specified number
//
// Parameters:
// 	seed: the seed to set as the number generator's seed
func (rg *RandomGenerator) Seed(seed int64) {

	rg.seed(seed)

}

// Random returns a pseudorandomly generated number within the range specified without resetting the seed.
// if no parameters are provided, Random will use the same range as last time.
// if there are two parameters, it will use them as the maximum and the minumum.
// if there is one parameter, it will use it as the maximum
//
// Parameters:
// 	max: the maximum value allowed to be generated
// 	min: the minimum value allowed to be generated
//
// Returns:
// 	the generated number
func (rg *RandomGenerator) Random(args ...int) int {

	if len(args) == 2 {

		rg.Max = args[0]
		rg.Min = args[1]

	} else if len(args) == 1 {

		rg.Max = args[0]
		rg.Min = 0

	}

	return rg.random()

}

// NextRandom returns a pseudorandomly generated number within the range and also resets the seed.
// if no parameters are provided, NextRandom will use the same range as last time.
// if there are two parameters, it will use them as the maximum and the minumum.
// if there is one parameter, it will use it as the maximum
//
// Parameters:
// 	max: the maximum value allowed to be generated
// 	min: the minimum value allowed to be generated
//
// Returns:
// 	the generated number
func (rg *RandomGenerator) NextRandom(args ...int) int {

	if len(args) == 2 {

		rg.Max = args[0]
		rg.Min = args[1]

	} else if len(args) == 1 {

		rg.Max = args[0]
		rg.Min = 0

	}

	rg.Reseed()
	return rg.random()

}

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
