package main

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup){//We are using point so that by function we are changing original variable
	//defer function delays the execution of a function until the sorrounding function has finished
	defer w.Done()//Done will be called when the goroutine is finished,it decrements variable by 1
	fmt.Println("doint task", id)
}

func main(){
	var wg sync.WaitGroup//We have to syncronize the tasks

	for i := 0;i<10;i++{
		wg.Add(1)//This is add 1 to the variable wg
		go task(i,&wg)
	}

	wg.Wait()//this will wait until variable is 0
}