package db

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

var rgx = *regexp.MustCompile(`group-title="(?P<GroupTitle>.*)"`)

type Groups struct {
	title string
}

type Channel struct {
	Id         string
	meta       string
	Source     *url.URL
	GroupTitle string
}

var groupsTable = make(map[string][]*Channel)
var channelTable = make(map[string]*Channel)

// Channel is in the form of http://server:port/username/password/channel_id
func NewChannel(extInf string, channelAddr string) (*Channel, error) {
	urlChannel, err := url.Parse(channelAddr)
	if err != nil {
		return nil, fmt.Errorf("error parsing url: %s", channelAddr)
	}

	// get group-title from #extinf line
	match := rgx.FindStringSubmatch(extInf)
	var groupTitle = "NONE"
	// match = [group-title="MOVIES" MOVIES]
	if len(match) >= 2 {
		groupTitle = match[1]
	}

	path := urlChannel.Path

	return &Channel{
		Id:         path[strings.LastIndex(path, "/")+1:], // extract only the channel_id
		meta:       extInf,
		Source:     urlChannel,
		GroupTitle: groupTitle,
	}, nil
}

func Reset() {
	channelTable = make(map[string]*Channel)
}

func RegisterChannel(channelMeta string, channelAddr string) (channel *Channel, err error) {
	channel, err = NewChannel(channelMeta, channelAddr)
	if err != nil {
		return
	}

	channelTable[channel.Id] = channel
	groupsTable[channel.GroupTitle] = append(groupsTable[channel.GroupTitle], channel)

	return
}

func GetAllTitles() (keys []string) {
	keys = make([]string, 0, len(groupsTable))
	for k := range groupsTable {
		keys = append(keys, k)
	}
	return
}

func GetChannel(id string) (channel *Channel, err error) {
	channel = channelTable[id]
	if channel == nil {
		err = fmt.Errorf("No channel available with id: %s ", id)
	}
	return
}

func GetAllChannelsWithGroupTitle(groupTitle string) (channels []*Channel) {

}
