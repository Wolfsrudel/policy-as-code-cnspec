// Code generated by protoc-gen-rangerrpc version DO NOT EDIT.
// source: cnspec_upstream.proto

package upstream

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"go.mondoo.com/cnspec/policy"
	ranger "go.mondoo.com/ranger-rpc"
	"go.mondoo.com/ranger-rpc/metadata"
	jsonpb "google.golang.org/protobuf/encoding/protojson"
	pb "google.golang.org/protobuf/proto"
)

// service interface definition

type Reporting interface {
	StoreReport(context.Context, *policy.ReportCollection) (*ReportID, error)
}

// client implementation

type ReportingClient struct {
	ranger.Client
	httpclient ranger.HTTPClient
	prefix     string
}

func NewReportingClient(addr string, client ranger.HTTPClient, plugins ...ranger.ClientPlugin) (*ReportingClient, error) {
	base, err := url.Parse(ranger.SanitizeUrl(addr))
	if err != nil {
		return nil, err
	}

	u, err := url.Parse("./Reporting")
	if err != nil {
		return nil, err
	}

	serviceClient := &ReportingClient{
		httpclient: client,
		prefix:     base.ResolveReference(u).String(),
	}
	serviceClient.AddPlugins(plugins...)
	return serviceClient, nil
}
func (c *ReportingClient) StoreReport(ctx context.Context, in *policy.ReportCollection) (*ReportID, error) {
	out := new(ReportID)
	err := c.DoClientRequest(ctx, c.httpclient, strings.Join([]string{c.prefix, "/StoreReport"}, ""), in, out)
	return out, err
}

// server implementation

type ReportingServerOption func(s *ReportingServer)

func WithUnknownFieldsForReportingServer() ReportingServerOption {
	return func(s *ReportingServer) {
		s.allowUnknownFields = true
	}
}

func NewReportingServer(handler Reporting, opts ...ReportingServerOption) http.Handler {
	srv := &ReportingServer{
		handler: handler,
	}

	for i := range opts {
		opts[i](srv)
	}

	service := ranger.Service{
		Name: "Reporting",
		Methods: map[string]ranger.Method{
			"StoreReport": srv.StoreReport,
		},
	}
	return ranger.NewRPCServer(&service)
}

type ReportingServer struct {
	handler            Reporting
	allowUnknownFields bool
}

func (p *ReportingServer) StoreReport(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
	var req policy.ReportCollection
	var err error

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("could not access header")
	}

	switch md.First("Content-Type") {
	case "application/protobuf", "application/octet-stream", "application/grpc+proto":
		err = pb.Unmarshal(*reqBytes, &req)
	default:
		// handle case of empty object
		if len(*reqBytes) > 0 {
			err = jsonpb.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(*reqBytes, &req)
		}
	}

	if err != nil {
		return nil, err
	}
	return p.handler.StoreReport(ctx, &req)
}
