package eventstore_db

import (
	"context"
	"github.com/EventStore/EventStore-Client-Go/v2/esdb"
)

type IESDB interface {
	Close() error
	AppendToStream(
		context context.Context,
		streamID string,
		opts esdb.AppendToStreamOptions,
		events ...esdb.EventData,
	) (*esdb.WriteResult, error)
	SetStreamMetadata(
		context context.Context,
		streamID string,
		opts esdb.AppendToStreamOptions,
		metadata esdb.StreamMetadata,
	) (*esdb.WriteResult, error)
	GetStreamMetadata(
		context context.Context,
		streamID string,
		opts esdb.ReadStreamOptions,
	) (*esdb.StreamMetadata, error)
	DeleteStream(
		parent context.Context,
		streamID string,
		opts esdb.DeleteStreamOptions,
	) (*esdb.DeleteResult, error)
	TombstoneStream(
		parent context.Context,
		streamID string,
		opts esdb.TombstoneStreamOptions,
	) (*esdb.DeleteResult, error)
	ReadStream(
		context context.Context,
		streamID string,
		opts esdb.ReadStreamOptions,
		count uint64,
	) (*esdb.ReadStream, error)
	ReadAll(
		context context.Context,
		opts esdb.ReadAllOptions,
		count uint64,
	) (*esdb.ReadStream, error)
	SubscribeToStream(
		parent context.Context,
		streamID string,
		opts esdb.SubscribeToStreamOptions,
	) (*esdb.Subscription, error)
	SubscribeToAll(
		parent context.Context,
		opts esdb.SubscribeToAllOptions,
	) (*esdb.Subscription, error)
	SubscribeToPersistentSubscription(
		ctx context.Context,
		streamName string,
		groupName string,
		options esdb.SubscribeToPersistentSubscriptionOptions,
	) (*esdb.PersistentSubscription, error)
	SubscribeToPersistentSubscriptionToAll(
		ctx context.Context,
		groupName string,
		options esdb.SubscribeToPersistentSubscriptionOptions,
	) (*esdb.PersistentSubscription, error)
	CreatePersistentSubscription(
		ctx context.Context,
		streamName string,
		groupName string,
		options esdb.PersistentStreamSubscriptionOptions,
	) error
	CreatePersistentSubscriptionToAll(
		ctx context.Context,
		groupName string,
		options esdb.PersistentAllSubscriptionOptions,
	) error
	UpdatePersistentSubscription(
		ctx context.Context,
		streamName string,
		groupName string,
		options esdb.PersistentStreamSubscriptionOptions,
	) error
	UpdatePersistentSubscriptionToAll(
		ctx context.Context,
		groupName string,
		options esdb.PersistentAllSubscriptionOptions,
	) error
	DeletePersistentSubscription(
		ctx context.Context,
		streamName string,
		groupName string,
		options esdb.DeletePersistentSubscriptionOptions,
	) error
	DeletePersistentSubscriptionToAll(
		ctx context.Context,
		groupName string,
		options esdb.DeletePersistentSubscriptionOptions,
	) error
	ReplayParkedMessages(
		ctx context.Context,
		streamName string,
		groupName string,
		options esdb.ReplayParkedMessagesOptions) error
	ReplayParkedMessagesToAll(ctx context.Context, groupName string, options esdb.ReplayParkedMessagesOptions) error
	ListAllPersistentSubscriptions(ctx context.Context, options esdb.ListPersistentSubscriptionsOptions) ([]esdb.PersistentSubscriptionInfo, error)
	ListPersistentSubscriptionsForStream(ctx context.Context, streamName string, options esdb.ListPersistentSubscriptionsOptions) ([]esdb.PersistentSubscriptionInfo, error)
	ListPersistentSubscriptionsToAll(ctx context.Context, options esdb.ListPersistentSubscriptionsOptions) ([]esdb.PersistentSubscriptionInfo, error)
	GetPersistentSubscriptionInfo(ctx context.Context, streamName string, groupName string, options esdb.GetPersistentSubscriptionOptions) (*esdb.PersistentSubscriptionInfo, error)
	GetPersistentSubscriptionInfoToAll(ctx context.Context, groupName string, options esdb.GetPersistentSubscriptionOptions) (*esdb.PersistentSubscriptionInfo, error)
	RestartPersistentSubscriptionSubsystem(ctx context.Context, options esdb.RestartPersistentSubscriptionSubsystemOptions) error
}
