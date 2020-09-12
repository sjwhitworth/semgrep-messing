package lock

import (
	"fmt"
	"sync"
)

func explicit() {
	m := sync.Mutex{}
	m.Lock()
	fmt.Println("the lock will be unlocked explicitly")
	m.Unlock()
}
