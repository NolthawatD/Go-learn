package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Num interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

type UserID int // this is an alias for int

func Add[T ~int](a T, b T) T {
	return a + b
}

/*
*
*  Normal
*
 */
func Normal() {
	result := Add(1, 2)
	println(result)

	a := UserID(11)
	b := UserID(22)

	result2 := Add(a, b)
	fmt.Println("use ~int aliased int type")
	println(result2)
}

/*
*
*  Slice mapping
*
 */

func MapValues[T constraints.Ordered](values []T, mapFunc func(T) T) []T {
	var newValues []T
	for _, v := range values {
		newValue := mapFunc(v)
		newValues = append(newValues, newValue)
	}
	return newValues
}

func Slice() {
	result1 := MapValues([]int{1, 2, 3}, func(n int) int {
		return n * 2
	})
	fmt.Printf("result 1: %v\n", result1)

	result2 := MapValues([]float32{1.1, 2.2, 3.3}, func(n float32) float32 {
		return n * 2
	})
	fmt.Printf("result 2: %v\n", result2)
}

/*
*
*  Struct
*
 */
type CustomData interface {
	constraints.Ordered | []byte | []rune
}

type User[T CustomData] struct {
	ID   int
	Name string
	Data T
}

func Struct() {
	u := User[string]{
		ID:   1,
		Name: "nolthawat",
		Data: "Can modify type of T",
	}

	fmt.Printf("user: %v\n", u)
}

/*
*
*  Comparable
*
 */
type CustomMapped[T comparable, V int | string] map[T]V // a(string) === b(string)

func Comparable() {
	m := make(CustomMapped[int, string])
	m[3] = "9"
	m[1] = "5"
	fmt.Printf("m: %v\n", m)
}

func main() {
	fmt.Println("Generics in Go")

	fmt.Println("Normal:")
	Normal()

	fmt.Println("Slice:")
	Slice()

	fmt.Println("Struct:")
	Struct()

	fmt.Println("Comparable:")
	Comparable()

}
