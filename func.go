// package main

// import "fmt"

// func multiply(n int) int {
// 	return n * n
// }

// func main() {
// 	val := 3
// 	fmt.Println("pass value ",multiply(val))

// }

//PASS STRUCT


package main 
import "fmt"


type Company struct {
	Name string
	Description string
}

func UpdateCompanyDetails(c Company) Company{
	c.Name = "gogoglinux"
	c.Description = "linux for golang"
	return c
}

func main () {
	res := Company{
		Name: "okay ide",
		Description: "linux foundation",

	}

	fmt.Println("pass valude",UpdateCompanyDetails(res))
}