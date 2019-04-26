// Code generated by sdkgen. DO NOT EDIT.

//nolint
package endpoint

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/endpoint"
)

//revive:disable

// ApiEndpointServiceClient is a endpoint.ApiEndpointServiceClient with
// lazy GRPC connection initialization.
type ApiEndpointServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

var _ endpoint.ApiEndpointServiceClient = &ApiEndpointServiceClient{}

// Get implements endpoint.ApiEndpointServiceClient
func (c *ApiEndpointServiceClient) Get(ctx context.Context, in *endpoint.GetApiEndpointRequest, opts ...grpc.CallOption) (*endpoint.ApiEndpoint, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return endpoint.NewApiEndpointServiceClient(conn).Get(ctx, in, opts...)
}

// List implements endpoint.ApiEndpointServiceClient
func (c *ApiEndpointServiceClient) List(ctx context.Context, in *endpoint.ListApiEndpointsRequest, opts ...grpc.CallOption) (*endpoint.ListApiEndpointsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return endpoint.NewApiEndpointServiceClient(conn).List(ctx, in, opts...)
}