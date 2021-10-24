package omp_grpc_template

import (
	desc "github.com/ozonmp/omp-grpc-template/pkg/omp-grpc-template"
)

type Implementation struct {
	desc.UnimplementedOmpGrpcTemplateServiceServer
}

func NewTemplateAPI() desc.OmpGrpcTemplateServiceServer {
	return &Implementation{}
}