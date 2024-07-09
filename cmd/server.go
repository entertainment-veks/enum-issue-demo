package main

import (
	"context"
	gen_proto "proto-bug-demo/pkg/proto/api"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	gen_proto.UnimplementedSwaggerServer
}

func (r *Server) GetLost(ctx context.Context, in *emptypb.Empty) (*gen_proto.MyEnum, error) {
	return &gen_proto.MyEnum{
		Value: gen_proto.MyEnum_LOST,
	}, nil
}

func (r *Server) GetOk(ctx context.Context, in *emptypb.Empty) (*gen_proto.MyEnum, error) {
	return &gen_proto.MyEnum{
		Value: gen_proto.MyEnum_OK,
	}, nil
}
