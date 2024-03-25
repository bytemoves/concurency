package main
import (
	"log"
	"os"
)

func main () {
	file , err := os.Open("sloth.txt")
	if err != nil {
		log.Fatal("error in opening the file: %v",err)
	
	}

	defer file.Close()

	byteReads := make([]byte,33)
	n,err := file.Read(byteReads)
	if err != nil {
		log.Fatalf("error from reading the file sloth.txt : %v",err)
	}

	log.Printf("We read \"%s\" into bytesRead (%d bytes)",string(byteReads),n)
        

}