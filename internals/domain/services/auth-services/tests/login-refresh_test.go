package auth_services_test

import (
	test_utils "auth-service/internals/test-utils"
	"testing"
)

func TestLoginRefreshToken(t *testing.T) {

	services := test_utils.SetupServices()
	user := createTestUser(services)

	result, err := services.AuthServices.Login(user.Email, defaultUser.Password, user.Realm.Code)

	if result == nil || err != nil {
		t.Errorf("Expected success, got %v", err)
	}

	result, err = services.AuthServices.LoginRefreshToken(result.RefreshToken)

	if result == nil || err != nil {
		t.Errorf("Expected success, got %v", err)
	}
}
