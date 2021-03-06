// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: github.com/pijalu/go.hands.two/frinsultproto/frinsult.proto

/*
Package frinsultproto is a generated protocol buffer package.

It is generated from these files:
	github.com/pijalu/go.hands.two/frinsultproto/frinsult.proto

It has these top-level messages:
	Frinsult
	Frinsults
	ByIDRequest
	VoteRequest
	Void
*/
package frinsultproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for FrinsultService service

type FrinsultService interface {
	GetFrinsultByID(ctx context.Context, in *ByIDRequest, opts ...client.CallOption) (*Frinsult, error)
	DeleteFrinsultByID(ctx context.Context, in *ByIDRequest, opts ...client.CallOption) (*Void, error)
	InsertFrinsult(ctx context.Context, in *Frinsult, opts ...client.CallOption) (*Frinsult, error)
	UpdateFrinsult(ctx context.Context, in *Frinsult, opts ...client.CallOption) (*Void, error)
	VoteFrinsultByID(ctx context.Context, in *VoteRequest, opts ...client.CallOption) (*Void, error)
	GetFrinsults(ctx context.Context, in *Void, opts ...client.CallOption) (*Frinsults, error)
}

type frinsultService struct {
	c    client.Client
	name string
}

func NewFrinsultService(name string, c client.Client) FrinsultService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "frinsultproto"
	}
	return &frinsultService{
		c:    c,
		name: name,
	}
}

func (c *frinsultService) GetFrinsultByID(ctx context.Context, in *ByIDRequest, opts ...client.CallOption) (*Frinsult, error) {
	req := c.c.NewRequest(c.name, "FrinsultService.GetFrinsultByID", in)
	out := new(Frinsult)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frinsultService) DeleteFrinsultByID(ctx context.Context, in *ByIDRequest, opts ...client.CallOption) (*Void, error) {
	req := c.c.NewRequest(c.name, "FrinsultService.DeleteFrinsultByID", in)
	out := new(Void)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frinsultService) InsertFrinsult(ctx context.Context, in *Frinsult, opts ...client.CallOption) (*Frinsult, error) {
	req := c.c.NewRequest(c.name, "FrinsultService.InsertFrinsult", in)
	out := new(Frinsult)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frinsultService) UpdateFrinsult(ctx context.Context, in *Frinsult, opts ...client.CallOption) (*Void, error) {
	req := c.c.NewRequest(c.name, "FrinsultService.UpdateFrinsult", in)
	out := new(Void)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frinsultService) VoteFrinsultByID(ctx context.Context, in *VoteRequest, opts ...client.CallOption) (*Void, error) {
	req := c.c.NewRequest(c.name, "FrinsultService.VoteFrinsultByID", in)
	out := new(Void)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frinsultService) GetFrinsults(ctx context.Context, in *Void, opts ...client.CallOption) (*Frinsults, error) {
	req := c.c.NewRequest(c.name, "FrinsultService.GetFrinsults", in)
	out := new(Frinsults)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FrinsultService service

type FrinsultServiceHandler interface {
	GetFrinsultByID(context.Context, *ByIDRequest, *Frinsult) error
	DeleteFrinsultByID(context.Context, *ByIDRequest, *Void) error
	InsertFrinsult(context.Context, *Frinsult, *Frinsult) error
	UpdateFrinsult(context.Context, *Frinsult, *Void) error
	VoteFrinsultByID(context.Context, *VoteRequest, *Void) error
	GetFrinsults(context.Context, *Void, *Frinsults) error
}

func RegisterFrinsultServiceHandler(s server.Server, hdlr FrinsultServiceHandler, opts ...server.HandlerOption) {
	type frinsultService interface {
		GetFrinsultByID(ctx context.Context, in *ByIDRequest, out *Frinsult) error
		DeleteFrinsultByID(ctx context.Context, in *ByIDRequest, out *Void) error
		InsertFrinsult(ctx context.Context, in *Frinsult, out *Frinsult) error
		UpdateFrinsult(ctx context.Context, in *Frinsult, out *Void) error
		VoteFrinsultByID(ctx context.Context, in *VoteRequest, out *Void) error
		GetFrinsults(ctx context.Context, in *Void, out *Frinsults) error
	}
	type FrinsultService struct {
		frinsultService
	}
	h := &frinsultServiceHandler{hdlr}
	s.Handle(s.NewHandler(&FrinsultService{h}, opts...))
}

type frinsultServiceHandler struct {
	FrinsultServiceHandler
}

func (h *frinsultServiceHandler) GetFrinsultByID(ctx context.Context, in *ByIDRequest, out *Frinsult) error {
	return h.FrinsultServiceHandler.GetFrinsultByID(ctx, in, out)
}

func (h *frinsultServiceHandler) DeleteFrinsultByID(ctx context.Context, in *ByIDRequest, out *Void) error {
	return h.FrinsultServiceHandler.DeleteFrinsultByID(ctx, in, out)
}

func (h *frinsultServiceHandler) InsertFrinsult(ctx context.Context, in *Frinsult, out *Frinsult) error {
	return h.FrinsultServiceHandler.InsertFrinsult(ctx, in, out)
}

func (h *frinsultServiceHandler) UpdateFrinsult(ctx context.Context, in *Frinsult, out *Void) error {
	return h.FrinsultServiceHandler.UpdateFrinsult(ctx, in, out)
}

func (h *frinsultServiceHandler) VoteFrinsultByID(ctx context.Context, in *VoteRequest, out *Void) error {
	return h.FrinsultServiceHandler.VoteFrinsultByID(ctx, in, out)
}

func (h *frinsultServiceHandler) GetFrinsults(ctx context.Context, in *Void, out *Frinsults) error {
	return h.FrinsultServiceHandler.GetFrinsults(ctx, in, out)
}
