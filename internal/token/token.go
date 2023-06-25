package token

import (
	"github.com/kataras/jwt"
	"github.com/lishimeng/app-starter/token"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/passport/internal/etc"
	"time"
)

func GenToken(org, uid, client string) (t string, err error) {
	var payload = token.JwtPayload{
		Org:    org,
		Uid:    uid,
		Client: client,
	}
	key := []byte(etc.Config.Token.Key)
	var issuer = etc.Config.Token.Issuer
	var alg = "HS256"
	var exp = time.Hour * 24 * 30 * 12 * 10
	var provider = token.NewJwtProvider(issuer,
		token.WithAlg(alg),
		token.WithKey(key, key),
		token.WithDefaultTTL(exp))
	tk, err := provider.Gen(payload)
	if err != nil {
		return
	}
	t = string(tk)
	return
}

func VerifyToken(t string) (to *jwt.VerifiedToken, err error) {
	key := []byte(etc.Config.Token.Key)
	var issuer = etc.Config.Token.Issuer
	var alg = "HS256"
	verifyProvider := token.NewJwtProvider(issuer,
		token.WithAlg(alg),
		token.WithKey(nil, key),
	)
	to, err = verifyProvider.Verify([]byte(t))
	log.Debug("to :%s", to)
	log.Debug("to :%s", string(to.Payload))
	if err != nil {
		return
	}
	return
}
