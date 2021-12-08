// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package aiplatform

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	aiplatformpb "google.golang.org/genproto/googleapis/cloud/aiplatform/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newIndexClientHook clientHook

// IndexCallOptions contains the retry settings for each method of IndexClient.
type IndexCallOptions struct {
	CreateIndex []gax.CallOption
	GetIndex    []gax.CallOption
	ListIndexes []gax.CallOption
	UpdateIndex []gax.CallOption
	DeleteIndex []gax.CallOption
}

func defaultIndexGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("aiplatform.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("aiplatform.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://aiplatform.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultIndexCallOptions() *IndexCallOptions {
	return &IndexCallOptions{
		CreateIndex: []gax.CallOption{},
		GetIndex:    []gax.CallOption{},
		ListIndexes: []gax.CallOption{},
		UpdateIndex: []gax.CallOption{},
		DeleteIndex: []gax.CallOption{},
	}
}

// internalIndexClient is an interface that defines the methods availaible from Vertex AI API.
type internalIndexClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateIndex(context.Context, *aiplatformpb.CreateIndexRequest, ...gax.CallOption) (*CreateIndexOperation, error)
	CreateIndexOperation(name string) *CreateIndexOperation
	GetIndex(context.Context, *aiplatformpb.GetIndexRequest, ...gax.CallOption) (*aiplatformpb.Index, error)
	ListIndexes(context.Context, *aiplatformpb.ListIndexesRequest, ...gax.CallOption) *IndexIterator
	UpdateIndex(context.Context, *aiplatformpb.UpdateIndexRequest, ...gax.CallOption) (*UpdateIndexOperation, error)
	UpdateIndexOperation(name string) *UpdateIndexOperation
	DeleteIndex(context.Context, *aiplatformpb.DeleteIndexRequest, ...gax.CallOption) (*DeleteIndexOperation, error)
	DeleteIndexOperation(name string) *DeleteIndexOperation
}

