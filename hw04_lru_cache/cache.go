package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

// Добавить значение в кэш по ключу.

func (c *lruCache) Set(key Key, value interface{}) bool {
	v, exists := c.items[key]

	if exists && v != nil {
		// если элемент присутствует в словаре, то обновить его значение и переместить элемент в начало очереди.
		v.Value = value
		c.items[key] = v       // Добавить в словарь новое значение.
		c.queue.MoveToFront(v) //  Переместить в начало очереди.
		return true
	}
	c.queue.PushFront(value)       //  Добавить в начало очереди.
	c.items[key] = c.queue.Front() // Добавить в словарь.

	// если размер очереди больше ёмкости кэша, то необходимо
	// удалить последний элемент из очереди и его значение из словаря).
	if c.queue.Len() > c.capacity {
		// Find the key which have to be delete.
		var rmKey Key
		for k1, v1 := range c.items {
			if v1.Value == c.queue.Back().Value || (v1.Prev == c.queue.Back().Prev && v1.Next == c.queue.Back().Next) {
				rmKey = k1
				break
			}
		}
		delete(c.items, rmKey)         // удалить ключ из словаря.
		c.queue.Remove(c.queue.Back()) // удалить последний элемент из очереди.
	}
	return false
}

// Получить значение из кэша по ключу.
func (c *lruCache) Get(key Key) (interface{}, bool) {
	v, exists := c.items[key]
	if !exists {
		return nil, false
	}
	if v.Value != c.queue.Front().Value {
		c.queue.MoveToFront(v)
	}
	return v.Value, true
}

func (c *lruCache) Clear() {
	c.items = make(map[Key]*ListItem, 0) // Удалить все значения из словаря.

	for iter := 0; iter < c.queue.Len(); iter++ { // Удалить все элементы кэша.
		c.queue.Remove(c.queue.Front())
	}
}
