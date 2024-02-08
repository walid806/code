package main
import "fmt"
func prim(a rune) bool {
	if (a>= 'A' && a <= 'Z' ) || (a>= 'a' && a<='z' ) || (a>= '0' && a<= '9'){
		return true
	}
	return false 
}
func Capitalize(s string) string {
	dd := []rune(s)
	letra := true 
	for i := 0 ; i < len (s) ;  i ++ {
		if prim(dd[i]) == true && letra {
			if dd[i] >= 'a' && dd[i]<='z' {
				dd[i] = 'A' - 'a' + dd[i]
			}
			letra = false
		}else if dd[i] >= 'A' && dd[i] <= 'Z'{
			dd[i] = 'a' - 'A' + dd[i]
		}else if prim(dd[i]) == false {
			letra = true
		}
	}
	return string(dd)
}
func main() {
	fmt.Println(Capitalize("Hello! How are you? How+are+things+4you?"))
}