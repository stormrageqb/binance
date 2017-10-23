package binance

import (
	"context"
)

type PingService struct {
	c *Client
}

func (s *PingService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   "GET",
		endpoint: "/api/v1/ping",
	}
	_, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	return
}

type ServerTimeService struct {
	c *Client
}

func (s *ServerTimeService) Do(ctx context.Context, opts ...RequestOption) (serverTime int64, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/api/v1/time",
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	j, err := newJSON(data)
	if err != nil {
		return
	}
	serverTime = j.Get("serverTime").MustInt64()
	return
}
