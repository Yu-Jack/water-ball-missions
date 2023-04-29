package domain

import "fmt"

type ChannelObserver interface {
	Action(channel *Channel, video *Video)
	GetName() string
}

type Channel struct {
	name        string
	videos      []*Video
	subscribers []ChannelObserver
}

func NewChannel(name string) *Channel {
	return &Channel{
		name: name,
	}
}

func (c *Channel) Upload(video *Video) {
	fmt.Printf("頻道 %s 上架了一則新影片 \"%s\"。\n", c.name, video.Title)
	c.videos = append(c.videos, video)
	c.notify(video)
}

func (c *Channel) notify(video *Video) {
	for _, s := range c.subscribers {
		s.Action(c, video)
	}
}

func (c *Channel) Register(subscriber ChannelObserver) {
	fmt.Printf("%s 訂閱了 %s。\n", subscriber.GetName(), c.name)
	c.subscribers = append(c.subscribers, subscriber)
}

func (c *Channel) Unregister(name string) {
	for i, sub := range c.subscribers {
		if sub.GetName() == name {
			c.subscribers = append(c.subscribers[:i], c.subscribers[i+1:]...)
			fmt.Printf("%s 解除訂閱了 %s。\n", name, c.name)
			return
		}
	}
}
