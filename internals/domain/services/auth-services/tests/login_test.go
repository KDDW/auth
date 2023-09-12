package auth_services_test

import (
	test_utils "auth-service/internals/test-utils"
	"testing"
)

func TestLoginWithRightCredentials(t *testing.T) {

	services := test_utils.SetupServices()
	user := test_utils.CreateTestUser(services)
	defaultUser := test_utils.DefaultUser

	result, err := services.AuthServices.Login(user.Email, defaultUser.Password, user.Realm.Code)

	if result == nil || err != nil {
		t.Errorf("Expected success, got %v", err)
	}
}

func TestLoginWithWrongCredentials(t *testing.T) {

	services := test_utils.SetupServices()
	user := test_utils.CreateTestUser(services)
	defaultUser := test_utils.DefaultUser

	result, err := services.AuthServices.Login(user.Email, "wrong-pwd", user.Realm.Code)

	if result != nil || err == nil {
		t.Errorf("Expected error, got %v", result)
	}

	result, err = services.AuthServices.Login("wrong-mail", defaultUser.Password, user.Realm.Code)

	if result != nil || err == nil {
		t.Errorf("Expected error, got %v", result)
	}

	result, err = services.AuthServices.Login(defaultUser.Email, defaultUser.Password, "wrong realm")

	if result != nil || err == nil {
		t.Errorf("Expected error, got %v", result)
	}
}
