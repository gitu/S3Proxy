package S3Proxy

import (
	"testing"
	"time"
)

func TestCacheObjectGet(t *testing.T) {
	// Needed for Options.ObjectCacheTTL
	LoadDefaultOptions()
	// Lower the TTL for testing purposes
	Options.ObjectCacheTTL = time.Duration(2 * time.Second)
	// Enter an Object item into the cache
	CacheObjectSet("object_test_add", "bucket", "/file/path")
	// Retrieve the object item from the cache
	object := CacheObjectGet("object_test_add")
	if object != nil {
		if object.Key != "object_test_add" {
			t.Error(
				"For", "object.Key",
				"Expected", "object_test_add",
				"Got", object.Key,
			)
		}
		if object.Bucket != "bucket" {
			t.Error(
				"For", "object.Bucket",
				"Expected", "bucket",
				"Got", "object.Bucket",
			)
		}
		if object.FilePath != "/file/path" {
			t.Error(
				"For", "object.FilePath",
				"Expected", "bucket",
				"Got", "/file/path",
			)
		}
	} else {
		t.Error(
			"For", "object",
			"Expect", "SS3ObjectCacheItem",
			"Got", nil,
		)
	}
}

func TestCacheObjectExpire(t *testing.T) {
	// Needed for Options.ObjectCacheTTL
	LoadDefaultOptions()
	// Lower the TTL for testing purposes
	Options.ObjectCacheTTL = time.Duration(2 * time.Second)
	// Enter an Object item into the cache
	CacheObjectSet("object_test_expire", "bucket", "/file/path")
	// Sleep long enough for the entry to expire
	time.Sleep(Options.ObjectCacheTTL + time.Second)
	object := CacheObjectGet("object_test_expire")
	// Confirm that the entry expired
	if object != nil {
		t.Error(
			"For", "CacheObjectGet",
			"Expected", "nil",
			"Got", object,
		)
	}
}
