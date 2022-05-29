package cache

import "time"

type Cache struct {
	kh map[string]cache
}

type cache struct {
	value    string
	deadline time.Time
}

func NewCache() Cache {

	return Cache{
		kh: make(map[string]cache),
	}
}

func (c Cache) Get(key string) (string, bool) {

	if a, ok := c.kh[key]; ok {
		if a.deadline.IsZero() || time.Now().Before(a.deadline) {

			return a.value, true
		}
	}
	return "", false

}

func (c *Cache) Put(key, value string) {
	c.kh[key] = cache{

		value:    value,
		deadline: time.Time{},
	}
}

func (c Cache) Keys() []string {
	res := []string{}
	for k, a := range c.kh {
		if a.deadline.IsZero() || time.Now().Before(a.deadline) {

			res = append(res, k)
		}
	}
	return res
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.kh[key] = cache{

		value:    value,
		deadline: deadline,
	}
}
