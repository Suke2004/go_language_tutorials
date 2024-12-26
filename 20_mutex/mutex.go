//If a single resouce is modified by multiple processes then it is not atomic,some tries to over write other
package main

import(
	"fmt"
	"sync"
)

type post struct{
	views int
	mu sync.Mutex//it is add to add mutex
}

// func(p *post) inc(wg *sync.WaitGroup){
// 	defer wg.Done()
// 	p.views +=1
// }//Due to sumultaneos modification of 2 or more processes we get error in the value.Mutex comes to play here

func(p *post) inc(wg *sync.WaitGroup){
	defer func(){
		p.mu.Unlock()
		wg.Done()//Using unclock in defer function is best practice as if in place of increment there may be any other modication logic which may throw an error.Then we may not rech to unlock.To prevent it we use in defer fn
	}()
	p.mu.Lock()
	p.views +=1

	// defer wg.Done()
	// p.mu.Lock()
	// p.views+=1
	// p.mu.Unlock()
}

func main(){
	var wg sync.WaitGroup
	myPost := post{
		views: 0,
	}
	for i:=0;i<1000;i++{
		wg.Add(1)
		go myPost.inc(&wg)
	}
	wg.Wait()
	// myPost.inc()
	// myPost.inc()
	fmt.Println(myPost.views)
} 