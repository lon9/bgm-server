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
	w = utils.SetJSONHeader(w)

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

	w.WriteHeader(200)
	response, err := json.Marshal(videos)
	utils.CheckError(w, err)
	fmt.Fprintln(w, string(response))
}

func videoUpdate(w http.ResponseWriter, r *http.Request) {
	w = utils.SetJSONHeader(w)
	if utils.CheckAuth(r) == false {
		w.Header().Set("WWW-Authenticate", `Basic realm="MY REALM"`)
		w.WriteHeader(401)
		w.Write([]byte("401 Unauthorized\n"))
		return
	}

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

func videoDelete(w http.ResponseWriter, r *http.Request) {
	w = utils.SetJSONHeader(w)
	if utils.CheckAuth(r) == false {
		w.Header().Set("WWW-Authenticate", `Basic realm="MY REALM"`)
		w.WriteHeader(401)
		w.Write([]byte("401 Unauthorized\n"))
		return
	}

	o := orm.NewOrm()
	vars := mux.Vars(r)
	videoID, err := strconv.Atoi(vars["videoId"])

	w = utils.SetJSONHeader(w)

	if err != nil {
		w.WriteHeader(400)
		return
	}
	video := Video{Id: videoID}

	if _, err := o.Delete(&video); err == nil {
		w.WriteHeader(204)
	} else {
		w.WriteHeader(404)
	}
}

func videoOptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST, DELETE")
}
