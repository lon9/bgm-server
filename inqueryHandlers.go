package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"bgm-server/Godeps/_workspace/src/github.com/astaxie/beego/orm"
	"bgm-server/Godeps/_workspace/src/github.com/gorilla/mux"
)

func inqueryIndex(w http.ResponseWriter, r *http.Request) {
	o := orm.NewOrm()

	params := r.URL.Query()
	page, _ := strconv.Atoi(params.Get("page"))
	maxResults, _ := strconv.Atoi(params.Get("maxResults"))
	dataType := params.Get("dataType")

	offset := maxResults * (page - 1)

	if page == 0 {
		page = 0
	}

	if maxResults == 0 {
		maxResults = 40
	}

	var inqueries Inqueries

	_, err := o.QueryTable("inquery").Limit(maxResults, offset).OrderBy("id").All(&inqueries)
	checkError(w, err)

	w.WriteHeader(200)

	if dataType == "json" {
		response, err := json.Marshal(inqueries)
		if err != nil {
			fmt.Fprintln(w, err)
		}
		w = setHeader(w)
		fmt.Fprintln(w, string(response))
	} else {
		//こっからテンプレート
		data := map[string]Inqueries{
			"Inqueries": inqueries,
		}
		tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/inquery/index.html"))
		w.Header().Set("Content-Type", "text/html")

		err = tmpl.Execute(w, data)
		if err != nil {
			panic(err)
		}
	}
}

func inqueryShow(w http.ResponseWriter, r *http.Request) {
	o := orm.NewOrm()

	params := r.URL.Query()
	dataType := params.Get("dataType")

	vars := mux.Vars(r)
	inqueryID, _ := strconv.Atoi(vars["inqueryId"])
	inquery := Inquery{Id: inqueryID}
	err := o.Read(&inquery)

	if err == orm.ErrNoRows {
		w.WriteHeader(204)
		fmt.Fprintln(w, "No result found.")
	} else if err == orm.ErrMissPK {
		fmt.Fprintln(w, "No primary key found.")
	} else {
		w.WriteHeader(200)
		if dataType == "json" {
			response, err := json.Marshal(inquery)
			checkError(w, err)
			w = setHeader(w)
			fmt.Fprintln(w, string(response))
		} else {
			//こっからテンプレート
			data := map[string]Inquery{
				"Inquery": inquery,
			}
			tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/inquery/show.html"))
			w.Header().Set("Content-Type", "text/html")

			err = tmpl.Execute(w, data)
			if err != nil {
				panic(err)
			}

		}
	}
}

func inqueryCreate(w http.ResponseWriter, r *http.Request) {
	o := orm.NewOrm()

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	checkError(w, err)
	if err := r.Body.Close(); err != nil {
		w.WriteHeader(400)
		checkError(w, err)
	}

	w = setHeader(w)

	var inquery Inquery
	if err := json.Unmarshal(body, &inquery); err != nil {
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	inquery.Created = time.Now()
	_, err = o.Insert(&inquery)
	if err == nil {
		w.WriteHeader(201)
	}
}

func inqueryDelete(w http.ResponseWriter, r *http.Request) {
	o := orm.NewOrm()

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["inqueryId"])
	w = setHeader(w)

	if err != nil {
		w.WriteHeader(400)
	}

	if _, err = o.Delete(&Inquery{Id: id}); err == nil {
		w.WriteHeader(204)
	}
}

func checkError(w http.ResponseWriter, err error) {
	if err != nil {
		fmt.Fprintln(w, err)
	}
}

func setHeader(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	return w
}
