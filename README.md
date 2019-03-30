# golang-simple-queue
This is a super simple linked list queue.

## Methods

`New()` - This method creates a new instance of the queue.

```
queue := queue.New()
```

`Append()` - This method appends a new item to the queue as a string.

```
item := `{id:1, message: "Thing to work on"}`
queue := queue.Append(item)
```


`Pull()` - This method pulls an item from the queue and removes its reference.

```
data := queue.Pull()
```

`Empty()` - Checks to see if there are any messages left on the queue.


## Simple example to show usage

```
package main

import (
	"github.com/safurr/golang-simple-queue"
	"fmt"
)

type Node struct {
    data string
    next *Node
}

func main() {
	queue := queue.New()
	item := `{id:1, message: "Thing to work on"}`
	queue.Append(item)
	for queue.Empty() == false {
		data := queue.Pull()
		fmt.Printf("Data: %s\n", data)
	}
}
```
