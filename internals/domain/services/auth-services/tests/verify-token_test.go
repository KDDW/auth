package auth_services_test

import (
	test_utils "auth-service/internals/test-utils"
	"auth-service/internals/utils/tokens"
	"testing"
)

func TestInvalidVerifyToken(t *testing.T) {

	services := test_utils.SetupServices()

	result, _ := services.AuthServices.VerifyToken("")

	if result != false {
		t.Errorf("Expected false, got %v", result)
	}

}

func TestValidVerifyToken(t *testing.T) {

	services := test_utils.SetupServices()
	user := test_utils.CreateTestUser(services)

	token, _ := tokens.GenerateAccessToken(user.ID)

	result, _ := services.AuthServices.VerifyToken(token)

	if result != true {
		t.Errorf("Expected true, got %v", result)
	}
}

func TestValidVerifyTokenWithoutUser(t *testing.T) {

	services := test_utils.SetupServices()

	token, _ := tokens.GenerateAccessToken(10)

	result, _ := services.AuthServices.VerifyToken(token)

	if result != false {
		t.Errorf("Expected false, got %v", result)
	}
}
