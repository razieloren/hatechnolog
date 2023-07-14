package youtube

import (
	"context"
	"fmt"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type Service struct {
	session *youtube.Service
}

func NewService(apiKey string) (*Service, error) {
	ctx := context.Background()
	session, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("new_service: %w", err)
	}
	return &Service{
		session: session,
	}, nil
}

func (service *Service) QueryChannelStats(params *ChannelQueryParams) (*ChannelStats, error) {
	channelResponse, err :=
		service.session.Channels.List(listChannelParts).ForUsername(params.Name).Do()
	if err != nil {
		return nil, fmt.Errorf("do: %w", err)
	}
	return &ChannelStats{
		ChannelName: params.Name,
		Subscribers: channelResponse.Items[0].Statistics.SubscriberCount,
		Views:       channelResponse.Items[0].Statistics.ViewCount,
	}, nil
}
