package jwt

import (
	"github.com/alecthomas/kingpin"
	"github.com/go-chi/jwtauth"
)

type JwtModule struct {
	Secret string
	APIKey string
}

func (m *JwtModule) Configure() {
	kingpin.Flag("jwt-secret", "JWT Secret key").
		Default("secret").
		Envar("JWT_SECRET_KEY").
		StringVar(&m.Secret)
	kingpin.Flag("jwt-api-key", "JWT API key").
		Default("secret").
		Envar("JWT_API_KEY").
		StringVar(&m.APIKey)
	kingpin.Parse()
}

func (m *JwtModule) ProvideJwtModule() JwtClient {
	m.Configure()

	return &jwtClient{
		jwtAuth: jwtauth.New("HS256", []byte(m.Secret), nil),
	}
}

type JwtClient interface {
	CreateJWT(values map[string]interface{}) string
	GetJWTAuth() *jwtauth.JWTAuth
}

type jwtClient struct {
	jwtAuth *jwtauth.JWTAuth
}

func (c jwtClient) GetJWTAuth() *jwtauth.JWTAuth {
	return c.jwtAuth
}

func (c *jwtClient) CreateJWT(values map[string]interface{}) string {
	_, tokenString, _ := c.jwtAuth.Encode(values)

	return tokenString
}
