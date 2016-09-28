# WeightedChoice library for golang
### reference page
http://eli.thegreenplace.net/2010/01/22/weighted-random-generation-in-python
### how to
```
package main

import (
	wc "github.com/cxfcxf/weightedchoice"
	"fmt"
	"time"
	"math/rand"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	wc := new(wc.WeightedChoice)

	wc.Weights = []int{2, 3, 5}

	fmt.Printf("result: %d\n", wc.BinarySearch().(int))
}
```
