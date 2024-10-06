package config

import (
	"fmt"
	"go-in-memory-cache-api/model"
	"sync"
	"time"
)

//	type ICache interface {
//		Set()
//		Get() (interface{}, bool)
//	}

type UserCacheKey int

type UserCacheNode struct {
	User model.User
	TTL  time.Time
}

type Cache struct {
	Users      map[UserCacheKey]*UserCacheNode
	Mutex      sync.RWMutex
	DefaultTTL time.Duration
}

func NewCache(defaultTTL time.Duration) *Cache {
	cache := &Cache{
		Users:      make(map[UserCacheKey]*UserCacheNode, 0),
		DefaultTTL: defaultTTL,
	}
	cache.StartCleanup()
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
	fmt.Println("reading from user cache")
	userCacheNode, ok = c.Users[userCacheKey]
	return
}

func (c *Cache) Set(userCacheKey UserCacheKey, user model.User) {
	c.Users[userCacheKey] = &UserCacheNode{
		User: user,
		TTL:  time.Now().Add(c.DefaultTTL),
	}
}
