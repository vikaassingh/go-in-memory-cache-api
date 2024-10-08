package cache

import (
	"go-in-memory-cache-api/model"
	"sync"
	"time"
)

var once sync.Once
var cache *Cache

type UserCacheKey uint

type UserCacheNode struct {
	User model.User
	TTL  time.Time
}

type Cache struct {
	Users      map[UserCacheKey]*UserCacheNode
	Mutex      sync.RWMutex
	DefaultTTL time.Duration
}

func init() {
	cache = GetCache(time.Second * 10)
}

func GetCache(defaultTTL time.Duration) *Cache {
	if cache == nil {
		once.Do(func() {
			cache = &Cache{
				Users:      make(map[UserCacheKey]*UserCacheNode, 0),
				DefaultTTL: defaultTTL,
			}
			go cache.StartCleanup()
		})
	}
	return cache
}

func (c *Cache) StartCleanup() {
	tricker := time.NewTicker(c.DefaultTTL)
	for {
		select {
		case <-tricker.C:
			c.CleanCache()
		}
	}
}

func (c *Cache) CleanCache() {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	for userCacheKey, userCacheNode := range c.Users {
		if time.Now().After(userCacheNode.TTL) {
			delete(c.Users, userCacheKey)
		}
	}
}

func (c *Cache) Get(userCacheKey UserCacheKey) (userCacheNode *UserCacheNode, ok bool) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	userCacheNode, ok = c.Users[userCacheKey]
	return
}

func (c *Cache) Set(userCacheKey UserCacheKey, user model.User) {
	// fmt.Printf("\nBefore Set c:%p-%v", c, c)
	if _, ok := c.Users[userCacheKey]; !ok {
		c.Users[userCacheKey] = &UserCacheNode{
			User: user,
			TTL:  time.Now().Add(c.DefaultTTL),
		}
	}
	// fmt.Printf("\nAfter Set cache:%p-%v", cache, cache)
	// fmt.Printf("\nAfter Set c:%p-%v", c, c)
}