// IndexClient is a client for interacting with Vertex AI API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// A service for creating and managing Vertex AI’s Index resources.
type IndexClient struct {
	// The internal transport-dependent client.
	internalClient internalIndexClient

	// The call options for this service.
	CallOptions *IndexCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *IndexClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *IndexClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *IndexClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateIndex creates an Index.
func (c *IndexClient) CreateIndex(ctx context.Context, req *aiplatformpb.CreateIndexRequest, opts ...gax.CallOption) (*CreateIndexOperation, error) {
	return c.internalClient.CreateIndex(ctx, req, opts...)
}

// CreateIndexOperation returns a new CreateIndexOperation from a given name.
// The name must be that of a previously created CreateIndexOperation, possibly from a different process.
func (c *IndexClient) CreateIndexOperation(name string) *CreateIndexOperation {
	return c.internalClient.CreateIndexOperation(name)
}

// GetIndex gets an Index.
func (c *IndexClient) GetIndex(ctx context.Context, req *aiplatformpb.GetIndexRequest, opts ...gax.CallOption) (*aiplatformpb.Index, error) {
	return c.internalClient.GetIndex(ctx, req, opts...)
}

// ListIndexes lists Indexes in a Location.
func (c *IndexClient) ListIndexes(ctx context.Context, req *aiplatformpb.ListIndexesRequest, opts ...gax.CallOption) *IndexIterator {
	return c.internalClient.ListIndexes(ctx, req, opts...)
}

// UpdateIndex updates an Index.
func (c *IndexClient) UpdateIndex(ctx context.Context, req *aiplatformpb.UpdateIndexRequest, opts ...gax.CallOption) (*UpdateIndexOperation, error) {
	return c.internalClient.UpdateIndex(ctx, req, opts...)
}

// UpdateIndexOperation returns a new UpdateIndexOperation from a given name.
// The name must be that of a previously created UpdateIndexOperation, possibly from a different process.
func (c *IndexClient) UpdateIndexOperation(name string) *UpdateIndexOperation {
	return c.internalClient.UpdateIndexOperation(name)
}

// DeleteIndex deletes an Index.
// An Index can only be deleted when all its
// DeployedIndexes had been undeployed.
func (c *IndexClient) DeleteIndex(ctx context.Context, req *aiplatformpb.DeleteIndexRequest, opts ...gax.CallOption) (*DeleteIndexOperation, error) {
	return c.internalClient.DeleteIndex(ctx, req, opts...)
}

// DeleteIndexOperation returns a new DeleteIndexOperation from a given name.
// The name must be that of a previously created DeleteIndexOperation, possibly from a different process.
func (c *IndexClient) DeleteIndexOperation(name string) *DeleteIndexOperation {
	return c.internalClient.DeleteIndexOperation(name)
}

// indexGRPCClient is a client for interacting with Vertex AI API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type indexGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing IndexClient
	CallOptions **IndexCallOptions

	// The gRPC API client.
	indexClient aiplatformpb.IndexServiceClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewIndexClient creates a new index service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// A service for creating and managing Vertex AI’s Index resources.
func NewIndexClient(ctx context.Context, opts ...option.ClientOption) (*IndexClient, error) {
	clientOpts := defaultIndexGRPCClientOptions()
	if newIndexClientHook != nil {
		hookOpts, err := newIndexClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := IndexClient{CallOptions: defaultIndexCallOptions()}

	c := &indexGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		indexClient:      aiplatformpb.NewIndexServiceClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	client.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	c.LROClient = &client.LROClient
	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *indexGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *indexGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *indexGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *indexGRPCClient) CreateIndex(ctx context.Context, req *aiplatformpb.CreateIndexRequest, opts ...gax.CallOption) (*CreateIndexOperation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateIndex[0:len((*c.CallOptions).CreateIndex):len((*c.CallOptions).CreateIndex)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.indexClient.CreateIndex(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &CreateIndexOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *indexGRPCClient) GetIndex(ctx context.Context, req *aiplatformpb.GetIndexRequest, opts ...gax.CallOption) (*aiplatformpb.Index, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetIndex[0:len((*c.CallOptions).GetIndex):len((*c.CallOptions).GetIndex)], opts...)
	var resp *aiplatformpb.Index
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.indexClient.GetIndex(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *indexGRPCClient) ListIndexes(ctx context.Context, req *aiplatformpb.ListIndexesRequest, opts ...gax.CallOption) *IndexIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListIndexes[0:len((*c.CallOptions).ListIndexes):len((*c.CallOptions).ListIndexes)], opts...)
	it := &IndexIterator{}
	req = proto.Clone(req).(*aiplatformpb.ListIndexesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*aiplatformpb.Index, string, error) {
		resp := &aiplatformpb.ListIndexesResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.indexClient.ListIndexes(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetIndexes(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *indexGRPCClient) UpdateIndex(ctx context.Context, req *aiplatformpb.UpdateIndexRequest, opts ...gax.CallOption) (*UpdateIndexOperation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "index.name", url.QueryEscape(req.GetIndex().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateIndex[0:len((*c.CallOptions).UpdateIndex):len((*c.CallOptions).UpdateIndex)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.indexClient.UpdateIndex(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &UpdateIndexOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *indexGRPCClient) DeleteIndex(ctx context.Context, req *aiplatformpb.DeleteIndexRequest, opts ...gax.CallOption) (*DeleteIndexOperation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteIndex[0:len((*c.CallOptions).DeleteIndex):len((*c.CallOptions).DeleteIndex)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.indexClient.DeleteIndex(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &DeleteIndexOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

// CreateIndexOperation manages a long-running operation from CreateIndex.
type CreateIndexOperation struct {
	lro *longrunning.Operation
}

// CreateIndexOperation returns a new CreateIndexOperation from a given name.
// The name must be that of a previously created CreateIndexOperation, possibly from a different process.
func (c *indexGRPCClient) CreateIndexOperation(name string) *CreateIndexOperation {
	return &CreateIndexOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *CreateIndexOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*aiplatformpb.Index, error) {
	var resp aiplatformpb.Index
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *CreateIndexOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*aiplatformpb.Index, error) {
	var resp aiplatformpb.Index
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *CreateIndexOperation) Metadata() (*aiplatformpb.CreateIndexOperationMetadata, error) {
	var meta aiplatformpb.CreateIndexOperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *CreateIndexOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *CreateIndexOperation) Name() string {
	return op.lro.Name()
}

// DeleteIndexOperation manages a long-running operation from DeleteIndex.
type DeleteIndexOperation struct {
	lro *longrunning.Operation
}

// DeleteIndexOperation returns a new DeleteIndexOperation from a given name.
// The name must be that of a previously created DeleteIndexOperation, possibly from a different process.
func (c *indexGRPCClient) DeleteIndexOperation(name string) *DeleteIndexOperation {
	return &DeleteIndexOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *DeleteIndexOperation) Wait(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.WaitWithInterval(ctx, nil, time.Minute, opts...)
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *DeleteIndexOperation) Poll(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.Poll(ctx, nil, opts...)
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *DeleteIndexOperation) Metadata() (*aiplatformpb.DeleteOperationMetadata, error) {
	var meta aiplatformpb.DeleteOperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *DeleteIndexOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *DeleteIndexOperation) Name() string {
	return op.lro.Name()
}

// UpdateIndexOperation manages a long-running operation from UpdateIndex.
type UpdateIndexOperation struct {
	lro *longrunning.Operation
}

// UpdateIndexOperation returns a new UpdateIndexOperation from a given name.
// The name must be that of a previously created UpdateIndexOperation, possibly from a different process.
func (c *indexGRPCClient) UpdateIndexOperation(name string) *UpdateIndexOperation {
	return &UpdateIndexOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *UpdateIndexOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*aiplatformpb.Index, error) {
	var resp aiplatformpb.Index
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *UpdateIndexOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*aiplatformpb.Index, error) {
	var resp aiplatformpb.Index
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *UpdateIndexOperation) Metadata() (*aiplatformpb.UpdateIndexOperationMetadata, error) {
	var meta aiplatformpb.UpdateIndexOperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *UpdateIndexOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *UpdateIndexOperation) Name() string {
	return op.lro.Name()
}

// IndexIterator manages a stream of *aiplatformpb.Index.
type IndexIterator struct {
	items    []*aiplatformpb.Index
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*aiplatformpb.Index, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *IndexIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *IndexIterator) Next() (*aiplatformpb.Index, error) {
	var item *aiplatformpb.Index
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *IndexIterator) bufLen() int {
	return len(it.items)
}

func (it *IndexIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
