// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: accesskey.proto

package client

import (
	fmt "fmt"
	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/core/services/accesskey/pb"
	grpc1 "google.golang.org/grpc"
	reflect "reflect"
	strings "strings"
)

var dependencies = []string{
	"grpc-client@erda.core.services.accesskey",
	"grpc-client",
}

// +provider
type provider struct {
	client Client
}

func (p *provider) Init(ctx servicehub.Context) error {
	var conn grpc.ClientConnInterface
	for _, dep := range dependencies {
		c, ok := ctx.Service(dep).(grpc.ClientConnInterface)
		if ok {
			conn = c
			break
		}
	}
	if conn == nil {
		return fmt.Errorf("not found connector in (%s)", strings.Join(dependencies, ", "))
	}
	p.client = New(conn)
	return nil
}

var (
	clientsType                = reflect.TypeOf((*Client)(nil)).Elem()
	accessKeyServiceClientType = reflect.TypeOf((*pb.AccessKeyServiceClient)(nil)).Elem()
	accessKeyServiceServerType = reflect.TypeOf((*pb.AccessKeyServiceServer)(nil)).Elem()
)

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	var opts []grpc1.CallOption
	for _, arg := range args {
		if opt, ok := arg.(grpc1.CallOption); ok {
			opts = append(opts, opt)
		}
	}
	switch ctx.Service() {
	case "erda.core.services.accesskey-client":
		return p.client
	case "erda.core.services.accesskey.AccessKeyService":
		return &accessKeyServiceWrapper{client: p.client.AccessKeyService(), opts: opts}
	case "erda.core.services.accesskey.AccessKeyService.client":
		return p.client.AccessKeyService()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case accessKeyServiceClientType:
		return p.client.AccessKeyService()
	case accessKeyServiceServerType:
		return &accessKeyServiceWrapper{client: p.client.AccessKeyService(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.core.services.accesskey-client", &servicehub.Spec{
		Services: []string{
			"erda.core.services.accesskey.AccessKeyService",
			"erda.core.services.accesskey-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			accessKeyServiceClientType,
			// server types
			accessKeyServiceServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
