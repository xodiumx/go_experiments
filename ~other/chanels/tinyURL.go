package main

import (
	"crypto/rand"
	"encoding/base64"
	"sync"
)

type Codec struct {
	mu         sync.Mutex
	urlMap     map[string]string
	keySize    int
	tinyDomain string
}

func Constructor() Codec {
	return Codec{
		urlMap:     make(map[string]string),
		keySize:    6,
		tinyDomain: "http://tinyurl.com/",
	}
}

func (c *Codec) generateKey() string {
	buf := make([]byte, c.keySize)
	_, err := rand.Read(buf)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(buf)[:8]
}

func (c *Codec) encode(longUrl string) string {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, url := range c.urlMap {
		if url == longUrl {
			return c.tinyDomain + key
		}
	}

	key := c.generateKey()
	c.urlMap[key] = longUrl
	return c.tinyDomain + key
}

func (c *Codec) decode(shortUrl string) string {
	c.mu.Lock()
	defer c.mu.Unlock()

	key := shortUrl[len(c.tinyDomain):]
	return c.urlMap[key]
}

//func main() {
//	obj := Constructor()
//	url := obj.encode("https://leetcode.com/problems/design-tinyurl")
//	ans := obj.decode(url)
//	fmt.Println(ans)
//}
