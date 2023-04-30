package domain

import "fmt"

type SubscriberAction func(cs *ChannelSubscriber, channel *Channel, video *Video)

type ChannelSubscriber struct {
	name             string
	channels         []*Channel
	subscriberAction SubscriberAction
}

func (cs *ChannelSubscriber) Action(channel *Channel, video *Video) {
	cs.subscriberAction(cs, channel, video)
}

func (cs *ChannelSubscriber) GetName() string {
	return cs.name
}

func (cs *ChannelSubscriber) Subscribe(channel *Channel) {
	cs.channels = append(cs.channels, channel)
	channel.Register(cs)
	fmt.Printf("%s 訂閱了 %s。\n", cs.name, channel.name)
}

func (cs *ChannelSubscriber) Unsubscribe(channel *Channel) {

	for i, c := range cs.channels {
		if c.name == channel.name {
			cs.channels = append(cs.channels[:i], cs.channels[i+1:]...)
			fmt.Printf("%s 解除訂閱了 %s。\n", cs.name, channel.name)
			return
		}
	}

	channel.Unregister(cs.name)
}

func (cs *ChannelSubscriber) LikeVideo(v *Video) {
	fmt.Printf("%s 對影片 \"%s\" 按讚。\n", cs.name, v.Title)
}

func NewChannelSubscriber(name string, subscriberAction SubscriberAction) *ChannelSubscriber {
	return &ChannelSubscriber{
		name:             name,
		subscriberAction: subscriberAction,
	}
}
