// Code generated by Kitex v0.10.3. DO NOT EDIT.

package followservice

import (
	relations "HuaTug.com/kitex_gen/relations"
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	RelationService(ctx context.Context, req *relations.RelationServiceRequest, callOptions ...callopt.Option) (r *relations.RelationServiceResponse, err error)
	FollowingList(ctx context.Context, req *relations.FollowingListRequest, callOptions ...callopt.Option) (r *relations.FollowingListResponse, err error)
	FollowerList(ctx context.Context, req *relations.FollowerListRequest, callOptions ...callopt.Option) (r *relations.FollowerListResponse, err error)
	FriendList(ctx context.Context, req *relations.FriendListRequest, callOptions ...callopt.Option) (r *relations.FriendListResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kFollowServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kFollowServiceClient struct {
	*kClient
}

func (p *kFollowServiceClient) RelationService(ctx context.Context, req *relations.RelationServiceRequest, callOptions ...callopt.Option) (r *relations.RelationServiceResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RelationService(ctx, req)
}

func (p *kFollowServiceClient) FollowingList(ctx context.Context, req *relations.FollowingListRequest, callOptions ...callopt.Option) (r *relations.FollowingListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FollowingList(ctx, req)
}

func (p *kFollowServiceClient) FollowerList(ctx context.Context, req *relations.FollowerListRequest, callOptions ...callopt.Option) (r *relations.FollowerListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FollowerList(ctx, req)
}

func (p *kFollowServiceClient) FriendList(ctx context.Context, req *relations.FriendListRequest, callOptions ...callopt.Option) (r *relations.FriendListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FriendList(ctx, req)
}