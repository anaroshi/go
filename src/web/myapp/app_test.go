package myapp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexPathHandler(t *testing.T) {
	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	//indexHandler(res, req)
	// if res.Code != http.StatusOK {
	// 	t.Fatal("Failed!!", res.Code)
	// }
	assert.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hellow World", string(data))
}

func TestBarHandler_WithoutName(t *testing.T) {
	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hellow world!", string(data))
}

func TestBarHandler_WithName(t *testing.T) {
	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar?name=sundor", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hellow sundor!", string(data))
}

func TestFooHandler_WithoutJson(t *testing.T) {
	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/foo", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	// json 값을 주지 않았기 때문에 당연히 에러다
	assert.Equal(http.StatusBadRequest, res.Code)
}

func TestFooHandler_WithJson(t *testing.T) {
	//json 처리
	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/foo", strings.NewReader(`{"first_name":"sundor","last_name":"sung","email":"sundor@hanmail.net"}`))

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusCreated, res.Code)

	user := new(User)
	err := json.NewDecoder(res.Body).Decode(user)
	assert.Nil(err)
	assert.Equal("sundor", user.FirstName)
	assert.Equal("sung", user.LastName)
	assert.Equal("sundor@hanmail.net", user.Email)

}
