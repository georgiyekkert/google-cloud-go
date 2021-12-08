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

package dialogflow

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newConversationProfilesClientHook clientHook

// ConversationProfilesCallOptions contains the retry settings for each method of ConversationProfilesClient.
type ConversationProfilesCallOptions struct {
	ListConversationProfiles  []gax.CallOption
	GetConversationProfile    []gax.CallOption
	CreateConversationProfile []gax.CallOption
	UpdateConversationProfile []gax.CallOption
	DeleteConversationProfile []gax.CallOption
}

func defaultConversationProfilesGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("dialogflow.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("dialogflow.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://dialogflow.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultConversationProfilesCallOptions() *ConversationProfilesCallOptions {
	return &ConversationProfilesCallOptions{
		ListConversationProfiles: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetConversationProfile: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		CreateConversationProfile: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		UpdateConversationProfile: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		DeleteConversationProfile: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalConversationProfilesClient is an interface that defines the methods availaible from Dialogflow API.
type internalConversationProfilesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListConversationProfiles(context.Context, *dialogflowpb.ListConversationProfilesRequest, ...gax.CallOption) *ConversationProfileIterator
	GetConversationProfile(context.Context, *dialogflowpb.GetConversationProfileRequest, ...gax.CallOption) (*dialogflowpb.ConversationProfile, error)
	CreateConversationProfile(context.Context, *dialogflowpb.CreateConversationProfileRequest, ...gax.CallOption) (*dialogflowpb.ConversationProfile, error)
	UpdateConversationProfile(context.Context, *dialogflowpb.UpdateConversationProfileRequest, ...gax.CallOption) (*dialogflowpb.ConversationProfile, error)
	DeleteConversationProfile(context.Context, *dialogflowpb.DeleteConversationProfileRequest, ...gax.CallOption) error
}

// ConversationProfilesClient is a client for interacting with Dialogflow API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service for managing ConversationProfiles.
type ConversationProfilesClient struct {
	// The internal transport-dependent client.
	internalClient internalConversationProfilesClient

	// The call options for this service.
	CallOptions *ConversationProfilesCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ConversationProfilesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ConversationProfilesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *ConversationProfilesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListConversationProfiles returns the list of all conversation profiles in the specified project.
func (c *ConversationProfilesClient) ListConversationProfiles(ctx context.Context, req *dialogflowpb.ListConversationProfilesRequest, opts ...gax.CallOption) *ConversationProfileIterator {
	return c.internalClient.ListConversationProfiles(ctx, req, opts...)
}

// GetConversationProfile retrieves the specified conversation profile.
func (c *ConversationProfilesClient) GetConversationProfile(ctx context.Context, req *dialogflowpb.GetConversationProfileRequest, opts ...gax.CallOption) (*dialogflowpb.ConversationProfile, error) {
	return c.internalClient.GetConversationProfile(ctx, req, opts...)
}

// CreateConversationProfile creates a conversation profile in the specified project.
//
// ConversationProfile.CreateTime and ConversationProfile.UpdateTime
// aren’t populated in the response. You can retrieve them via
// GetConversationProfile API.
func (c *ConversationProfilesClient) CreateConversationProfile(ctx context.Context, req *dialogflowpb.CreateConversationProfileRequest, opts ...gax.CallOption) (*dialogflowpb.ConversationProfile, error) {
	return c.internalClient.CreateConversationProfile(ctx, req, opts...)
}

// UpdateConversationProfile updates the specified conversation profile.
//
// ConversationProfile.CreateTime and ConversationProfile.UpdateTime
// aren’t populated in the response. You can retrieve them via
// GetConversationProfile API.
func (c *ConversationProfilesClient) UpdateConversationProfile(ctx context.Context, req *dialogflowpb.UpdateConversationProfileRequest, opts ...gax.CallOption) (*dialogflowpb.ConversationProfile, error) {
	return c.internalClient.UpdateConversationProfile(ctx, req, opts...)
}

// DeleteConversationProfile deletes the specified conversation profile.
func (c *ConversationProfilesClient) DeleteConversationProfile(ctx context.Context, req *dialogflowpb.DeleteConversationProfileRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteConversationProfile(ctx, req, opts...)
}

// conversationProfilesGRPCClient is a client for interacting with Dialogflow API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type conversationProfilesGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing ConversationProfilesClient
	CallOptions **ConversationProfilesCallOptions

	// The gRPC API client.
	conversationProfilesClient dialogflowpb.ConversationProfilesClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewConversationProfilesClient creates a new conversation profiles client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service for managing ConversationProfiles.
func NewConversationProfilesClient(ctx context.Context, opts ...option.ClientOption) (*ConversationProfilesClient, error) {
	clientOpts := defaultConversationProfilesGRPCClientOptions()
	if newConversationProfilesClientHook != nil {
		hookOpts, err := newConversationProfilesClientHook(ctx, clientHookParams{})
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
	client := ConversationProfilesClient{CallOptions: defaultConversationProfilesCallOptions()}

	c := &conversationProfilesGRPCClient{
		connPool:                   connPool,
		disableDeadlines:           disableDeadlines,
		conversationProfilesClient: dialogflowpb.NewConversationProfilesClient(connPool),
		CallOptions:                &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *conversationProfilesGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *conversationProfilesGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *conversationProfilesGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *conversationProfilesGRPCClient) ListConversationProfiles(ctx context.Context, req *dialogflowpb.ListConversationProfilesRequest, opts ...gax.CallOption) *ConversationProfileIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListConversationProfiles[0:len((*c.CallOptions).ListConversationProfiles):len((*c.CallOptions).ListConversationProfiles)], opts...)
	it := &ConversationProfileIterator{}
	req = proto.Clone(req).(*dialogflowpb.ListConversationProfilesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*dialogflowpb.ConversationProfile, string, error) {
		resp := &dialogflowpb.ListConversationProfilesResponse{}
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
			resp, err = c.conversationProfilesClient.ListConversationProfiles(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetConversationProfiles(), resp.GetNextPageToken(), nil
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

func (c *conversationProfilesGRPCClient) GetConversationProfile(ctx context.Context, req *dialogflowpb.GetConversationProfileRequest, opts ...gax.CallOption) (*dialogflowpb.ConversationProfile, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetConversationProfile[0:len((*c.CallOptions).GetConversationProfile):len((*c.CallOptions).GetConversationProfile)], opts...)
	var resp *dialogflowpb.ConversationProfile
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.conversationProfilesClient.GetConversationProfile(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *conversationProfilesGRPCClient) CreateConversationProfile(ctx context.Context, req *dialogflowpb.CreateConversationProfileRequest, opts ...gax.CallOption) (*dialogflowpb.ConversationProfile, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateConversationProfile[0:len((*c.CallOptions).CreateConversationProfile):len((*c.CallOptions).CreateConversationProfile)], opts...)
	var resp *dialogflowpb.ConversationProfile
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.conversationProfilesClient.CreateConversationProfile(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *conversationProfilesGRPCClient) UpdateConversationProfile(ctx context.Context, req *dialogflowpb.UpdateConversationProfileRequest, opts ...gax.CallOption) (*dialogflowpb.ConversationProfile, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "conversation_profile.name", url.QueryEscape(req.GetConversationProfile().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateConversationProfile[0:len((*c.CallOptions).UpdateConversationProfile):len((*c.CallOptions).UpdateConversationProfile)], opts...)
	var resp *dialogflowpb.ConversationProfile
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.conversationProfilesClient.UpdateConversationProfile(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *conversationProfilesGRPCClient) DeleteConversationProfile(ctx context.Context, req *dialogflowpb.DeleteConversationProfileRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteConversationProfile[0:len((*c.CallOptions).DeleteConversationProfile):len((*c.CallOptions).DeleteConversationProfile)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.conversationProfilesClient.DeleteConversationProfile(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// ConversationProfileIterator manages a stream of *dialogflowpb.ConversationProfile.
type ConversationProfileIterator struct {
	items    []*dialogflowpb.ConversationProfile
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
	InternalFetch func(pageSize int, pageToken string) (results []*dialogflowpb.ConversationProfile, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *ConversationProfileIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ConversationProfileIterator) Next() (*dialogflowpb.ConversationProfile, error) {
	var item *dialogflowpb.ConversationProfile
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ConversationProfileIterator) bufLen() int {
	return len(it.items)
}

func (it *ConversationProfileIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
