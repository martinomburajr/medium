package main

import (
	"testing"
)

func TestBaseHandler(t *testing.T) {
	//recorder := httptest.NewRecorder()
	//
	//server := httptest.NewServer(Router())
	//
	//type args struct {
	//	writer  http.ResponseWriter
	//	request *http.Request
	//}
	//tests := []struct {
	//	name string
	//	args args
	//	wantResponseCode int
	//}{
	//	{"test-post-request", args{recorder, httptest.NewRequest(http.MethodGet, server.URL + "/", nil)}, 400},
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		server.Start()
	//		defer server.Close()
	//
	//		BaseHandler(tt.args.writer, tt.args.request)
	//
	//		if tt.wantResponseCode != tt.args.writer
	//	})
	//}
}
