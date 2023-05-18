package test

import (
	"flag"
	"testing"
	"time"

	"github.com/Wishrem/wuso/config"
	"github.com/Wishrem/wuso/pkg/errno"
	. "github.com/Wishrem/wuso/pkg/utils/jwt"
	"github.com/golang-jwt/jwt"
	. "github.com/smartystreets/goconvey/convey"
)

var path string

func init() {
	path = *flag.String("config", ".", "test config path")
}

func TestJWT(t *testing.T) {
	config.Init(path)
	Convey("Test token parsing", t, func() {
		Convey("Test correct parsing", func() {
			token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ3aXNocmVtIiwidXNlcl9pZCI6MTUxNjIzOTAyMiwiZXhwIjo0ODM5OTY5NjkwfQ.I-4j3lGLz6M2nQQebsPplEjUPa95tnpxiWA_941hINU"
			ac := Claims{
				UserId: 1516239022,
				StandardClaims: jwt.StandardClaims{
					Issuer:    "wishrem",
					ExpiresAt: 4839969690, // 2123
				},
			}
			c, err := Parse(token, config.JWT.Secret)
			Convey("Should not be error", func() {
				So(nil, ShouldEqual, err)
			})
			So(ac, ShouldResemble, *c)
		})

		Convey("Test wrong Issuer", func() {
			token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ0ZXN0IiwidXNlcl9pZCI6MTUxNjIzOTAyMiwiZXhwIjo0ODM5OTY5NjkwfQ.L3MuhUHVZTE9vxGtbZIl48nWOmmPaBXTYKB9nTnccFA"
			_, err := Parse(token, config.JWT.Secret)
			So(errno.AuthorizationFailed, ShouldBeError, err)
		})

		Convey("Test expired token", func() {
			token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ3aXNocmVtIiwidXNlcl9pZCI6MTUxNjIzOTAyMiwiZXhwIjowfQ.1jmx63v_qWISiQ1-DWgUVeS8SJ_avIK0Mufs1DEMXFg"
			_, err := Parse(token, config.JWT.Secret)
			So(errno.AuthorizationExpired, ShouldBeError, err)
		})
	})

	Convey("Test token generation", t, func() {
		s, err := Generate(1516239022, config.JWT.Secret)
		ac := &Claims{
			UserId: 1516239022,
			StandardClaims: jwt.StandardClaims{
				Issuer:    "wishrem",
				ExpiresAt: time.Now().Add(config.JWT.ExpireTime).Unix(),
			},
		}
		Convey("Should not be error", func() {
			So(nil, ShouldEqual, err)
		})

		Convey("Should be same", func() {
			c, _ := Parse(s, config.JWT.Secret)
			So(ac, ShouldResemble, c)
		})
	})
}
