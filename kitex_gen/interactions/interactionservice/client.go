// Code generated by Kitex v0.10.3. DO NOT EDIT.

package interactionservice

import (
	interactions "HuaTug.com/kitex_gen/interactions"
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	LikeAction(ctx context.Context, req *interactions.LikeActionRequest, callOptions ...callopt.Option) (r *interactions.LikeActionResponse, err error)
	LikeList(ctx context.Context, req *interactions.LikeListRequest, callOptions ...callopt.Option) (r *interactions.LikeListResponse, err error)
	CreateComment(ctx context.Context, req *interactions.CreateCommentRequest, callOptions ...callopt.Option) (r *interactions.CreateCommentResponse, err error)
	ListComment(ctx context.Context, req *interactions.ListCommentRequest, callOptions ...callopt.Option) (r *interactions.ListCommentResponse, err error)
	DeleteComment(ctx context.Context, req *interactions.CommentDeleteRequest, callOptions ...callopt.Option) (r *interactions.CommentDeleteResponse, err error)
	VideoPopularList(ctx context.Context, req *interactions.VideoPopularListRequest, callOptions ...callopt.Option) (r *interactions.VideoPopularListResponse, err error)
	DeleteVideoInfo(ctx context.Context, req *interactions.DeleteVideoInfoRequest, callOptions ...callopt.Option) (r *interactions.DeleteVideoInfoResponse, err error)
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
	return &kInteractionServiceClient{
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

type kInteractionServiceClient struct {
	*kClient
}

func (p *kInteractionServiceClient) LikeAction(ctx context.Context, req *interactions.LikeActionRequest, callOptions ...callopt.Option) (r *interactions.LikeActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.LikeAction(ctx, req)
}

func (p *kInteractionServiceClient) LikeList(ctx context.Context, req *interactions.LikeListRequest, callOptions ...callopt.Option) (r *interactions.LikeListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.LikeList(ctx, req)
}

func (p *kInteractionServiceClient) CreateComment(ctx context.Context, req *interactions.CreateCommentRequest, callOptions ...callopt.Option) (r *interactions.CreateCommentResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateComment(ctx, req)
}

func (p *kInteractionServiceClient) ListComment(ctx context.Context, req *interactions.ListCommentRequest, callOptions ...callopt.Option) (r *interactions.ListCommentResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListComment(ctx, req)
}

func (p *kInteractionServiceClient) DeleteComment(ctx context.Context, req *interactions.CommentDeleteRequest, callOptions ...callopt.Option) (r *interactions.CommentDeleteResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteComment(ctx, req)
}

func (p *kInteractionServiceClient) VideoPopularList(ctx context.Context, req *interactions.VideoPopularListRequest, callOptions ...callopt.Option) (r *interactions.VideoPopularListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.VideoPopularList(ctx, req)
}

func (p *kInteractionServiceClient) DeleteVideoInfo(ctx context.Context, req *interactions.DeleteVideoInfoRequest, callOptions ...callopt.Option) (r *interactions.DeleteVideoInfoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteVideoInfo(ctx, req)
}