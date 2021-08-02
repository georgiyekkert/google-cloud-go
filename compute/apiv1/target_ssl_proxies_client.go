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

var newTargetSslProxiesClientHook clientHook

// TargetSslProxiesCallOptions contains the retry settings for each method of TargetSslProxiesClient.
type TargetSslProxiesCallOptions struct {
	Delete             []gax.CallOption
	Get                []gax.CallOption
	Insert             []gax.CallOption
	List               []gax.CallOption
	SetBackendService  []gax.CallOption
	SetProxyHeader     []gax.CallOption
	SetSslCertificates []gax.CallOption
	SetSslPolicy       []gax.CallOption
}

// internalTargetSslProxiesClient is an interface that defines the methods availaible from Google Compute Engine API.
type internalTargetSslProxiesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	Delete(context.Context, *computepb.DeleteTargetSslProxyRequest, ...gax.CallOption) (*Operation, error)
	Get(context.Context, *computepb.GetTargetSslProxyRequest, ...gax.CallOption) (*computepb.TargetSslProxy, error)
	Insert(context.Context, *computepb.InsertTargetSslProxyRequest, ...gax.CallOption) (*Operation, error)
	List(context.Context, *computepb.ListTargetSslProxiesRequest, ...gax.CallOption) *TargetSslProxyIterator
	SetBackendService(context.Context, *computepb.SetBackendServiceTargetSslProxyRequest, ...gax.CallOption) (*Operation, error)
	SetProxyHeader(context.Context, *computepb.SetProxyHeaderTargetSslProxyRequest, ...gax.CallOption) (*Operation, error)
	SetSslCertificates(context.Context, *computepb.SetSslCertificatesTargetSslProxyRequest, ...gax.CallOption) (*Operation, error)
	SetSslPolicy(context.Context, *computepb.SetSslPolicyTargetSslProxyRequest, ...gax.CallOption) (*Operation, error)
}

// TargetSslProxiesClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The TargetSslProxies API.
type TargetSslProxiesClient struct {
	// The internal transport-dependent client.
	internalClient internalTargetSslProxiesClient

	// The call options for this service.
	CallOptions *TargetSslProxiesCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *TargetSslProxiesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *TargetSslProxiesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *TargetSslProxiesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// Delete deletes the specified TargetSslProxy resource.
func (c *TargetSslProxiesClient) Delete(ctx context.Context, req *computepb.DeleteTargetSslProxyRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// Get returns the specified TargetSslProxy resource. Gets a list of available target SSL proxies by making a list() request.
func (c *TargetSslProxiesClient) Get(ctx context.Context, req *computepb.GetTargetSslProxyRequest, opts ...gax.CallOption) (*computepb.TargetSslProxy, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// Insert creates a TargetSslProxy resource in the specified project using the data included in the request.
func (c *TargetSslProxiesClient) Insert(ctx context.Context, req *computepb.InsertTargetSslProxyRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List retrieves the list of TargetSslProxy resources available to the specified project.
func (c *TargetSslProxiesClient) List(ctx context.Context, req *computepb.ListTargetSslProxiesRequest, opts ...gax.CallOption) *TargetSslProxyIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// SetBackendService changes the BackendService for TargetSslProxy.
func (c *TargetSslProxiesClient) SetBackendService(ctx context.Context, req *computepb.SetBackendServiceTargetSslProxyRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.SetBackendService(ctx, req, opts...)
}

// SetProxyHeader changes the ProxyHeaderType for TargetSslProxy.
func (c *TargetSslProxiesClient) SetProxyHeader(ctx context.Context, req *computepb.SetProxyHeaderTargetSslProxyRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.SetProxyHeader(ctx, req, opts...)
}

// SetSslCertificates changes SslCertificates for TargetSslProxy.
func (c *TargetSslProxiesClient) SetSslCertificates(ctx context.Context, req *computepb.SetSslCertificatesTargetSslProxyRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.SetSslCertificates(ctx, req, opts...)
}

// SetSslPolicy sets the SSL policy for TargetSslProxy. The SSL policy specifies the server-side support for SSL features. This affects connections between clients and the SSL proxy load balancer. They do not affect the connection between the load balancer and the backends.
func (c *TargetSslProxiesClient) SetSslPolicy(ctx context.Context, req *computepb.SetSslPolicyTargetSslProxyRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.SetSslPolicy(ctx, req, opts...)
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type targetSslProxiesRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewTargetSslProxiesRESTClient creates a new target ssl proxies rest client.
//
// The TargetSslProxies API.
func NewTargetSslProxiesRESTClient(ctx context.Context, opts ...option.ClientOption) (*TargetSslProxiesClient, error) {
	clientOpts := append(defaultTargetSslProxiesRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	c := &targetSslProxiesRESTClient{
		endpoint:   endpoint,
		httpClient: httpClient,
	}
	c.setGoogleClientInfo()

	return &TargetSslProxiesClient{internalClient: c, CallOptions: &TargetSslProxiesCallOptions{}}, nil
}

func defaultTargetSslProxiesRESTClientOptions() []option.ClientOption {
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
func (c *targetSslProxiesRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *targetSslProxiesRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *targetSslProxiesRESTClient) Connection() *grpc.ClientConn {
	return nil
}

// Delete deletes the specified TargetSslProxy resource.
func (c *targetSslProxiesRESTClient) Delete(ctx context.Context, req *computepb.DeleteTargetSslProxyRequest, opts ...gax.CallOption) (*Operation, error) {
	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/targetSslProxies/%v", req.GetProject(), req.GetTargetSslProxy())

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

// Get returns the specified TargetSslProxy resource. Gets a list of available target SSL proxies by making a list() request.
func (c *targetSslProxiesRESTClient) Get(ctx context.Context, req *computepb.GetTargetSslProxyRequest, opts ...gax.CallOption) (*computepb.TargetSslProxy, error) {
	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/targetSslProxies/%v", req.GetProject(), req.GetTargetSslProxy())

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
	rsp := &computepb.TargetSslProxy{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// Insert creates a TargetSslProxy resource in the specified project using the data included in the request.
func (c *targetSslProxiesRESTClient) Insert(ctx context.Context, req *computepb.InsertTargetSslProxyRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetTargetSslProxyResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/targetSslProxies", req.GetProject())

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

// List retrieves the list of TargetSslProxy resources available to the specified project.
func (c *targetSslProxiesRESTClient) List(ctx context.Context, req *computepb.ListTargetSslProxiesRequest, opts ...gax.CallOption) *TargetSslProxyIterator {
	it := &TargetSslProxyIterator{}
	req = proto.Clone(req).(*computepb.ListTargetSslProxiesRequest)
	m := protojson.MarshalOptions{AllowPartial: true, UseProtoNames: false}
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.TargetSslProxy, string, error) {
		resp := &computepb.TargetSslProxyList{}
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
		baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/targetSslProxies", req.GetProject())

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

// SetBackendService changes the BackendService for TargetSslProxy.
func (c *targetSslProxiesRESTClient) SetBackendService(ctx context.Context, req *computepb.SetBackendServiceTargetSslProxyRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetTargetSslProxiesSetBackendServiceRequestResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/targetSslProxies/%v/setBackendService", req.GetProject(), req.GetTargetSslProxy())

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

// SetProxyHeader changes the ProxyHeaderType for TargetSslProxy.
func (c *targetSslProxiesRESTClient) SetProxyHeader(ctx context.Context, req *computepb.SetProxyHeaderTargetSslProxyRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetTargetSslProxiesSetProxyHeaderRequestResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/targetSslProxies/%v/setProxyHeader", req.GetProject(), req.GetTargetSslProxy())

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

// SetSslCertificates changes SslCertificates for TargetSslProxy.
func (c *targetSslProxiesRESTClient) SetSslCertificates(ctx context.Context, req *computepb.SetSslCertificatesTargetSslProxyRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetTargetSslProxiesSetSslCertificatesRequestResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/targetSslProxies/%v/setSslCertificates", req.GetProject(), req.GetTargetSslProxy())

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

// SetSslPolicy sets the SSL policy for TargetSslProxy. The SSL policy specifies the server-side support for SSL features. This affects connections between clients and the SSL proxy load balancer. They do not affect the connection between the load balancer and the backends.
func (c *targetSslProxiesRESTClient) SetSslPolicy(ctx context.Context, req *computepb.SetSslPolicyTargetSslProxyRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetSslPolicyReferenceResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/targetSslProxies/%v/setSslPolicy", req.GetProject(), req.GetTargetSslProxy())

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

// TargetSslProxyIterator manages a stream of *computepb.TargetSslProxy.
type TargetSslProxyIterator struct {
	items    []*computepb.TargetSslProxy
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
	InternalFetch func(pageSize int, pageToken string) (results []*computepb.TargetSslProxy, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *TargetSslProxyIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *TargetSslProxyIterator) Next() (*computepb.TargetSslProxy, error) {
	var item *computepb.TargetSslProxy
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *TargetSslProxyIterator) bufLen() int {
	return len(it.items)
}

func (it *TargetSslProxyIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
