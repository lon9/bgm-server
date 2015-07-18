package main

import (
	"bgm-server/Godeps/_workspace/src/github.com/astaxie/beego/orm"
	"time"
)

// Inquery is mel.
type Inquery struct {
	Id      int       `json: "id"`
	Content string    `json: "content"`
	Created time.Time `json: "created"`
}

// Video is model of YouTube videos.
type Video struct {
	Id              int        `json: "id"`
	VideoId         string     `json: "videoId" orm: "unique"`
	Title           string     `json: "title"`
	Artist          string     `json: "artist"`
	Highthumbnail   *Thumbnail `json: "highThumbnail" orm: "rel(one);on_delete(set_null)"`
	MediumThumbnail *Thumbnail `json: "mediumThumbnail" orm: "rel(one);on_delete(set_null)"`
	PublishedDate   time.Time  `json: "publishedDate"`
	Created         time.Time  `json: "created" orm: "auto_now_add;type(datetime)"`
	Updated         time.Time  `json: "updated" orm: "auto_now;type(datetime)"`
	liked           int        `json: "liked"`
}

// Thumbnail is model of thumbnail of Video.
type Thumbnail struct {
	Id     int    `json: "id"`
	URL    string `json: "url"`
	Width  int    `json: "width"`
	Height int    `json: "height"`
}

// Inqueries are models.
type Inqueries []Inquery

// Videos is sequence of Video.
type Videos []Video

func init() {
	orm.RegisterModel(new(Inquery))
	orm.RegisterModel(new(Video))
	orm.RegisterModel(new(Thumbnail))

	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		panic(err)
	}
}
