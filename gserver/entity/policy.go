package entity

import (
	"time"
)

const (
	PolicyTable PolicyTableConfig = iota
	PolicyId
	PolicyName
	PolicyRoleId
	PolicyPermissionId
	PolicyUserGroupId
	PolicyCreatedAt
	PolicyUpdatedAt
)

type Policy struct {
	Id           int       `json:"id"`
	Name         string    `validate:"required"json:"name"`
	RoleId       int       `validate:"required"json:"role_id"`
	PermissionId int       `validate:"required"json:"permission_id"`
	UserGroupId  int       `validate:"required"json:"user_group_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type PolicyTableConfig int

func (pc PolicyTableConfig) String() string {
	switch pc {
	case PolicyTable:
		return "policies"
	case PolicyId:
		return "id"
	case PolicyName:
		return "name"
	case PolicyRoleId:
		return "role_id"
	case PolicyPermissionId:
		return "permission_id"
	case PolicyUserGroupId:
		return "user_group_id"
	case PolicyCreatedAt:
		return "created_at"
	case PolicyUpdatedAt:
		return "updated_at"
	}
	panic("Unknown value")
}
