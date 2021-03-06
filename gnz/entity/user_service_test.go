package entity

import (
	"strings"
	"testing"
)

func TestUserServiceString(t *testing.T) {
	table := UserServiceTable.String()
	if !strings.EqualFold(table, "user_services") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	id := UserServiceId.String()
	if !strings.EqualFold(id, "id") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	internalId := UserServiceInternalId.String()
	if !strings.EqualFold(internalId, "internal_id") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	userId := UserServiceUserUuid.String()
	if !strings.EqualFold(userId, "user_uuid") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	serviceId := UserServiceServiceUuid.String()
	if !strings.EqualFold(serviceId, "service_uuid") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	createdAt := UserServiceCreatedAt.String()
	if !strings.EqualFold(createdAt, "created_at") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}

	updatedAt := UserServiceUpdatedAt.String()
	if !strings.EqualFold(updatedAt, "updated_at") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}
}
