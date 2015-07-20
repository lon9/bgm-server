package main

import (
	"bgm-server/Godeps/_workspace/src/github.com/astaxie/beego/orm"
	"bgm-server/Godeps/_workspace/src/github.com/gorilla/mux"
	"bgm-server/utils"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

//Videoを一覧する関数order=likeでおすすめの多い順で取得
func videoIndex(w http.ResponseWriter, r *http.Request) {
	o := orm.NewOrm()

	params := r.URL.Query()
	page, _ := strconv.Atoi(params.Get("page"))
	maxResults, _ := strconv.Atoi(params.Get("maxResults"))
	order := params.Get("order")

	offset := maxResults * (page - 1)

	if page == 0 {
		page = 0
	}

	if maxResults == 0 {
		maxResults = 40
	}

	var videos Videos

	var err error
	//DBからデータを取ってくる
	if order == "like" {
		_, err = o.QueryTable("video").Limit(maxResults, offset).OrderBy("-liked").All(&videos)
	} else {
		_, err = o.QueryTable("video").Limit(maxResults, offset).OrderBy("id").All(&videos)
	}
	if err != nil {
		panic(err)
	}

	//Get thumbnails.
	for i, video := range videos {
		if video.HighThumbnail != nil {
			_, err := o.LoadRelated(&video, "HighThumbnail")
			if err != nil {
				panic(err)
			}
		}
		if video.MediumThumbnail != nil {
			_, err = o.LoadRelated(&video, "MediumThumbnail")
			if err != nil {
				panic(err)
			}
		}
		videos[i] = video
	}
	w.WriteHeader(200)

	response, err := json.Marshal(videos)
	utils.CheckError(w, err)
	w = utils.SetJSONHeader(w)

	fmt.Fprintln(w, string(response))
}

func videoUpdate(w http.ResponseWriter, r *http.Request) {
	o := orm.NewOrm()

	//パラメタからアップデートしたいvideoIdを取得
	vars := mux.Vars(r)
	videoID, _ := strconv.Atoi(vars["videoId"])
	video := Video{Id: videoID}

	//bodyからデータを取り出す。
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	checkError(w, err)
	if err := r.Body.Close(); err != nil {
		w.WriteHeader(400)
		utils.CheckError(w, err)
		panic(err)
	}

	w = utils.SetJSONHeader(w)

	var newVideo Video
	//jsonをパース
	if err := json.Unmarshal(body, &newVideo); err != nil {
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		panic(err)
	}

	if o.Read(&video) == nil {
		//データがぞんざいすればアップデートする
		video.VideoId = newVideo.VideoId
		video.Title = newVideo.Title
		video.Artist = newVideo.Artist
		if _, err := o.Update(&video); err == nil {
			w.WriteHeader(200)
			response, err := json.Marshal(video)
			utils.CheckError(w, err)
			fmt.Fprintln(w, string(response))
		}
	}
}
