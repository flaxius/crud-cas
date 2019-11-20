package oauth

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/flaxius/portia/pkg/pb/authentication"
	"sync"
	"time"
)

type serviceTokenServer struct {
	oauth *oauth.User
	m     sync.Mutex
}

var jwtKey = []byte("clave_de_maxima_seguridad_XD")

//NewTokenService aa
func NewTokenService() oauth.Oauth2ServiceServer {
	return &serviceTokenServer{}
}

//Claims FLA
type Claims struct {
	Username string `json:"username"`
	admin    bool   `json:admin`
	jwt.StandardClaims
}

func (s *serviceTokenServer) Create(ctx context.Context, req *oauth.CreateRequest) (*oauth.CreateResponse, error) {
	s.m.Lock()
	defer s.m.Unlock()
	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute)
	//var date &oauth.User{}

	// Create the JWT claims, which includes the username and expiry time
	var claims = &Claims{
		"XE77772",
		false,
		jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			Audience:  "Novatos",
			Issuer:    "Dios",
			Subject:   "Pruebas",
			ExpiresAt: expirationTime.Unix(),
		},
	} // Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		fmt.Println("AAA")
	}
	return &oauth.CreateResponse{
		Type:        "Bearer",
		AccessToken: tokenString,
	}, nil
}
