package main

import (
	"fmt"
	"reflect"
)

func main() {
	var str = []string {"abuja", "ogun", "ibadan", "lagos", "osun", "jos"}

	fmt.Println(itemExist(str, "abuja"))
	fmt.Println(itemExist(str, "abia"))
}

func itemExist(slice interface{}, item interface{}) bool {
	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice {
		panic("Invalid data type")
	}

	for i := 0; i < s.Len() ; i++ {
		if s.Index(i).Interface() ==  item {
			return true
		}

	}

	return false
}