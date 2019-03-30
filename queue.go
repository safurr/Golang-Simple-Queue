package queue

type Node struct {
    data string
    next *Node
}

type Queue struct {
    head *Node
}

func (q *Queue) Append(newNode *Node) {
    if q.head == nil {
        q.head = newNode
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
