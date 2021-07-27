package myapp

import (
	"cnestWeb/utils"
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
	mux.ServeHTTP(res,req)
	assert.Equal(http.StatusOK, res.Code)	
	// if res.Code != http.StatusOK {
	// 	t.Fatal("failed!! ", res.Code)
	// }

	data, err := ioutil.ReadAll(res.Body)
	utils.HandleErr(err)
	assert.Equal("Hello World", string(data))
}

func TestBarPathHandler_WithoutName(t *testing.T) {
	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res,req)
	assert.Equal(http.StatusOK, res.Code)
	data, err := ioutil.ReadAll(res.Body)
	utils.HandleErr(err)
	assert.Equal("Hello world!", string(data))
}

func TestBarPathHandler_WithName(t *testing.T) {
	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar?name=Ann", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res,req)
	assert.Equal(http.StatusOK, res.Code)
	
	data, err := ioutil.ReadAll(res.Body)
	utils.HandleErr(err)
	assert.Equal("Hello Ann!", string(data))
}

func TestFooHandler_WithoutJson(t *testing.T) {
	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/foo", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusBadRequest, res.Code)
}

func TestFooHandler_WithJson(t *testing.T) {
	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/foo", 
		strings.NewReader(`{"firstname":"junghwa","lastname":"sung","email":"sundor@hanmail.net"}`))

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusCreated, res.Code)

	user := new(User)
	err := json.NewDecoder(res.Body).Decode(user)
	assert.Nil(err)
	assert.Equal("junghwa",user.FirstName)
	assert.Equal("sung",user.LastName)
}
