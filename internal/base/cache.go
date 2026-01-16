package base

type Node struct {
	Key   string
	Value string
	Prev  *Node
	Next  *Node
}

type LruCache struct {
	capacity map[string]*Node
	size     int
	Head     *Node
	Tail     *Node
}

func NewLruCache(size int) *LruCache {
	return &LruCache{
		size:     size,
		capacity: make(map[string]*Node),
	}
}

func (l *LruCache) Put(key string, value string) {
	count := len(l.capacity)
	existsNode := l.capacity[key] // есть ли уже такой ключ
	if existsNode != nil {        //да - обновляем и переносим в начало
		existsNode.Value = value
		if existsNode != l.Head { // нода не в начале кэша
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
		l.capacity[key] = existsNode // обновляею в мапе или не надо?
		return
	}
	if count >= l.size { // переполнился
		tailKey := l.Tail.Key
		if l.Tail.Prev != nil {
			l.Tail.Prev.Next = nil
			l.Tail = l.Tail.Prev
		} else {
			l.Tail = nil
			l.Head = nil
			delete(l.capacity, tailKey)
		}
	}
	newNode := &Node{
		Key:   key,
		Value: value,
	}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		l.capacity[key] = newNode
	} else {
		newNode.Next = l.Head
		l.Head.Prev = newNode
		newNode.Prev = nil
		l.Head = newNode
		l.capacity[key] = newNode
	}
}

func (l *LruCache) Get(key string) *string {
	value := l.capacity[key]
	if value != nil { //есть нода
		if value == l.Head {
			return &value.Value
		}
		if value == l.Tail {
			l.Tail = value.Prev
		}
		if value.Next != nil {
			value.Next.Prev = value.Prev
		}
		if value.Prev != nil {
			value.Prev.Next = value.Next
		}
		value.Next = l.Head
		value.Prev = nil
		l.Head.Prev = value
		l.Head = value
		return &value.Value
	}
	return nil
}
