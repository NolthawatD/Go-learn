package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

// https://dev.to/pallat/eriiynruueruueng-fuzzing-ain-go-2jh5
/*
	- สามารถรันเทสด้วยคำสั่ง go test -run=FuzzReverse ในกรณีที่มีเทสอื่นที่เราไม่อยากเทสอยู่ด้วย
	- go test -fuzz=Fuzz ตรวจสอบการสร้าง input แบบสุ่มเข้าไปดูบ้างว่าจะมีอะไรผิดพลาดเกิดขึ้น
	  ปัญหานี้จะถูกเขียนลงไปในไฟล์ seed corpus ซึ่งมันจะถูกรันในการสั่ง go test ครั้งถัดไป
*/

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev, revErr := Reverse(input)
	doubleRev, doubleRevErr := Reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q, err: %v\n", rev, revErr)
	fmt.Printf("reversed again: %q, err: %v\n", doubleRev, doubleRevErr)
}

func Reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r), nil
}
