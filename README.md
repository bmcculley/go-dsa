# Go Data Structures and Algorithms

## Example Use

```golang
package main

import (
	"fmt"
	"github.com/bmcculley/go-dsa/DataStructures/stack/llstack"
)

func main() {
	var stack llstack.Stack
	for i := 1; i < 10; i++ {
 		stack.Push(i)
 	}
 	for !stack.IsEmpty() {
		fmt.Println(stack.Pop())
	}
}
```