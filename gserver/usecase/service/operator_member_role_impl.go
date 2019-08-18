package service

import (
	"strconv"
	"strings"

	"github.com/tomoyane/grant-n-z/gserver/common/driver"
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"

	"github.com/tomoyane/grant-n-z/gserver/usecase/repository"
)

type operatorMemberRoleServiceImpl struct {
	operatorMemberRoleRepository repository.OperatorMemberRoleRepository
	userRepository               repository.UserRepository
	roleRepository               repository.RoleRepository
}

func NewOperatorMemberRoleService() OperatorMemberRoleService {
	log.Logger.Info("Inject `operatorMemberRoleRepository`,`userRepository`,`roleRepository` to `OperatorMemberRoleService`")
	return operatorMemberRoleServiceImpl{
		operatorMemberRoleRepository: repository.NewOperatorMemberRoleRepository(driver.Db),
		userRepository:               repository.NewUserRepository(driver.Db),
		roleRepository:               repository.NewRoleRepository(driver.Db),
	}
}

func (omrs operatorMemberRoleServiceImpl) Get(queryParam string) ([]*entity.OperatorMemberRole, *model.ErrorResponse) {
	if strings.EqualFold(queryParam, "") {
		return omrs.GetAll()
	}

	i, castErr := strconv.Atoi(queryParam)
	if castErr != nil {
		log.Logger.Warn("The user_id is only integer")
		return nil, model.BadRequest(castErr.Error())
	}

	entities, err := omrs.GetByUserId(i)
	if err != nil {
		return nil, err
	}

	if entities == nil {
		return []*entity.OperatorMemberRole{}, nil
	}

	return entities, nil
}

func (omrs operatorMemberRoleServiceImpl) GetAll() ([]*entity.OperatorMemberRole, *model.ErrorResponse) {
	return omrs.operatorMemberRoleRepository.FindAll()
}

func (omrs operatorMemberRoleServiceImpl) GetByUserId(userId int) ([]*entity.OperatorMemberRole, *model.ErrorResponse) {
	return omrs.operatorMemberRoleRepository.FindByUserId(userId)
}

func (omts operatorMemberRoleServiceImpl) GetRoleNameByUserId(userId int) ([]string, *model.ErrorResponse) {
	return nil, nil
}

func (omrs operatorMemberRoleServiceImpl) Insert(entity *entity.OperatorMemberRole) (*entity.OperatorMemberRole, *model.ErrorResponse) {
	if userEntity, _ := omrs.userRepository.FindById(entity.UserId); userEntity == nil {
		log.Logger.Warn("Not found user id")
		return nil, model.BadRequest("Not found user id")
	}

	if roleEntity, _ := omrs.roleRepository.FindById(entity.RoleId); roleEntity == nil {
		log.Logger.Warn("Not found role id")
		return nil, model.BadRequest("Not found role id")
	}

	return omrs.operatorMemberRoleRepository.Save(*entity)
}
