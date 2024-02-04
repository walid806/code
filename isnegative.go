package main
import "fmt"

func main() {
	IsNegative(10)
	IsNegative(0)
	IsNegative(-1)
}
func IsNegative(nb int) {
if nb > 0 {
fmt.Println("1")
}
if nb == 0  {
	fmt.Println("0")
}
if nb < 0 {
	fmt.Println("-1")
}
}