package main

import (
	"github.com/smartystreets/goconvey/convey"
	_ "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingRoute(t *testing.T) {
	convey.Convey("Given a HTTP request for /invalid/123", t, func() {
		//req := httptest.NewRequest("GET", "/ping", nil)
		resp := httptest.NewRecorder()
		convey.Convey("When the request is handled by the Router", func() {
			convey.Convey("Then the response should be a 404", func() {
				convey.So(resp.Code, convey.ShouldEqual, 404)
			})
		})
	})
	router := setupRouter()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
