package S3Proxy

import "time"

type options struct {
	CacheDir       string
	BindAddress    string
	Region         string
	Bucket         string
	Key            string
	ObjectCacheTTL time.Duration
}

// Options is a globally available struct for storing runtime options
var Options = options{}

// LoadDefaultOptions is used to load in the default options into the globally
// available options struct. This is typically one of the first things called
// on start up. After this all options can be overridden
func LoadDefaultOptions() {
	Options.CacheDir = "/tmp/S3Proxy/"
	Options.BindAddress = ":9090"
	Options.ObjectCacheTTL = time.Duration(1 * time.Minute)
	Options.Region = "eu-west-1"
	Options.Bucket = "example-bucket"
}
