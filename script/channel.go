package script

import (
	"fmt"
	"sync"
)

func ChannelMain() {
	nameChan := make(chan string)
	var wg sync.WaitGroup
	wg.Add(4)
	go newname("duy", &nameChan, &wg)
	go newname("duya", &nameChan, &wg)
	go newname("duys", &nameChan, &wg)
	go newname("duyw", &nameChan, &wg)

	fmt.Println(<-nameChan)
	fmt.Println(<-nameChan)
	fmt.Println(<-nameChan)
	fmt.Println(<-nameChan)
	wg.Wait()

}

func newname(name string, nameChan *chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	*nameChan <- name
}
