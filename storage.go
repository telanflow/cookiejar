package cookiejar

import (
	"github.com/gomodule/redigo/redis"
	"net/http"
)

type Storage interface {
	Set(string, map[string]entry)
	Get(string) map[string]entry
	Delete(string)
}

func NewFileJar(filename string, o *Options) (http.CookieJar, error) {
	store := &FileDrive{
		filename: filename,
		entries:  make(map[string]map[string]entry),
	}
	store.readEntries()

	return New(store, o)
}

func NewEntriesJar(o *Options) (http.CookieJar, error) {
	store := &EntriesDrive{
		entries: make(map[string]map[string]entry),
	}
	return New(store, o)
}

func NewRedisJar(pool *redis.Pool, o *Options) (http.CookieJar, error) {
	store := &RedisDrive{
		pool:       pool,
		namespaces: "cookiejar",
		entries:    make(map[string]map[string]entry),
	}
	store.readEntries()
	return New(store, o)
}
