// package main

// import "fmt"

// func main() {
// 	MyChan := make(chan string)
// 	MyChan <- "okay this is a channel"

// 	fmt.Println(MyChan)

// }

package main
import "fmt"

func main () {
	linkchannel := make(chan string)
	go func () {
		links := []string{"https/google.com", "http://facebook.com", "www.amazon.com"}
		for _,link := range links {
			linkchannel <- link
		}
	} ()

	link1 := <- linkchannel
	link2 := <- linkchannel
	link3 := <- linkchannel

	fmt.Println(link1)
	fmt.Println(link2)
	fmt.Println(link3)
	

}