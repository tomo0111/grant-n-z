package service

import (
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/tomoyane/grant-n-z/domain/repository"
	"time"
	"github.com/satori/go.uuid"
	"github.com/dgrijalva/jwt-go"
	"github.com/tomoyane/grant-n-z/domain"
	"net/http"
	"github.com/tomoyane/grant-n-z/infra"
)

type TokenService struct {
	TokenRepository repository.TokenRepository
}

func (t TokenService) GenerateJwt(username string, userUuid uuid.UUID, isAdmin bool) string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["user_uuid"] = userUuid.String()
	claims["expires"] = time.Now().Add(time.Hour * 365).String()
	claims["role"] = "user"

	if isAdmin {
		claims["role"] = "admin"
	}

	signedToken, err := token.SignedString([]byte(infra.Yaml.App.PrivateKey))
	if err != nil {
		domain.ErrorResponse{}.Print(http.StatusInternalServerError, "failed generate jwt", "")
		return ""
	}

	return signedToken
}

func (t TokenService) ValidJwt(token string) (map[string]string, bool) {
	resultMap := map[string]string{}

	parseToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(infra.Yaml.App.PrivateKey), nil
	})

	claims := parseToken.Claims.(jwt.MapClaims)
	if err != nil && !parseToken.Valid {
		return resultMap, false
	}

	resultMap["username"] = claims["username"].(string)
	resultMap["user_uuid"] = claims["user_uuid"].(string)
	resultMap["expires"] = claims["expires"].(string)
	resultMap["role"] = claims["role"].(string)

	return resultMap, true
}

func (t TokenService) GetTokenByUserId(userId string) *entity.Token {
	return t.TokenRepository.FindByUserId(userId)
}

func (t TokenService) InsertToken(userUuid uuid.UUID, token string, refreshToken string) *entity.Token {
	data := entity.Token{
		TokenType: "Bearer",
		Token: token,
		RefreshToken: refreshToken,
		UserUuid: userUuid,
	}
	return t.TokenRepository.Save(data)
}