package helper

import "fmt"

func ErrorPanic(err error) {
	if err != nil {
		// defer fmt.Println("err\v", err)
		panic(err)
	}
}
func ErrorLog(des string, err error) {
	if err != nil {
		// defer fmt.Println("err\v", err)
		fmt.Println("\v \v", des, err)
	}
}
