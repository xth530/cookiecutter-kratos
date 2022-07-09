package service

import "github.com/google/wire"

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(New{{cookiecutter.repo_service_name}}Service)
