package main
import "fmt"


func main () {

	var checkInterface interface {} = "linux tut"
	checkType := checkInterface.(string)
	fmt.Println(checkType)

}