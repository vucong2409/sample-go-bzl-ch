package router

import "net/http"

func Ping(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte(`{"pong": "pong"}`))
}
