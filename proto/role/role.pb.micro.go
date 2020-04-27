// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/role/role.proto

package go_micro_srv_basic_role

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"

	context "context"

	client "github.com/micro/go-micro/v2/client"

	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Role service

type RoleService interface {
	Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error)
	AddOrUpdate(ctx context.Context, in *RoleAddOrUpdateRequest, opts ...client.CallOption) (*Response, error)
}

type roleService struct {
	c    client.Client
	name string
}

func NewRoleService(name string, c client.Client) RoleService {
	return &roleService{
		c:    c,
		name: name,
	}
}

func (c *roleService) Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error) {
	req := c.c.NewRequest(c.name, "Role.Get", in)
	out := new(GetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleService) AddOrUpdate(ctx context.Context, in *RoleAddOrUpdateRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Role.AddOrUpdate", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Role service

type RoleHandler interface {
	Get(context.Context, *GetRequest, *GetResponse) error
	AddOrUpdate(context.Context, *RoleAddOrUpdateRequest, *Response) error
}

func RegisterRoleHandler(s server.Server, hdlr RoleHandler, opts ...server.HandlerOption) error {
	type role interface {
		Get(ctx context.Context, in *GetRequest, out *GetResponse) error
		AddOrUpdate(ctx context.Context, in *RoleAddOrUpdateRequest, out *Response) error
	}
	type Role struct {
		role
	}
	h := &roleHandler{hdlr}
	return s.Handle(s.NewHandler(&Role{h}, opts...))
}

type roleHandler struct {
	RoleHandler
}

func (h *roleHandler) Get(ctx context.Context, in *GetRequest, out *GetResponse) error {
	return h.RoleHandler.Get(ctx, in, out)
}

func (h *roleHandler) AddOrUpdate(ctx context.Context, in *RoleAddOrUpdateRequest, out *Response) error {
	return h.RoleHandler.AddOrUpdate(ctx, in, out)
}
