package db

import (
	"fmt"
	"net/url"
	"strings"
)

type Channel struct {
	Id     string
	Source *url.URL
}

var channelsDB = make(map[string]*Channel)

// Channel is in the form of http://server:port/username/password/channel_id
func NewChannel(channelAddr string) (*Channel, error) {
	urlChannel, err := url.Parse(channelAddr)
	if err != nil {
		return nil, fmt.Errorf("error parsing url: %s", channelAddr)
	}

	path := urlChannel.Path

	return &Channel{
		Id:     path[strings.LastIndex(path, "/")+1:],
		Source: urlChannel,
	}, nil
}

func Reset() {
	channelsDB = make(map[string]*Channel)
}

func RegisterChannel(channelAddr string) (channel *Channel, err error) {
	channel, err = NewChannel(channelAddr)
	if err != nil {
		return
	}

	channelsDB[channel.Id] = channel
	return
}

	return channel, nil
}
