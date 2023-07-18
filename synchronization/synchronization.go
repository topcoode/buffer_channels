package synchronization

import (
	"fmt"
	"time"
)

func check(done chan bool) {
	fmt.Print("Welcome to...")
	time.Sleep(time.Second)
	fmt.Println("TutorialsPoint")

	done <- true
	fmt.Println("done-------->", done) //done---------> 0xc00005e070
}

func main() {
	done := make(chan bool, 1)
	go check(done)

	<-done
}

//output:=
//Welcome to...TutorialsPoint
