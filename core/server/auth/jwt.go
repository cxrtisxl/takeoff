package auth

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtPrivateKey *ecdsa.PrivateKey
var jwtPublicKey *ecdsa.PublicKey
var kid string

type JWK struct {
	Kty string `json:"kty"`
	Use string `json:"use"`
	Alg string `json:"alg"`
	Kid string `json:"kid"`
	Crv string `json:"crv"`
	X   string `json:"x"`
	Y   string `json:"y"`
}

type JWKS struct {
	Keys []JWK `json:"keys"`
}

var cachedJWKS []byte

func initJWT() {
	var err error

	// Getting private key from JWT_KEY_B64
	pemBytes, err := base64.StdEncoding.DecodeString(os.Getenv("JWT_KEY_B64"))
	if err != nil {
		panic(err)
	}
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		panic("failed to parse PEM")
	}
	jwtPrivateKey, err = x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	// Setting public key
	jwtPublicKey = &jwtPrivateKey.PublicKey

	// Generating kid (key ID)
	bytes, err := x509.MarshalPKIXPublicKey(jwtPublicKey)
	if err != nil {
		panic(err)
	}
	hash := sha256.Sum256(bytes)
	kid = hex.EncodeToString(hash[:16])

	// Preparing JWKS
	xStr := base64.RawURLEncoding.EncodeToString(jwtPublicKey.X.Bytes())
	yStr := base64.RawURLEncoding.EncodeToString(jwtPublicKey.Y.Bytes())
	jwks := JWKS{Keys: []JWK{
		{
			Kty: "EC",
			Use: "sig",
			Alg: "ES256",
			Kid: kid,
			Crv: "P-256",
			X:   xStr,
			Y:   yStr,
		},
	}}
	cachedJWKS, err = json.Marshal(jwks)
	if err != nil {
		panic(err)
	}
}

func NewJWT(pk *ecdsa.PrivateKey) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	tokenString, err := token.SignedString(pk)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyJWT(tokenString string) {
	token, err := jwt.Parse(tokenString,
		func(_ *jwt.Token) (any, error) {
			return jwtPublicKey, nil
		},
		jwt.WithValidMethods([]string{jwt.SigningMethodES256.Alg()}),
	)
	if err != nil {
		log.Fatal(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		log.Println(claims["foo"], claims["nbf"])
	} else {
		log.Println(err)
	}
}
