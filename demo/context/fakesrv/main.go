package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

type Results struct {
	ResponseData struct {
		Results []Content
	}
}

// A Result contains the title and URL of a search result.
type Content struct {
	Title, SubTitle string
}

func main() {

	http.HandleFunc("/context_demo", handleContext)
	http.ListenAndServe(":9000", nil)
}

func handleContext(resp http.ResponseWriter, req *http.Request) {
	defer func() {
		if e := recover(); e != nil {
			if msg, ok := e.(string); ok {
				resp.Write([]byte(msg))
			} else {
				panic(e)
			}
		}
	}()
	check_error := func(err error, msg string) {
		if err != nil {
			if msg != "" {
				panic(err.Error() + ":" + msg)
			} else {
				panic(err.Error())
			}
		}
	}
	if req.Method == "GET" {
		q := req.FormValue("q")
		seg := strings.Split(q, ":")
		if len(seg) < 2 {
			log.Println("query format wrong")
			resp.Write([]byte("query format wrong"))
			return
		}
		title := seg[0]
		num, err := strconv.Atoi(seg[1])
		check_error(err, "")
		rs := Results{}
		for i := 0; i < num; i++ {
			rs.ResponseData.Results = append(rs.ResponseData.Results,
				Content{fmt.Sprintf("%s %d", title, i), RandomString(20)})
		}
		buff := bytes.NewBuffer(nil)
		err = json.NewEncoder(buff).Encode(rs)
		check_error(err, "")
		time.Sleep(time.Second * 2)
		resp.Write(buff.Bytes())
	} else {
		resp.Write([]byte("请使用get方法!"))
	}
}
func RandomString(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
