package base

type Node struct {
	Key   string
	Value string
	Prev  *Node
	Next  *Node
}

type LruCache struct {
	count int
	size  int
	Head  *Node
	Tail  *Node
}

func NewLruCache(size int) *LruCache {
	return &LruCache{
		size:  size,
		count: 0,
	}
}

func (l *LruCache) Put(key string, value string) {
	existsNode := l.Head
	for existsNode != nil { // ищем ключ
		if existsNode.Key == key {
			existsNode.Value = value
			if existsNode != l.Head {
				if existsNode.Next != nil {
					existsNode.Next.Prev = existsNode.Prev
				}
				if existsNode.Prev != nil {
					existsNode.Prev.Next = existsNode.Next
				}
				if existsNode == l.Tail {
					l.Tail = existsNode.Prev
				}
				existsNode.Next = l.Head
				existsNode.Prev = nil
				l.Head = existsNode
			}
			return
		}
		existsNode = existsNode.Next
	}

	if l.count >= l.size {
		if l.Tail.Prev != nil {
			l.Tail.Prev.Next = nil
			l.Tail = l.Tail.Prev
		} else {
			l.Head = nil
			l.Tail = nil
		}
		l.count--
	}

	newNode := &Node{Key: key, Value: value}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		newNode.Next = l.Head
		l.Head.Prev = newNode
		newNode.Prev = nil
		l.Head = newNode
	}
	l.count++
}

func (l *LruCache) Get(key string) *string {
	firstNode := l.Head
	for firstNode != nil {
		if firstNode.Key == key {
			if firstNode == l.Head {
				return &firstNode.Value
			}
			if firstNode == l.Tail {
				l.Tail = firstNode.Prev
			}
			if firstNode.Next != nil {
				firstNode.Next.Prev = firstNode.Prev
			}
			if firstNode.Prev != nil {
				firstNode.Prev.Next = firstNode.Next
			}
			firstNode.Next = l.Head
			firstNode.Prev = nil
			l.Head.Prev = firstNode
			l.Head = firstNode
			return &firstNode.Value
		}
		firstNode = firstNode.Next
	}
	return nil
}
