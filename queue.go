package queue

type Node struct {
    data string
    next *Node
}

type Queue struct {
    head *Node
}

func (q *Queue) Append(input string) {
    newNode := &Node{data: input}
    if q.head == nil {
        q.head = &Node
        return
    }

    currentNode := q.head
    for currentNode.next != nil {
        currentNode = currentNode.next
    }
    currentNode.next = newNode
}

func (q *Queue) Pull() string {
    if q.head == nil {
        return ""
    }
    data := q.head.data
    q.head = q.head.next
    return data
}

func (q *Queue) Size() int {
    size := len(q)
    return size
}
