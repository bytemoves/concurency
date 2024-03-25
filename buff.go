package main
 
import (
   "fmt"
)
 
func main() {
 
   linkChannel := make(chan string, 5)
   go ping(linkChannel)
 
   fmt.Println(<-linkChannel)
   fmt.Println(<-linkChannel)
   fmt.Println(<-linkChannel)
 
}
 
func ping(c chan string) {
   links := []string{"https://www.golinuxcloud.com/", "https://www.tesla.com/", "https://www.google.com/"}
   for _, link := range links {
       c <- link
   }
}
