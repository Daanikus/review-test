package main

import (
	"fmt"
)

func main() {
	msg := "Exported"
	Export(msg)
	sl := []int{2, 3, 4, 5, 6, 7}
	_ = GetSubSlice(1, sl)
	_ = Do()
}

func Export(msg string) error {
	fmt.Println("Exported")
	fmt.Println(msg)
	return nil
}

func GetSubSlice(start int, sl []int) []int {
	return sl[start:len(sl)]
}

func Do() error {
	err := Export("Hello")
	if err != nil {
		return err
	}
	return nil
}
