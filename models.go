package main

import (
	"bgm-server/Godeps/_workspace/src/github.com/astaxie/beego/orm"
	"time"
)

// Inquery is mel.
type Inquery struct {
	Id      int       `json:"id"`
	Content string    `json:"content"`
	Created time.Time `json:"created"`
}

// Video is model of YouTube videos.
type Video struct {
	Id            int       `json:"id"`
	VideoId       string    `json:"videoId" orm:"unique"`
	Title         string    `json:"title"`
	Artist        string    `json:"artist,omitempty" orm:"null"`
	ThumbnailUrl  string    `json:"url,omitempty" orm:"null"`
	PublishedDate time.Time `json:"publishedDate,omitempty" orm:"null"`
	Created       time.Time `json:"created" orm:"auto_now_add;type(datetime);null"`
	Updated       time.Time `json:"updated" orm:"auto_now;type(datetime);null"`
	Liked         int       `json:"liked"`
}

// Inqueries are models.
type Inqueries []Inquery

// Videos is sequence of Video.
type Videos []Video

func init() {
	orm.RegisterModel(new(Inquery), new(Video))

	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		panic(err)
	}
}
