package lock

import (
	"fmt"
	"sync"
)

func deferred() {
	m := sync.Mutex{}
	m.Lock()
	defer m.Unlock()

	fmt.Println("the lock will be unlocked via defer")
}
