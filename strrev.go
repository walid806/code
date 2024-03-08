package main

import "fmt"

func main() {
	s := "Hello World!"
	s = StrRev(s)
	fmt.Println(s)
}
func StrRev(s string) string {
	ss := ""
	for _, i := range s {
		ss = string(i) + ss
	}
	return string(ss)

}
