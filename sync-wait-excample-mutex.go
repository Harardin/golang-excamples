/* RWMutex or Mutex
Allows to stop and continue in case you work with slice from different goroutines
*/

// Excample

package main

import (
	"sync"
)

func main() {
  var rw sync.RWMutex
  
  rw.Lock()
    // DO SOMETHING
    // This one always stop
  rw.Unlock()
  
  rw.RLock()
    // Read from slice for excample
    // This one not always stop
  rw.RUnlock()
}
