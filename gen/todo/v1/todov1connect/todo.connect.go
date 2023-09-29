// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: todo/v1/todo.proto

package todov1connect

import (
	v1 "connectgo/gen/todo/v1"
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion0_1_0

const (
	// ToDoServiceName is the fully-qualified name of the ToDoService service.
	ToDoServiceName = "todo.v1.ToDoService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ToDoServiceCreateToDoProcedure is the fully-qualified name of the ToDoService's CreateToDo RPC.
	ToDoServiceCreateToDoProcedure = "/todo.v1.ToDoService/CreateToDo"
)

// ToDoServiceClient is a client for the todo.v1.ToDoService service.
type ToDoServiceClient interface {
	CreateToDo(context.Context, *connect.Request[v1.CreateToDoRequest]) (*connect.Response[v1.CreateToDoResponse], error)
}

// NewToDoServiceClient constructs a client for the todo.v1.ToDoService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewToDoServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ToDoServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &toDoServiceClient{
		createToDo: connect.NewClient[v1.CreateToDoRequest, v1.CreateToDoResponse](
			httpClient,
			baseURL+ToDoServiceCreateToDoProcedure,
			opts...,
		),
	}
}

// toDoServiceClient implements ToDoServiceClient.
type toDoServiceClient struct {
	createToDo *connect.Client[v1.CreateToDoRequest, v1.CreateToDoResponse]
}

// CreateToDo calls todo.v1.ToDoService.CreateToDo.
func (c *toDoServiceClient) CreateToDo(ctx context.Context, req *connect.Request[v1.CreateToDoRequest]) (*connect.Response[v1.CreateToDoResponse], error) {
	return c.createToDo.CallUnary(ctx, req)
}

// ToDoServiceHandler is an implementation of the todo.v1.ToDoService service.
type ToDoServiceHandler interface {
	CreateToDo(context.Context, *connect.Request[v1.CreateToDoRequest]) (*connect.Response[v1.CreateToDoResponse], error)
}

// NewToDoServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewToDoServiceHandler(svc ToDoServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	toDoServiceCreateToDoHandler := connect.NewUnaryHandler(
		ToDoServiceCreateToDoProcedure,
		svc.CreateToDo,
		opts...,
	)
	return "/todo.v1.ToDoService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ToDoServiceCreateToDoProcedure:
			toDoServiceCreateToDoHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedToDoServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedToDoServiceHandler struct{}

func (UnimplementedToDoServiceHandler) CreateToDo(context.Context, *connect.Request[v1.CreateToDoRequest]) (*connect.Response[v1.CreateToDoResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("todo.v1.ToDoService.CreateToDo is not implemented"))
}
