package oauth

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/flaxius/crud-cas/pkg/pb/authentication"
	"sync"
	"github.com/flaxius/crud-cas/pkg/cassandra"
	cascfg "github.com/flaxius/crud-cas/pkg/cassandra/config"
	"time"
)
const (
	insertServiceName = `INSERT INTO service_names(service_name) VALUES (?)`
	queryServiceNames = `SELECT user_code,user_jwk_enc FROM userstore WHERE user_id=9495`
)

// Span is the database representation of a span.
type UserStore struct {
	user_code      string // deprecated
	user_jwk_enc string
}

type serviceTokenServer struct {
	oauth *oauth.User
	m     sync.Mutex
	session cassandra.Session
	queryServiceNames string
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

func (s *serviceTokenServer) Create(ctx context.Context, req *oauth.UserStoreIdReq) (*oauth.UserStoreIdRes, error) {
	s.m.Lock()
	defer s.m.Unlock()
	cConfig := &cascfg.Configuration{
		Servers:              []string{"22.6.4.241"},
		Keyspace:             "securestore",
		ConnectionsPerHost:   10,
		Timeout:              time.Millisecond * 750,
		ProtoVersion:         4,
		Port:9042,
	}
	cqlSession, _ := cConfig.NewSession()
	papa := cqlSession.Query(queryServiceNames).Iter()
	var user_code string
	var user_jwk_enc string
	var operations []UserStore
	for papa.Scan(&user_code,&user_jwk_enc) {
		//operations = append(operations, operation)
		dbSpan := UserStore{
			user_code:    user_code,
			user_jwk_enc: user_jwk_enc,
		}
		operations = append(operations,dbSpan)
	}
	papa.Close()
	fmt.Println(operations)
	return &oauth.UserStoreIdRes{
		UserId:        1,
		UserJwkEnc:"A",
		UserJwkSig:"B",
	}, nil
}
