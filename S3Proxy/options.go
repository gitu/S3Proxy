package S3Proxy

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type options struct {
	CacheDir       string
	BindAddress    string
	Region         string
	Bucket         string
	TokenKey       string
	TokenMethod    jwt.SigningMethod
	CookieMaxAge   int
	ObjectCacheTTL time.Duration
	AwsCredentials *credentials.Credentials
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
	Options.TokenKey = "test_keys/sample_key"
	Options.TokenMethod = jwt.SigningMethodRS512
	Options.CookieMaxAge = 3600
	Options.AwsCredentials = credentials.NewSharedCredentials("", "default")
}
