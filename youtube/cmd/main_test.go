package main

import (
	"time"

	"youtube/internal/domain"
)

func Example() {
	pew := domain.NewChannel("PewDiePie")
	water := domain.NewChannel("水球軟體學院")

	p1 := domain.NewChannelSubscriber("水球", func(name string, channel *domain.Channel, video *domain.Video) {
		if video.Length >= 3*time.Minute {
			video.Liked(name)
		}
	})

	p2 := domain.NewChannelSubscriber("火球", func(name string, channel *domain.Channel, video *domain.Video) {
		if video.Length <= 1*time.Minute {
			channel.Unsubscribed(name)
		}
	})

	water.Subscribed(p1)
	pew.Subscribed(p1)

	water.Subscribed(p2)
	pew.Subscribed(p2)

	water.Upload(domain.NewVideo("C1M1S2", "這個世界正是物件導向的呢！", 4*time.Minute))
	pew.Upload(domain.NewVideo("Hello guys", "Clickbait", 30*time.Second))
	water.Upload(domain.NewVideo("C1M1S3", "物件 vs. 類別", 1*time.Minute))
	pew.Upload(domain.NewVideo("Minecraft", "Let’s play Minecraft", 30*time.Minute))

	// Output:
	// 水球 訂閱了 水球軟體學院。
	// 水球 訂閱了 PewDiePie。
	// 火球 訂閱了 水球軟體學院。
	// 火球 訂閱了 PewDiePie。
	// 頻道 水球軟體學院 上架了一則新影片 "C1M1S2"。
	// 水球 對影片 "C1M1S2" 按讚。
	// 頻道 PewDiePie 上架了一則新影片 "Hello guys"。
	// 火球 解除訂閱了 PewDiePie。
	// 頻道 水球軟體學院 上架了一則新影片 "C1M1S3"。
	// 火球 解除訂閱了 水球軟體學院。
	// 頻道 PewDiePie 上架了一則新影片 "Minecraft"。
	// 水球 對影片 "Minecraft" 按讚。
}
