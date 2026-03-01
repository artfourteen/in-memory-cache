package cache

type ICache interface {
	Set(key string, value interface{})
	Get(key string) interface{}
	Delete(key string)
}

type Cache struct {
	memory map[string]interface{}
}

func (c *Cache) Set(key string, value interface{}) {
	c.memory[key] = value
}

func (c *Cache) Get(key string) interface{} {
	return c.memory[key]
}

func (c *Cache) Delete(key string) {
	delete(c.memory, key)
}

func New() *Cache {
	return &Cache{memory: make(map[string]interface{})}
}
