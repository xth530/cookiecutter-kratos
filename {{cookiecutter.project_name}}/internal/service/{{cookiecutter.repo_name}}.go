package service

import (
	"context"

	v1 "{{cookiecutter.module_name}}/api/{{cookiecutter.repo_name}}/v1"
	"{{cookiecutter.module_name}}/internal/biz"
)

// {{cookiecutter.repo_service_name}}Service is a {{cookiecutter.repo_name}} service.
type {{cookiecutter.repo_service_name}}Service struct {
	v1.Unimplemented{{cookiecutter.repo_service_name}}Server

	uc *biz.{{cookiecutter.repo_service_name}}Usecase
}

// New{{cookiecutter.repo_service_name}}Service new a {{cookiecutter.repo_name}} service.
func New{{cookiecutter.repo_service_name}}Service(uc *biz.{{cookiecutter.repo_service_name}}Usecase) *{{cookiecutter.repo_service_name}}Service {
	return &{{cookiecutter.repo_service_name}}Service{uc: uc}
}

// SayHello implements {{cookiecutter.repo_name}}.{{cookiecutter.repo_service_name}}Server.
func (s *{{cookiecutter.repo_service_name}}Service) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.Create{{cookiecutter.repo_service_name}}(ctx, &biz.{{cookiecutter.repo_service_name}}{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}
