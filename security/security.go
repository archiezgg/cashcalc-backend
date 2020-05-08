package security

import (
	"time"

	"github.com/IstvanN/cashcalc-backend/models"
	jwt "github.com/dgrijalva/jwt-go"
)

var signingKey = []byte("rekettye")

// Claims is the struct for the Token Claims including role
// and standard JWT claims
type Claims struct {
	Role models.Role `json:"role"`
	jwt.StandardClaims
}

// CreateToken takes a Role as param and creates a signed token
func CreateToken(role models.Role) (string, error) {
	claims := Claims{
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
