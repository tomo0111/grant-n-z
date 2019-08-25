package service

import (
	"strings"

	"github.com/tomoyane/grant-n-z/gserver/common/property"
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

var tsInstance TokenService

type tokenServiceImpl struct {
	userService               UserService
	operatorMemberRoleService OperatorMemberRoleService
}

func GetTokenServiceInstance() TokenService {
	if tsInstance == nil {
		tsInstance = NewTokenService()
	}
	return tsInstance
}

func NewTokenService() TokenService {
	log.Logger.Info("New `TokenService` instance")
	log.Logger.Info("Inject `UserService`, `OperatorMemberRoleService` to `TokenService`")
	return tokenServiceImpl{
		userService:               NewUserService(),
		operatorMemberRoleService: NewOperatorMemberRoleService(),
	}
}

func (tsi tokenServiceImpl) Generate(queryParam string, userEntity entity.User) (*string, *model.ErrorResponse) {
	if strings.EqualFold(queryParam, property.AuthOperator) {
		return tsi.operatorToken(userEntity)
	} else if strings.EqualFold(queryParam, "") {
		return tsi.userToken(userEntity)
	} else {
		return nil, model.BadRequest("Not support type of query parameter")
	}
}

func (tsi tokenServiceImpl) operatorToken(userEntity entity.User) (*string, *model.ErrorResponse) {
	// TODO: Cache

	user, err := tsi.userService.GetUserWithRoleByEmail(userEntity.Email)
	if err != nil || user == nil {
		return nil, model.BadRequest("Failed to email or password")
	}
	if !tsi.userService.ComparePw(user.Password, userEntity.Password) {
		return nil, model.BadRequest("Failed to email or password")
	}
	if user.RoleId != property.OperatorRoleId {
		return nil, model.BadRequest("Can not issue token")
	}

	return tsi.userService.GenerateJwt(&userEntity, property.OperatorRoleId), nil
}

func (tsi tokenServiceImpl) serviceToken(userEntity entity.User) (*string, *model.ErrorResponse) {
	return nil, nil
}

func (tsi tokenServiceImpl) userToken(userEntity entity.User) (*string, *model.ErrorResponse) {
	// TODO: Cache
	// TODO: Set user policy

	user, err := tsi.userService.GetUserByEmail(userEntity.Email)
	if err != nil || user == nil {
		return nil, model.BadRequest("Failed to email or password")
	}

	if !tsi.userService.ComparePw(user.Password, userEntity.Password) {
		return nil, model.BadRequest("Failed to email or password")
	}

	return tsi.userService.GenerateJwt(&userEntity, property.OperatorRoleId), nil
}