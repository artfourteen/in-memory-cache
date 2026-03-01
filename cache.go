package cache

type ICache interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
	Delete(key string)
}

type Cache struct {
	memory map[string]interface{}
}

func (c *Cache) Set(key string, value interface{}) {
	c.memory[key] = value
}

func (c *Cache) Get(key string) (interface{}, bool) {
	value, ok := c.memory[key]
	return value, ok
}

func (c *Cache) Delete(key string) {
	delete(c.memory, key)
}

func New() ICache {
	return &Cache{memory: make(map[string]interface{})}
}
