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

package compute

import (
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"sort"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	httptransport "google.golang.org/api/transport/http"
	computepb "google.golang.org/genproto/googleapis/cloud/compute/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newGlobalOperationsClientHook clientHook

// GlobalOperationsCallOptions contains the retry settings for each method of GlobalOperationsClient.
type GlobalOperationsCallOptions struct {
	AggregatedList []gax.CallOption
	Delete         []gax.CallOption
	Get            []gax.CallOption
	List           []gax.CallOption
	Wait           []gax.CallOption
}

// internalGlobalOperationsClient is an interface that defines the methods availaible from Google Compute Engine API.
type internalGlobalOperationsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	AggregatedList(context.Context, *computepb.AggregatedListGlobalOperationsRequest, ...gax.CallOption) *OperationsScopedListPairIterator
	Delete(context.Context, *computepb.DeleteGlobalOperationRequest, ...gax.CallOption) (*computepb.DeleteGlobalOperationResponse, error)
	Get(context.Context, *computepb.GetGlobalOperationRequest, ...gax.CallOption) (*Operation, error)
	List(context.Context, *computepb.ListGlobalOperationsRequest, ...gax.CallOption) *OperationIterator
	Wait(context.Context, *computepb.WaitGlobalOperationRequest, ...gax.CallOption) (*Operation, error)
}

// GlobalOperationsClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The GlobalOperations API.
type GlobalOperationsClient struct {
	// The internal transport-dependent client.
	internalClient internalGlobalOperationsClient

	// The call options for this service.
	CallOptions *GlobalOperationsCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *GlobalOperationsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *GlobalOperationsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *GlobalOperationsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// AggregatedList retrieves an aggregated list of all operations.
func (c *GlobalOperationsClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListGlobalOperationsRequest, opts ...gax.CallOption) *OperationsScopedListPairIterator {
	return c.internalClient.AggregatedList(ctx, req, opts...)
}

// Delete deletes the specified Operations resource.
func (c *GlobalOperationsClient) Delete(ctx context.Context, req *computepb.DeleteGlobalOperationRequest, opts ...gax.CallOption) (*computepb.DeleteGlobalOperationResponse, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// Get retrieves the specified Operations resource. Gets a list of operations by making a list() request.
func (c *GlobalOperationsClient) Get(ctx context.Context, req *computepb.GetGlobalOperationRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// List retrieves a list of Operation resources contained within the specified project.
func (c *GlobalOperationsClient) List(ctx context.Context, req *computepb.ListGlobalOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// Wait waits for the specified Operation resource to return as DONE or for the request to approach the 2 minute deadline, and retrieves the specified Operation resource. This method differs from the GET method in that it waits for no more than the default deadline (2 minutes) and then returns the current state of the operation, which might be DONE or still in progress.
//
// This method is called on a best-effort basis. Specifically:
//
//   In uncommon cases, when the server is overloaded, the request might return before the default deadline is reached, or might return after zero seconds.
//
//   If the default deadline is reached, there is no guarantee that the operation is actually done when the method returns. Be prepared to retry if the operation is not DONE.
func (c *GlobalOperationsClient) Wait(ctx context.Context, req *computepb.WaitGlobalOperationRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Wait(ctx, req, opts...)
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type globalOperationsRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewGlobalOperationsRESTClient creates a new global operations rest client.
//
// The GlobalOperations API.
func NewGlobalOperationsRESTClient(ctx context.Context, opts ...option.ClientOption) (*GlobalOperationsClient, error) {
	clientOpts := append(defaultGlobalOperationsRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	c := &globalOperationsRESTClient{
		endpoint:   endpoint,
		httpClient: httpClient,
	}
	c.setGoogleClientInfo()

	return &GlobalOperationsClient{internalClient: c, CallOptions: &GlobalOperationsCallOptions{}}, nil
}

func defaultGlobalOperationsRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://compute.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://compute.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://compute.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *globalOperationsRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *globalOperationsRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *globalOperationsRESTClient) Connection() *grpc.ClientConn {
	return nil
}

// AggregatedList retrieves an aggregated list of all operations.
func (c *globalOperationsRESTClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListGlobalOperationsRequest, opts ...gax.CallOption) *OperationsScopedListPairIterator {
	it := &OperationsScopedListPairIterator{}
	req = proto.Clone(req).(*computepb.AggregatedListGlobalOperationsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]OperationsScopedListPair, string, error) {
		resp := &computepb.OperationAggregatedList{}
		if pageSize > math.MaxInt32 {
			req.MaxResults = proto.Uint32(math.MaxInt32)
		} else if pageSize != 0 {
			req.MaxResults = proto.Uint32(uint32(pageSize))
		}
		if pageToken != "" {
			req.PageToken = proto.String(pageToken)
		}

		baseUrl, _ := url.Parse(c.endpoint)
		baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/aggregated/operations", req.GetProject())

		params := url.Values{}
		if req != nil && req.Filter != nil {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req != nil && req.IncludeAllScopes != nil {
			params.Add("includeAllScopes", fmt.Sprintf("%v", req.GetIncludeAllScopes()))
		}
		if req != nil && req.MaxResults != nil {
			params.Add("maxResults", fmt.Sprintf("%v", req.GetMaxResults()))
		}
		if req != nil && req.OrderBy != nil {
			params.Add("orderBy", fmt.Sprintf("%v", req.GetOrderBy()))
		}
		if req != nil && req.PageToken != nil {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}
		if req != nil && req.ReturnPartialSuccess != nil {
			params.Add("returnPartialSuccess", fmt.Sprintf("%v", req.GetReturnPartialSuccess()))
		}

		baseUrl.RawQuery = params.Encode()

		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return nil, "", err
		}

		// Set the headers
		for k, v := range c.xGoogMetadata {
			httpReq.Header[k] = v
		}

		httpReq.Header["Content-Type"] = []string{"application/json"}
		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return nil, "", err
		}
		defer httpRsp.Body.Close()

		if httpRsp.StatusCode != http.StatusOK {
			return nil, "", fmt.Errorf(httpRsp.Status)
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return nil, "", err
		}

		unm.Unmarshal(buf, resp)
		it.Response = resp

		elems := make([]OperationsScopedListPair, 0, len(resp.GetItems()))
		for k, v := range resp.GetItems() {
			elems = append(elems, OperationsScopedListPair{k, v})
		}
		sort.Slice(elems, func(i, j int) bool { return elems[i].Key < elems[j].Key })

		return elems, resp.GetNextPageToken(), nil
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
	it.pageInfo.MaxSize = int(req.GetMaxResults())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// Delete deletes the specified Operations resource.
func (c *globalOperationsRESTClient) Delete(ctx context.Context, req *computepb.DeleteGlobalOperationRequest, opts ...gax.CallOption) (*computepb.DeleteGlobalOperationResponse, error) {
	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/operations/%v", req.GetProject(), req.GetOperation())

	httpReq, err := http.NewRequest("DELETE", baseUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.DeleteGlobalOperationResponse{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// Get retrieves the specified Operations resource. Gets a list of operations by making a list() request.
func (c *globalOperationsRESTClient) Get(ctx context.Context, req *computepb.GetGlobalOperationRequest, opts ...gax.CallOption) (*Operation, error) {
	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/operations/%v", req.GetProject(), req.GetOperation())

	httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, err
	}
	op := &Operation{proto: rsp}
	return op, err
}

// List retrieves a list of Operation resources contained within the specified project.
func (c *globalOperationsRESTClient) List(ctx context.Context, req *computepb.ListGlobalOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	it := &OperationIterator{}
	req = proto.Clone(req).(*computepb.ListGlobalOperationsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.Operation, string, error) {
		resp := &computepb.OperationList{}
		if pageSize > math.MaxInt32 {
			req.MaxResults = proto.Uint32(math.MaxInt32)
		} else if pageSize != 0 {
			req.MaxResults = proto.Uint32(uint32(pageSize))
		}
		if pageToken != "" {
			req.PageToken = proto.String(pageToken)
		}

		baseUrl, _ := url.Parse(c.endpoint)
		baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/operations", req.GetProject())

		params := url.Values{}
		if req != nil && req.Filter != nil {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req != nil && req.MaxResults != nil {
			params.Add("maxResults", fmt.Sprintf("%v", req.GetMaxResults()))
		}
		if req != nil && req.OrderBy != nil {
			params.Add("orderBy", fmt.Sprintf("%v", req.GetOrderBy()))
		}
		if req != nil && req.PageToken != nil {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}
		if req != nil && req.ReturnPartialSuccess != nil {
			params.Add("returnPartialSuccess", fmt.Sprintf("%v", req.GetReturnPartialSuccess()))
		}

		baseUrl.RawQuery = params.Encode()

		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return nil, "", err
		}

		// Set the headers
		for k, v := range c.xGoogMetadata {
			httpReq.Header[k] = v
		}

		httpReq.Header["Content-Type"] = []string{"application/json"}
		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return nil, "", err
		}
		defer httpRsp.Body.Close()

		if httpRsp.StatusCode != http.StatusOK {
			return nil, "", fmt.Errorf(httpRsp.Status)
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return nil, "", err
		}

		unm.Unmarshal(buf, resp)
		it.Response = resp
		return resp.GetItems(), resp.GetNextPageToken(), nil
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
	it.pageInfo.MaxSize = int(req.GetMaxResults())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// Wait waits for the specified Operation resource to return as DONE or for the request to approach the 2 minute deadline, and retrieves the specified Operation resource. This method differs from the GET method in that it waits for no more than the default deadline (2 minutes) and then returns the current state of the operation, which might be DONE or still in progress.
//
// This method is called on a best-effort basis. Specifically:
//
//   In uncommon cases, when the server is overloaded, the request might return before the default deadline is reached, or might return after zero seconds.
//
//   If the default deadline is reached, there is no guarantee that the operation is actually done when the method returns. Be prepared to retry if the operation is not DONE.
func (c *globalOperationsRESTClient) Wait(ctx context.Context, req *computepb.WaitGlobalOperationRequest, opts ...gax.CallOption) (*Operation, error) {
	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/operations/%v/wait", req.GetProject(), req.GetOperation())

	httpReq, err := http.NewRequest("POST", baseUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, err
	}
	op := &Operation{proto: rsp}
	return op, err
}

// OperationIterator manages a stream of *computepb.Operation.
type OperationIterator struct {
	items    []*computepb.Operation
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
	InternalFetch func(pageSize int, pageToken string) (results []*computepb.Operation, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *OperationIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *OperationIterator) Next() (*computepb.Operation, error) {
	var item *computepb.Operation
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *OperationIterator) bufLen() int {
	return len(it.items)
}

func (it *OperationIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// OperationsScopedListPair is a holder type for string/*computepb.OperationsScopedList map entries
type OperationsScopedListPair struct {
	Key   string
	Value *computepb.OperationsScopedList
}

// OperationsScopedListPairIterator manages a stream of OperationsScopedListPair.
type OperationsScopedListPairIterator struct {
	items    []OperationsScopedListPair
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
	InternalFetch func(pageSize int, pageToken string) (results []OperationsScopedListPair, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *OperationsScopedListPairIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *OperationsScopedListPairIterator) Next() (OperationsScopedListPair, error) {
	var item OperationsScopedListPair
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *OperationsScopedListPairIterator) bufLen() int {
	return len(it.items)
}

func (it *OperationsScopedListPairIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
