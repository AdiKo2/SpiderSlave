package Utils

import "fmt"

//type Deck[C any] struct {
//	cards []C
//}

type Slice2D [][]any

type SliceString2D [][]string

type SliceString []string

type Deck5 [][]any

//func (d *Deck[C]) ToStringSlice() SliceString {
//	var slice []string
//	for _, value := range d.cards {
//		slice = append(slice, fmt.Sprintf("%v", value))
//	}
//
//	return slice
//}

func (d Slice2D) ToStringSlice2() SliceString2D {
	var bigSlice SliceString2D
	for _, user := range d {
		var slice []string
		for _, parameter := range user {
			slice = append(slice, fmt.Sprintf("%v", parameter))
		}
		bigSlice = append(bigSlice, slice)
	}
	return bigSlice
}

func (d SliceString2D) BatchUsers(batchSize int, function func([][]string)) {
	length := len(d)
	for i := 0; i < length; i++ {
		//		for i := 0; i < length; i + batchSize {
		j := i + batchSize
		if j > length {
			j = length
		}
		function(d[i:j])
	}
}

func (d Slice2D) BatchUsers2(batchSize int, function func(records [][]string) error) {
	amount := len(d)
	for startIndex := 0; startIndex < amount; startIndex++ {

		//		for startIndex := 0; startIndex < amount; startIndex + batchSize {
		endIndex := startIndex + batchSize
		if endIndex > amount {
			endIndex = amount
		}
		if err := function(d[startIndex:endIndex].ToStringSlice2()); err != nil {
			fmt.Println("current error is:", err)
		}
	}
}
