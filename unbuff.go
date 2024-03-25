package main

import (
	"fmt"
	"time"
)

//---send data to a channel --///

func SendData (ch chan string) {
	fmt.Println("sending data into a channel..")
	ch <- "heyllo"
	fmt.Println("data has been recieved")
}

//--- func getting data from the channel/receiving data ---///

func getData(ch chan string) {
	time.Sleep(2*time.Second)
	fmt.Println("string recieved for  processing : ",<-ch) //recieving data

}

func main() {
	ch := make(chan string)
	go SendData(ch)
	go getData(ch)
	fmt.Scanln()
}