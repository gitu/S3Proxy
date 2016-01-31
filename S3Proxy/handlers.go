package S3Proxy

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"io/ioutil"
	"net/http"
	"regexp"
)

var CookieStore = sessions.NewCookieStore(securecookie.GenerateRandomKey(64))

// The default handler used for everything else
func DefaultHandler(w http.ResponseWriter, req *http.Request) {
	key := req.URL.Path[1:]

	session, err := CookieStore.Get(req, "tokencache")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	formToken := req.Form.Get("token")
	if formToken != "" {
		token, err := jwt.Parse(formToken, func(t *jwt.Token) (interface{}, error) {
			return ioutil.ReadFile(Options.TokenKey)
		})
		if token.Method == Options.TokenMethod && err == nil && token.Valid {
			session.Values["path"] = token.Claims["path"]
			session.Values["client"] = token.Claims["client"]
			session.Values["exp"] = token.Claims["exp"]
			session.Save(req, w)
		} else {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
	}

	if path, ok := session.Values["path"].(string); ok {
		matched, err := regexp.MatchString(path, key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		if !matched {
			http.Error(w, "access to path not allowed", http.StatusForbidden)
			return
		}
	} else {
		http.Error(w, "path error", http.StatusForbidden)
		return
	}

	obj, s3err := S3GetObject(Options.Bucket, key, Options.Region)
	if s3err != nil {
		http.Error(w, s3err.Message, s3err.Code)
		return
	}

	http.ServeFile(w, req, obj)
	return
}
