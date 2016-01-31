package S3Proxy

import (
	"encoding/json"
	"time"
)

var s3Objects = map[string]*S3ObjectCacheItem{}

type s3CacheItem struct {
	Timestamp time.Time
	TTL       time.Duration
}

func (c s3CacheItem) String() string {
	tmp := &c
	tmp.TTL = c.TimeLeft()
	out, _ := json.Marshal(tmp)
	return string(out)
}

// TimeLeft returns the time duration before the CacheItem expires
func (c s3CacheItem) TimeLeft() time.Duration {
	return c.TTL - time.Since(c.Timestamp)
}

// S3ObjectCacheItem represents a cache entry for meta data for an S3 object
// and not the S3Object itself. The idea is to reduce S3 API calls for Object
// information
type S3ObjectCacheItem struct {
	CacheItem s3CacheItem
	Key       string
	Bucket    string
	FilePath  string
}

func (o S3ObjectCacheItem) String() string {
	out, _ := json.Marshal(o)
	return string(out)
}

// CacheObjectGet checks if the S3Object information is in the local cache
func CacheObjectGet(key string) *S3ObjectCacheItem {
	object, hit := s3Objects[key]
	if hit {
		if time.Since(object.CacheItem.Timestamp) <= object.CacheItem.TTL {
			LogInfo("S3 Object Cache Hit - " + object.String())
			return object
		}
		// The cache entry has expired so remove it
		delete(s3Objects, key)
	}
	LogInfo("S3 Object Cache Miss - {\"Key\":\"" + key + "\"}")
	return nil
}

// CacheObjectSet enters a S3 Object's info into the cache
func CacheObjectSet(key, bucket, filepath string) *S3ObjectCacheItem {
	object := S3ObjectCacheItem{
		CacheItem: s3CacheItem{Timestamp: time.Now(), TTL: Options.ObjectCacheTTL},
		Key:       key,
		Bucket:    bucket,
		FilePath:  filepath,
	}
	s3Objects[key] = &object
	LogInfo("S3 Object Cache Set - " + object.String())
	return &object
}
