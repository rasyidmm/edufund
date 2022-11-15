package jwt_builder

import (
	"edufund/src/shared/tracing"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"time"
)

type TokenRequest struct {
	Username string
	Fullname string
}

type TokenResponse struct {
	Token   string
	Expired string
}
type JwtCustomClaims struct {
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	UUID     string `json:"uuid"`
	jwt.StandardClaims
}

var SecretKey = []byte("a6a72f7f-2d25-4e54-9758-74d70c040b46")

func CreateJwtToken(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "CreateJwtToken")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqdata := in.(*TokenRequest)

	Expired := time.Now().Add(5 * time.Hour)

	claims := JwtCustomClaims{
		Fullname: reqdata.Fullname,
		Username: reqdata.Username,
		UUID:     uuid.NewString(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: Expired.Unix(),
			NotBefore: time.Now().Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString(SecretKey)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	res := &TokenResponse{
		Token:   t,
		Expired: Expired.Format("2006-01-02 15:04:05"),
	}

	tracing.LogResponse(sp, res)
	return res, nil
}
