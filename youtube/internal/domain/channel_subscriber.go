package domain

type SubscriberAction func(name string, channel *Channel, video *Video)

type channelSubscriber struct {
	name             string
	subscriberAction SubscriberAction
}

func (cs *channelSubscriber) Action(channel *Channel, video *Video) {
	cs.subscriberAction(cs.name, channel, video)
}

func (cs *channelSubscriber) GetName() string {
	return cs.name
}

func NewChannelSubscriber(name string, subscriberAction SubscriberAction) ChannelObserver {
	return &channelSubscriber{
		name:             name,
		subscriberAction: subscriberAction,
	}
}
