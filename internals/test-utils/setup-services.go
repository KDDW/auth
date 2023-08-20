package test_utils

import (
	"auth-service/internals/domain/services"
	"auth-service/internals/infra/repositories"
)

func SetupServices() *services.Services {
	repositories := repositories.GetRepositories(nil)
	return services.GetServices(repositories)
}
