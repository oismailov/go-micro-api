// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/api.proto

package api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
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

// Client API for Dtu service

type DtuService interface {
	AccountLookup(ctx context.Context, in *AccountRequest, opts ...client.CallOption) (*AccountResponse, error)
	GetProducts(ctx context.Context, in *AccountResponse, opts ...client.CallOption) (*ProductsResponse, error)
	SendTransfer(ctx context.Context, in *TransferRequest, opts ...client.CallOption) (*TransferResponse, error)
}

type dtuService struct {
	c    client.Client
	name string
}

func NewDtuService(name string, c client.Client) DtuService {
	return &dtuService{
		c:    c,
		name: name,
	}
}

func (c *dtuService) AccountLookup(ctx context.Context, in *AccountRequest, opts ...client.CallOption) (*AccountResponse, error) {
	req := c.c.NewRequest(c.name, "Dtu.AccountLookup", in)
	out := new(AccountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dtuService) GetProducts(ctx context.Context, in *AccountResponse, opts ...client.CallOption) (*ProductsResponse, error) {
	req := c.c.NewRequest(c.name, "Dtu.GetProducts", in)
	out := new(ProductsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dtuService) SendTransfer(ctx context.Context, in *TransferRequest, opts ...client.CallOption) (*TransferResponse, error) {
	req := c.c.NewRequest(c.name, "Dtu.SendTransfer", in)
	out := new(TransferResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Dtu service

type DtuHandler interface {
	AccountLookup(context.Context, *AccountRequest, *AccountResponse) error
	GetProducts(context.Context, *AccountResponse, *ProductsResponse) error
	SendTransfer(context.Context, *TransferRequest, *TransferResponse) error
}

func RegisterDtuHandler(s server.Server, hdlr DtuHandler, opts ...server.HandlerOption) error {
	type dtu interface {
		AccountLookup(ctx context.Context, in *AccountRequest, out *AccountResponse) error
		GetProducts(ctx context.Context, in *AccountResponse, out *ProductsResponse) error
		SendTransfer(ctx context.Context, in *TransferRequest, out *TransferResponse) error
	}
	type Dtu struct {
		dtu
	}
	h := &dtuHandler{hdlr}
	return s.Handle(s.NewHandler(&Dtu{h}, opts...))
}

type dtuHandler struct {
	DtuHandler
}

func (h *dtuHandler) AccountLookup(ctx context.Context, in *AccountRequest, out *AccountResponse) error {
	return h.DtuHandler.AccountLookup(ctx, in, out)
}

func (h *dtuHandler) GetProducts(ctx context.Context, in *AccountResponse, out *ProductsResponse) error {
	return h.DtuHandler.GetProducts(ctx, in, out)
}

func (h *dtuHandler) SendTransfer(ctx context.Context, in *TransferRequest, out *TransferResponse) error {
	return h.DtuHandler.SendTransfer(ctx, in, out)
}
