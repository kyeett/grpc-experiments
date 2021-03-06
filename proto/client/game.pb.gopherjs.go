// Code generated by protoc-gen-gopherjs. DO NOT EDIT.
// source: proto/game.proto

/*
	Package client is a generated protocol buffer package.

	Web exposes a backend server over gRPC.

	It is generated from these files:
		proto/game.proto

	It has these top-level messages:
		Empty
		PlayerID
		EntityResponse
*/
package client

import jspb "github.com/johanbrandhorst/protobuf/jspb"
import google_protobuf "github.com/johanbrandhorst/protobuf/ptypes/empty"

import (
	context "context"

	grpcweb "github.com/johanbrandhorst/protobuf/grpcweb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the jspb package it is being compiled against.
const _ = jspb.JspbPackageIsVersion2

type Empty struct {
}

// MarshalToWriter marshals Empty to the provided writer.
func (m *Empty) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	return
}

// Marshal marshals Empty to a slice of bytes.
func (m *Empty) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a Empty from the provided reader.
func (m *Empty) UnmarshalFromReader(reader jspb.Reader) *Empty {
	for reader.Next() {
		if m == nil {
			m = &Empty{}
		}

		switch reader.GetFieldNumber() {
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a Empty from a slice of bytes.
func (m *Empty) Unmarshal(rawBytes []byte) (*Empty, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

type PlayerID struct {
	ID string
}

// GetID gets the ID of the PlayerID.
func (m *PlayerID) GetID() (x string) {
	if m == nil {
		return x
	}
	return m.ID
}

// MarshalToWriter marshals PlayerID to the provided writer.
func (m *PlayerID) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if len(m.ID) > 0 {
		writer.WriteString(1, m.ID)
	}

	return
}

// Marshal marshals PlayerID to a slice of bytes.
func (m *PlayerID) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a PlayerID from the provided reader.
func (m *PlayerID) UnmarshalFromReader(reader jspb.Reader) *PlayerID {
	for reader.Next() {
		if m == nil {
			m = &PlayerID{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			m.ID = reader.ReadString()
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a PlayerID from a slice of bytes.
func (m *PlayerID) Unmarshal(rawBytes []byte) (*PlayerID, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

type EntityResponse struct {
	Payload []byte
}

// GetPayload gets the Payload of the EntityResponse.
func (m *EntityResponse) GetPayload() (x []byte) {
	if m == nil {
		return x
	}
	return m.Payload
}

// MarshalToWriter marshals EntityResponse to the provided writer.
func (m *EntityResponse) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if len(m.Payload) > 0 {
		writer.WriteBytes(1, m.Payload)
	}

	return
}

// Marshal marshals EntityResponse to a slice of bytes.
func (m *EntityResponse) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a EntityResponse from the provided reader.
func (m *EntityResponse) UnmarshalFromReader(reader jspb.Reader) *EntityResponse {
	for reader.Next() {
		if m == nil {
			m = &EntityResponse{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			m.Payload = reader.ReadBytes()
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a EntityResponse from a slice of bytes.
func (m *EntityResponse) Unmarshal(rawBytes []byte) (*EntityResponse, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpcweb.Client

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpcweb package it is being compiled against.
const _ = grpcweb.GrpcWebPackageIsVersion3

// Client API for Backend service

type BackendClient interface {
	NewPlayer(ctx context.Context, in *google_protobuf.Empty, opts ...grpcweb.CallOption) (*PlayerID, error)
	// rpc PerformAction(ActionRequest) returns (ActionRequest) {}
	EntityStream(ctx context.Context, in *google_protobuf.Empty, opts ...grpcweb.CallOption) (Backend_EntityStreamClient, error)
}

type backendClient struct {
	client *grpcweb.Client
}

// NewBackendClient creates a new gRPC-Web client.
func NewBackendClient(hostname string, opts ...grpcweb.DialOption) BackendClient {
	return &backendClient{
		client: grpcweb.NewClient(hostname, "game.Backend", opts...),
	}
}

func (c *backendClient) NewPlayer(ctx context.Context, in *google_protobuf.Empty, opts ...grpcweb.CallOption) (*PlayerID, error) {
	resp, err := c.client.RPCCall(ctx, "NewPlayer", in.Marshal(), opts...)
	if err != nil {
		return nil, err
	}

	return new(PlayerID).Unmarshal(resp)
}

func (c *backendClient) EntityStream(ctx context.Context, in *google_protobuf.Empty, opts ...grpcweb.CallOption) (Backend_EntityStreamClient, error) {
	srv, err := c.client.NewClientStream(ctx, false, true, "EntityStream", opts...)
	if err != nil {
		return nil, err
	}

	err = srv.SendMsg(in.Marshal())
	if err != nil {
		return nil, err
	}

	return &backendEntityStreamClient{srv}, nil
}

type Backend_EntityStreamClient interface {
	Recv() (*EntityResponse, error)
	grpcweb.ClientStream
}

type backendEntityStreamClient struct {
	grpcweb.ClientStream
}

func (x *backendEntityStreamClient) Recv() (*EntityResponse, error) {
	resp, err := x.RecvMsg()
	if err != nil {
		return nil, err
	}

	return new(EntityResponse).Unmarshal(resp)
}
