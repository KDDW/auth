package auth_services_test

import (
	test_utils "auth-service/internals/test-utils"
	"testing"
)

func TestLoginRefreshToken(t *testing.T) {

	services := test_utils.SetupServices()
	user := test_utils.CreateTestUser(services)
	defaultUser := test_utils.DefaultUser

	result, err := services.AuthServices.Login(user.Email, defaultUser.Password, user.Realm.Code)

	if err != nil {
		t.Errorf("Expected success, got %v", err)
	}

	if result == nil {
		t.Errorf("Expected result != nil, got %v", result)
	}

	result, err = services.AuthServices.LoginRefreshToken(result.RefreshToken)

	if err != nil {
		t.Errorf("Expected success, got %v", err)
	}

	if result == nil {
		t.Errorf("Expected result != nil, got %v", result)
	}
}
