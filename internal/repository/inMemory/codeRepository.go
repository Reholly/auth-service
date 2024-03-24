package inMemory

import (
	"auth-service/internal/domain/repository"
	"github.com/patrickmn/go-cache"
	"time"
)

type Cache struct {
	cache             *cache.Cache
	defaultExpiration time.Duration
}

func NewCache(defaultExpiration, cleanupTime time.Duration) repository.CodeRepository {
	return &Cache{
		cache:             cache.New(defaultExpiration, cleanupTime),
		defaultExpiration: defaultExpiration,
	}
}

func (c *Cache) GetByUsername(username string, codeType repository.CodeType) (string, error) {
	value, found := c.cache.Get(username + string(codeType))
	if !found {
		return "", repository.ErrorNotFound
	}

	s, ok := value.(string)
	if !ok {
		return "", repository.ErrorNotFound
	}

	return s, nil
}
func (c *Cache) GetUsernameByCode(code string) (string, error) {
	value, found := c.cache.Get(code)
	if !found {
		return "", repository.ErrorNotFound
	}

	s, ok := value.(string)
	if !ok {
		return "", repository.ErrorNotFound
	}

	return s, nil
}
func (c *Cache) Remove(username string, codeType repository.CodeType) {
	c.cache.Delete(username + string(codeType))
}

func (c *Cache) Set(username, code string, codeType repository.CodeType) {
	c.cache.Set(username+string(codeType), code, c.defaultExpiration)
	c.cache.Set(code, username, c.defaultExpiration)
}
