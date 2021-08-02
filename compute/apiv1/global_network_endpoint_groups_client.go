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
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"

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

var newGlobalNetworkEndpointGroupsClientHook clientHook

// GlobalNetworkEndpointGroupsCallOptions contains the retry settings for each method of GlobalNetworkEndpointGroupsClient.
type GlobalNetworkEndpointGroupsCallOptions struct {
	AttachNetworkEndpoints []gax.CallOption
	Delete                 []gax.CallOption
	DetachNetworkEndpoints []gax.CallOption
	Get                    []gax.CallOption
	Insert                 []gax.CallOption
	List                   []gax.CallOption
	ListNetworkEndpoints   []gax.CallOption
}

// internalGlobalNetworkEndpointGroupsClient is an interface that defines the methods availaible from Google Compute Engine API.
type internalGlobalNetworkEndpointGroupsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	AttachNetworkEndpoints(context.Context, *computepb.AttachNetworkEndpointsGlobalNetworkEndpointGroupRequest, ...gax.CallOption) (*Operation, error)
	Delete(context.Context, *computepb.DeleteGlobalNetworkEndpointGroupRequest, ...gax.CallOption) (*Operation, error)
	DetachNetworkEndpoints(context.Context, *computepb.DetachNetworkEndpointsGlobalNetworkEndpointGroupRequest, ...gax.CallOption) (*Operation, error)
	Get(context.Context, *computepb.GetGlobalNetworkEndpointGroupRequest, ...gax.CallOption) (*computepb.NetworkEndpointGroup, error)
	Insert(context.Context, *computepb.InsertGlobalNetworkEndpointGroupRequest, ...gax.CallOption) (*Operation, error)
	List(context.Context, *computepb.ListGlobalNetworkEndpointGroupsRequest, ...gax.CallOption) *NetworkEndpointGroupIterator
	ListNetworkEndpoints(context.Context, *computepb.ListNetworkEndpointsGlobalNetworkEndpointGroupsRequest, ...gax.CallOption) *NetworkEndpointWithHealthStatusIterator
}

// GlobalNetworkEndpointGroupsClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The GlobalNetworkEndpointGroups API.
type GlobalNetworkEndpointGroupsClient struct {
	// The internal transport-dependent client.
	internalClient internalGlobalNetworkEndpointGroupsClient

	// The call options for this service.
	CallOptions *GlobalNetworkEndpointGroupsCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *GlobalNetworkEndpointGroupsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *GlobalNetworkEndpointGroupsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *GlobalNetworkEndpointGroupsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// AttachNetworkEndpoints attach a network endpoint to the specified network endpoint group.
func (c *GlobalNetworkEndpointGroupsClient) AttachNetworkEndpoints(ctx context.Context, req *computepb.AttachNetworkEndpointsGlobalNetworkEndpointGroupRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.AttachNetworkEndpoints(ctx, req, opts...)
}

// Delete deletes the specified network endpoint group.Note that the NEG cannot be deleted if there are backend services referencing it.
func (c *GlobalNetworkEndpointGroupsClient) Delete(ctx context.Context, req *computepb.DeleteGlobalNetworkEndpointGroupRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// DetachNetworkEndpoints detach the network endpoint from the specified network endpoint group.
func (c *GlobalNetworkEndpointGroupsClient) DetachNetworkEndpoints(ctx context.Context, req *computepb.DetachNetworkEndpointsGlobalNetworkEndpointGroupRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.DetachNetworkEndpoints(ctx, req, opts...)
}

// Get returns the specified network endpoint group. Gets a list of available network endpoint groups by making a list() request.
func (c *GlobalNetworkEndpointGroupsClient) Get(ctx context.Context, req *computepb.GetGlobalNetworkEndpointGroupRequest, opts ...gax.CallOption) (*computepb.NetworkEndpointGroup, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// Insert creates a network endpoint group in the specified project using the parameters that are included in the request.
func (c *GlobalNetworkEndpointGroupsClient) Insert(ctx context.Context, req *computepb.InsertGlobalNetworkEndpointGroupRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List retrieves the list of network endpoint groups that are located in the specified project.
func (c *GlobalNetworkEndpointGroupsClient) List(ctx context.Context, req *computepb.ListGlobalNetworkEndpointGroupsRequest, opts ...gax.CallOption) *NetworkEndpointGroupIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// ListNetworkEndpoints lists the network endpoints in the specified network endpoint group.
func (c *GlobalNetworkEndpointGroupsClient) ListNetworkEndpoints(ctx context.Context, req *computepb.ListNetworkEndpointsGlobalNetworkEndpointGroupsRequest, opts ...gax.CallOption) *NetworkEndpointWithHealthStatusIterator {
	return c.internalClient.ListNetworkEndpoints(ctx, req, opts...)
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type globalNetworkEndpointGroupsRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewGlobalNetworkEndpointGroupsRESTClient creates a new global network endpoint groups rest client.
//
// The GlobalNetworkEndpointGroups API.
func NewGlobalNetworkEndpointGroupsRESTClient(ctx context.Context, opts ...option.ClientOption) (*GlobalNetworkEndpointGroupsClient, error) {
	clientOpts := append(defaultGlobalNetworkEndpointGroupsRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	c := &globalNetworkEndpointGroupsRESTClient{
		endpoint:   endpoint,
		httpClient: httpClient,
	}
	c.setGoogleClientInfo()

	return &GlobalNetworkEndpointGroupsClient{internalClient: c, CallOptions: &GlobalNetworkEndpointGroupsCallOptions{}}, nil
}

func defaultGlobalNetworkEndpointGroupsRESTClientOptions() []option.ClientOption {
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
func (c *globalNetworkEndpointGroupsRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *globalNetworkEndpointGroupsRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *globalNetworkEndpointGroupsRESTClient) Connection() *grpc.ClientConn {
	return nil
}

// AttachNetworkEndpoints attach a network endpoint to the specified network endpoint group.
func (c *globalNetworkEndpointGroupsRESTClient) AttachNetworkEndpoints(ctx context.Context, req *computepb.AttachNetworkEndpointsGlobalNetworkEndpointGroupRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetGlobalNetworkEndpointGroupsAttachEndpointsRequestResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/networkEndpointGroups/%v/attachNetworkEndpoints", req.GetProject(), req.GetNetworkEndpointGroup())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
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

// Delete deletes the specified network endpoint group.Note that the NEG cannot be deleted if there are backend services referencing it.
func (c *globalNetworkEndpointGroupsRESTClient) Delete(ctx context.Context, req *computepb.DeleteGlobalNetworkEndpointGroupRequest, opts ...gax.CallOption) (*Operation, error) {
	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/networkEndpointGroups/%v", req.GetProject(), req.GetNetworkEndpointGroup())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

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
	rsp := &computepb.Operation{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, err
	}
	op := &Operation{proto: rsp}
	return op, err
}

// DetachNetworkEndpoints detach the network endpoint from the specified network endpoint group.
func (c *globalNetworkEndpointGroupsRESTClient) DetachNetworkEndpoints(ctx context.Context, req *computepb.DetachNetworkEndpointsGlobalNetworkEndpointGroupRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetGlobalNetworkEndpointGroupsDetachEndpointsRequestResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/networkEndpointGroups/%v/detachNetworkEndpoints", req.GetProject(), req.GetNetworkEndpointGroup())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
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

// Get returns the specified network endpoint group. Gets a list of available network endpoint groups by making a list() request.
func (c *globalNetworkEndpointGroupsRESTClient) Get(ctx context.Context, req *computepb.GetGlobalNetworkEndpointGroupRequest, opts ...gax.CallOption) (*computepb.NetworkEndpointGroup, error) {
	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/networkEndpointGroups/%v", req.GetProject(), req.GetNetworkEndpointGroup())

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
	rsp := &computepb.NetworkEndpointGroup{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// Insert creates a network endpoint group in the specified project using the parameters that are included in the request.
func (c *globalNetworkEndpointGroupsRESTClient) Insert(ctx context.Context, req *computepb.InsertGlobalNetworkEndpointGroupRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetNetworkEndpointGroupResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/networkEndpointGroups", req.GetProject())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
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

// List retrieves the list of network endpoint groups that are located in the specified project.
func (c *globalNetworkEndpointGroupsRESTClient) List(ctx context.Context, req *computepb.ListGlobalNetworkEndpointGroupsRequest, opts ...gax.CallOption) *NetworkEndpointGroupIterator {
	it := &NetworkEndpointGroupIterator{}
	req = proto.Clone(req).(*computepb.ListGlobalNetworkEndpointGroupsRequest)
	m := protojson.MarshalOptions{AllowPartial: true, UseProtoNames: false}
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.NetworkEndpointGroup, string, error) {
		resp := &computepb.NetworkEndpointGroupList{}
		if pageSize > math.MaxInt32 {
			req.MaxResults = proto.Uint32(math.MaxInt32)
		} else {
			req.MaxResults = proto.Uint32(uint32(pageSize))
		}
		req.PageToken = proto.String(pageToken)

		jsonReq, err := m.Marshal(req)
		if err != nil {
			return nil, "", err
		}

		baseUrl, _ := url.Parse(c.endpoint)
		baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/networkEndpointGroups", req.GetProject())

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

		httpReq, err := http.NewRequest("GET", baseUrl.String(), bytes.NewReader(jsonReq))
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

// ListNetworkEndpoints lists the network endpoints in the specified network endpoint group.
func (c *globalNetworkEndpointGroupsRESTClient) ListNetworkEndpoints(ctx context.Context, req *computepb.ListNetworkEndpointsGlobalNetworkEndpointGroupsRequest, opts ...gax.CallOption) *NetworkEndpointWithHealthStatusIterator {
	it := &NetworkEndpointWithHealthStatusIterator{}
	req = proto.Clone(req).(*computepb.ListNetworkEndpointsGlobalNetworkEndpointGroupsRequest)
	m := protojson.MarshalOptions{AllowPartial: true, UseProtoNames: false}
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.NetworkEndpointWithHealthStatus, string, error) {
		resp := &computepb.NetworkEndpointGroupsListNetworkEndpoints{}
		if pageSize > math.MaxInt32 {
			req.MaxResults = proto.Uint32(math.MaxInt32)
		} else {
			req.MaxResults = proto.Uint32(uint32(pageSize))
		}
		req.PageToken = proto.String(pageToken)

		jsonReq, err := m.Marshal(req)
		if err != nil {
			return nil, "", err
		}

		baseUrl, _ := url.Parse(c.endpoint)
		baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/networkEndpointGroups/%v/listNetworkEndpoints", req.GetProject(), req.GetNetworkEndpointGroup())

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

		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
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

// NetworkEndpointGroupIterator manages a stream of *computepb.NetworkEndpointGroup.
type NetworkEndpointGroupIterator struct {
	items    []*computepb.NetworkEndpointGroup
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
	InternalFetch func(pageSize int, pageToken string) (results []*computepb.NetworkEndpointGroup, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *NetworkEndpointGroupIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *NetworkEndpointGroupIterator) Next() (*computepb.NetworkEndpointGroup, error) {
	var item *computepb.NetworkEndpointGroup
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *NetworkEndpointGroupIterator) bufLen() int {
	return len(it.items)
}

func (it *NetworkEndpointGroupIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// NetworkEndpointWithHealthStatusIterator manages a stream of *computepb.NetworkEndpointWithHealthStatus.
type NetworkEndpointWithHealthStatusIterator struct {
	items    []*computepb.NetworkEndpointWithHealthStatus
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
	InternalFetch func(pageSize int, pageToken string) (results []*computepb.NetworkEndpointWithHealthStatus, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *NetworkEndpointWithHealthStatusIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *NetworkEndpointWithHealthStatusIterator) Next() (*computepb.NetworkEndpointWithHealthStatus, error) {
	var item *computepb.NetworkEndpointWithHealthStatus
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *NetworkEndpointWithHealthStatusIterator) bufLen() int {
	return len(it.items)
}

func (it *NetworkEndpointWithHealthStatusIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
