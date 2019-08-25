package db

import (
	"fmt"
	"net/url"
	"strings"
)

type Channel struct {
	id     string
	source *url.URL
}

var channelsDB = make(map[string]*Channel)

// Channel is in the form of http://server:port/username/password/channel_id
func NewChannel(channelAddr string) (*Channel, error) {
	urlChannel, err := url.Parse(channelAddr)
	if err != nil {
		return nil, fmt.Errorf("Error parsing url: *s", channelAddr)
	}

	path := urlChannel.Path

	return &Channel{
		id:     path[strings.LastIndex(path, "/")+1:],
		source: urlChannel,
	}, nil
}

func ClearDB() {
	channelsDB = make(map[string]*Channel)
}

func RegisterChannel(m3uChannelAddr string) (*Channel, error) {
	channel, err := NewChannel(m3uChannelAddr)
	if err != nil {
		return nil, err
	}

	channelsDB[channel.id] = channel

	return channel, nil
}
