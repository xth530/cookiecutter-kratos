package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(New{{cookiecutter.repo_service_name}}Usecase)
