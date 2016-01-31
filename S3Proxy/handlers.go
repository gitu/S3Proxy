package S3Proxy

import (
	"net/http"
//	"github.com/dgrijalva/jwt-go"
)


// The default handler used for everything else
func DefaultHandler(w http.ResponseWriter, req *http.Request) {
	key := req.URL.Path[1:]


	obj, err := S3GetObject(Options.Bucket, key, Options.Region)
	if err != nil {
		http.Error(w, err.Message, err.Code)
		return
	}

	http.ServeFile(w, req, obj)
	return
}
