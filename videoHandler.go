package main

import (
	"bgm-server/Godeps/_workspace/src/github.com/astaxie/beego/orm"
	"bgm-server/utils"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

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
		_, err = o.QueryTable("video").Limit(maxResults, offset).OrderBy("liked").All(&videos)
	} else {
		_, err = o.QueryTable("video").Limit(maxResults, offset).OrderBy("id").All(&videos)
	}
	utils.CheckError(w, err)

	w.WriteHeader(200)

	response, err := json.Marshal(videos)
	utils.CheckError(w, err)
	w = utils.SetJSONHeader(w)

	fmt.Fprintln(w, string(response))
}

func videoUpdate(w http.ResponseWriter, r *http.Request) {
	o := orm.NewOrm()

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	checkError(w, err)
	if err := r.Body.Close(); err != nil {
		w.WriteHeader(400)
		utils.CheckError(w, err)
	}

	w = utils.SetJSONHeader(w)

	var video Video

	//jsonをパース
	if err := json.Unmarshal(body, &video); err != nil {
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			utils.CheckError(w, err)
		}
	}

	//DBに保存
	if _, _, err := o.ReadOrCreate(&video, "videoId"); err == nil {
		video.Liked = video.Liked + 1
		if _, err := o.Update(&video); err == nil {
			w.WriteHeader(200)
			fmt.Fprintln(w, "更新しました。")
		}
	}
}
