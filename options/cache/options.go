package main

import (
	"fmt"
	"log"
	"time"
)

type Option func(c *Cache) error

func CacheMaxAge(d time.Duration) Option {
	return func(c *Cache) error {
		return nil
	}
}

func CacheMaxEntries(s int) Option {
	return func(c *Cache) error {
		return nil
	}
}

type Cache struct {
}

func NewCache(opts ...Option) (*Cache, error) {
	c := &Cache{}
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, fmt.Errorf("error while applying option: %w", err)
		}
	}

	return c, nil
}

func main() {
	c, err := NewCache(
		CacheMaxAge(1*time.Hour),
		CacheMaxEntries(10),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(c)
}
