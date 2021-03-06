// @generated Code generated by thrift-gen. Do not modify.

// Package role is generated code used to make or handle TChannel calls using Thrift.
package role

import (
	"fmt"

	athrift "github.com/apache/thrift/lib/go/thrift"
	"github.com/uber/tchannel-go/thrift"
)

// Interfaces for the service and client for the services defined in the IDL.

// TChanRoleService is the interface that defines the server handler and client interface.
type TChanRoleService interface {
	GetMembers(ctx thrift.Context, role string) ([]string, error)
	SetRole(ctx thrift.Context, role string) error
}

// Implementation of a client and service handler.

type tchanRoleServiceClient struct {
	thriftService string
	client        thrift.TChanClient
}

func NewTChanRoleServiceInheritedClient(thriftService string, client thrift.TChanClient) *tchanRoleServiceClient {
	return &tchanRoleServiceClient{
		thriftService,
		client,
	}
}

// NewTChanRoleServiceClient creates a client that can be used to make remote calls.
func NewTChanRoleServiceClient(client thrift.TChanClient) TChanRoleService {
	return NewTChanRoleServiceInheritedClient("RoleService", client)
}

func (c *tchanRoleServiceClient) GetMembers(ctx thrift.Context, role string) ([]string, error) {
	var resp RoleServiceGetMembersResult
	args := RoleServiceGetMembersArgs{
		Role: role,
	}
	success, err := c.client.Call(ctx, c.thriftService, "GetMembers", &args, &resp)
	if err == nil && !success {
	}

	return resp.GetSuccess(), err
}

func (c *tchanRoleServiceClient) SetRole(ctx thrift.Context, role string) error {
	var resp RoleServiceSetRoleResult
	args := RoleServiceSetRoleArgs{
		Role: role,
	}
	success, err := c.client.Call(ctx, c.thriftService, "SetRole", &args, &resp)
	if err == nil && !success {
	}

	return err
}

type tchanRoleServiceServer struct {
	handler TChanRoleService
}

// NewTChanRoleServiceServer wraps a handler for TChanRoleService so it can be
// registered with a thrift.Server.
func NewTChanRoleServiceServer(handler TChanRoleService) thrift.TChanServer {
	return &tchanRoleServiceServer{
		handler,
	}
}

func (s *tchanRoleServiceServer) Service() string {
	return "RoleService"
}

func (s *tchanRoleServiceServer) Methods() []string {
	return []string{
		"GetMembers",
		"SetRole",
	}
}

func (s *tchanRoleServiceServer) Handle(ctx thrift.Context, methodName string, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	args, err := s.GetArgs(methodName, protocol)
	if err != nil {
		return false, nil, err
	}
	return s.HandleArgs(ctx, methodName, args)
}

func (s *tchanRoleServiceServer) GetArgs(methodName string, protocol athrift.TProtocol) (args interface{}, err error) {
	switch methodName {
	case "GetMembers":
		args, err = s.readGetMembers(protocol)
	case "SetRole":
		args, err = s.readSetRole(protocol)

	default:
		err = fmt.Errorf("method %v not found in service %v", methodName, s.Service())
	}
	return
}

func (s *tchanRoleServiceServer) HandleArgs(ctx thrift.Context, methodName string, args interface{}) (bool, athrift.TStruct, error) {
	switch methodName {
	case "GetMembers":
		return s.handleGetMembers(ctx, args.(RoleServiceGetMembersArgs))
	case "SetRole":
		return s.handleSetRole(ctx, args.(RoleServiceSetRoleArgs))

	default:
		return false, nil, fmt.Errorf("method %v not found in service %v", methodName, s.Service())
	}
}

func (s *tchanRoleServiceServer) readGetMembers(protocol athrift.TProtocol) (interface{}, error) {
	var req RoleServiceGetMembersArgs

	if err := req.Read(protocol); err != nil {
		return nil, err
	}
	return req, nil
}

func (s *tchanRoleServiceServer) handleGetMembers(ctx thrift.Context, req RoleServiceGetMembersArgs) (bool, athrift.TStruct, error) {
	var res RoleServiceGetMembersResult
	r, err :=
		s.handler.GetMembers(ctx, req.Role)

	if err != nil {
		return false, nil, err
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanRoleServiceServer) readSetRole(protocol athrift.TProtocol) (interface{}, error) {
	var req RoleServiceSetRoleArgs

	if err := req.Read(protocol); err != nil {
		return nil, err
	}
	return req, nil
}

func (s *tchanRoleServiceServer) handleSetRole(ctx thrift.Context, req RoleServiceSetRoleArgs) (bool, athrift.TStruct, error) {
	var res RoleServiceSetRoleResult
	err :=
		s.handler.SetRole(ctx, req.Role)

	if err != nil {
		return false, nil, err
	} else {
	}

	return err == nil, &res, nil
}
