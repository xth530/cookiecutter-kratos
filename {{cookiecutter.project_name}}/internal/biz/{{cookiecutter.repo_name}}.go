package biz

import (
	"context"

	v1 "{{cookiecutter.module_name}}/api/{{cookiecutter.repo_name}}/v1"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// {{cookiecutter.repo_service_name}} is a {{cookiecutter.repo_service_name}} model.
type {{cookiecutter.repo_service_name}} struct {
	Hello string
}

// {{cookiecutter.repo_service_name}}Repo is a Greater repo.
type {{cookiecutter.repo_service_name}}Repo interface {
	Save(context.Context, *{{cookiecutter.repo_service_name}}) (*{{cookiecutter.repo_service_name}}, error)
	Update(context.Context, *{{cookiecutter.repo_service_name}}) (*{{cookiecutter.repo_service_name}}, error)
	FindByID(context.Context, int64) (*{{cookiecutter.repo_service_name}}, error)
	ListByHello(context.Context, string) ([]*{{cookiecutter.repo_service_name}}, error)
	ListAll(context.Context) ([]*{{cookiecutter.repo_service_name}}, error)
}

// {{cookiecutter.repo_service_name}}Usecase is a {{cookiecutter.repo_service_name}} usecase.
type {{cookiecutter.repo_service_name}}Usecase struct {
	repo {{cookiecutter.repo_service_name}}Repo
	log  *log.Helper
}

// New{{cookiecutter.repo_service_name}}Usecase new a {{cookiecutter.repo_service_name}} usecase.
func New{{cookiecutter.repo_service_name}}Usecase(repo {{cookiecutter.repo_service_name}}Repo, logger log.Logger) *{{cookiecutter.repo_service_name}}Usecase {
	return &{{cookiecutter.repo_service_name}}Usecase{repo: repo, log: log.NewHelper(logger)}
}

// Create{{cookiecutter.repo_service_name}} creates a {{cookiecutter.repo_service_name}}, and returns the new {{cookiecutter.repo_service_name}}.
func (uc *{{cookiecutter.repo_service_name}}Usecase) Create{{cookiecutter.repo_service_name}}(ctx context.Context, g *{{cookiecutter.repo_service_name}}) (*{{cookiecutter.repo_service_name}}, error) {
	uc.log.WithContext(ctx).Infof("Create{{cookiecutter.repo_service_name}}: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
