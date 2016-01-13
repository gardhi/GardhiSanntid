package main

import(
	"fmt"
	"runtime"
)

var i = 0

func inc(quit chan int){
	for x := 0; x< 1000000; x++{
		i++
		
	}

	quit <- 1
}

func dec(quit chan int){
	for x := 0; x< 1000000; x++{
		i--
		
	}

	quit <- 1
}

func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())
	routineQuit := make(chan int)
	go inc(routineQuit)
	go dec(routineQuit)
	<- routineQuit
	<- routineQuit

	fmt.Println("Done, result = ", i)

}
	
