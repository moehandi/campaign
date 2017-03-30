package common

import (
	"crypto/rsa"
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gorilla/context"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"github.com/Sirupsen/logrus"
)

// AppClaims provides custom claim for JWT
type AppClaims struct {
	UserName string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

// using asymmetric crypto/RSA keys
// location of private/public key files
const (
	// openssl gen rsa -out app.rsa 1024
	privKeyPath = "keys/app.rsa"
	// openssl rsa -in app.rsa -pubout > app.rsa.pub
	pubKeyPath = "keys/app.rsa.pub"
)

// Private key for signing and public key for verification
var (
	//verifyKey, signKey []byte
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

// Read the key files before starting http handlers
func initKeys() {

	signBytes, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		logrus.Fatalf("[initKeys]: %s\n", err)
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		logrus.Fatalf("[initKeys]: %s\n", err)
	}

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		logrus.Fatalf("[initKeys]: %s\n", err)
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		logrus.Fatalf("[initKeys]: %s\n", err)
	}
}

// GenerateJWT generates a new JWT token
func GenerateJWT(name, role string) (string, error) {
	// Create the Claims
	claims := AppClaims{
		name,
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 365).Unix(),
			Issuer:    "moehandi",
			//Audience: "",
			//Id: "",
			//IssuedAt: 28839,
			//NotBefore: int64,
			//Subject:   "subject",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	ss, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}

	return ss, nil
}

// Authorize Middleware for validating JWT tokens
func Authorize(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	// Get token from request
	token, err := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, &AppClaims{}, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return verifyKey, nil
	})

	if err != nil {
		switch err.(type) {

		case *jwt.ValidationError: // JWT validation error
			vErr := err.(*jwt.ValidationError)

			switch vErr.Errors {
			case jwt.ValidationErrorExpired: //JWT expired
				ShowAppError(
					w,
					err,
					"Access Token is expired, get a new Token",
					401,
				)
				return

			default:
				ShowAppError(w,
					err,
					"Error while parsing the Access Token!",
					500,
				)
				return
			}

		default:
			ShowAppError(w,
				err,
				"Error while parsing Access Token!",
				500)
			return
		}

	}
	if token.Valid {
		// Set user name to HTTP context
		context.Set(r, "user", token.Claims.(*AppClaims).UserName)
		next(w, r)
	} else {
		ShowAppError(
			w,
			err,
			"Invalid Access Token",
			401,
		)
	}
}

// TokenFromAuthHeader is a "TokenExtractor" that takes a given request and extracts
// the JWT token from the Authorization header.
func TokenFromAuthHeader(r *http.Request) (string, error) {
	// Look for an Authorization header
	if ah := r.Header.Get("Authorization"); ah != "" {
		// Should be a bearer token
		if len(ah) > 6 && strings.ToUpper(ah[0:6]) == "BEARER" {
			return ah[7:], nil
		}
	}
	return "", errors.New("No token in the HTTP request")
}
