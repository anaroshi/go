package myapp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())	
	defer ts.Close()
	
	resp, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	data, _ := ioutil.ReadAll(resp.Body)
	assert.Equal("Hello Ann", string(data))

}

func TestUsers(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())	
	defer ts.Close()
	
	resp, err := http.Get(ts.URL+"/users")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	data, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(string(data), "No Users")
}

func TestGetUserInfo(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())	
	defer ts.Close()
	
	resp, err := http.Get(ts.URL+"/users/89")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "No User Id:89")
}


func TestCreateUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())	
	defer ts.Close()
	
	resp, err := http.Post(ts.URL+"/users", "application/json", strings.NewReader(`{"first_name":"junghwa","last_name":"sung","email":"sundor@hanmail.net"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	user := new(User)
	err = json.NewDecoder(resp.Body).Decode(user)
	assert.Nil(err)
	assert.NotEqual(0, user.ID)
	assert.Equal("junghwa", user.FirstName)
	assert.Equal("sung",user.LastName)

	id := user.ID
	resp, err = http.Get(ts.URL+"/users/"+strconv.Itoa(id))
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	user2 := new(User)
	err = json.NewDecoder(resp.Body).Decode(user2)
	assert.NoError(err)
	assert.Equal(user.ID, user2.ID)
	assert.Equal(user.FirstName, user2.FirstName)
}


func TestDeleteUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	req, _ := http.NewRequest("DELETE", ts.URL+"/users/1",nil)
	resp, err := http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode) 
	// data, _ := ioutil.ReadAll(resp.Body)
	// log.Print(data)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "No User ID:1")

	// 등록
	resp, err = http.Post(ts.URL+"/users", "application/json", 
			strings.NewReader(`{"first_name":"junghwa","last_name":"sung","email":"sundor@hanmail.net"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)
	
	// 등록 확인
	user := new(User)
	err = json.NewDecoder(resp.Body).Decode(user)
	assert.Nil(err)
	assert.NotEqual(0, user.ID)

	// 삭제
	req, _ = http.NewRequest("DELETE", ts.URL+"/users/1",nil)
	resp, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode) 
	data, _ = ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "Deleted User ID:1")
}

func TestUpdateUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	// Update
	req, _ := http.NewRequest("PUT", ts.URL+"/users", 
			strings.NewReader(`{"id":1, "first_name":"updated","last_name":"updated","email":"updated@hanmail.net"}`))
	resp, err := http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "No User ID:1")

	// 등록
	resp, err = http.Post(ts.URL+"/users", "application/json", 
			strings.NewReader(`{"first_name":"junghwa","last_name":"sung","email":"sundor@hanmail.net"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	// 등록 확인
	user := new(User)
	err = json.NewDecoder(resp.Body).Decode(user)
	assert.Nil(err)
	assert.NotEqual(0, user.ID)

	
	updateStr := fmt.Sprintf(`{"id":%d, "first_name":"jason"}`, user.ID)

	// Update
	req, _ = http.NewRequest("PUT", ts.URL+"/users", 
			strings.NewReader(updateStr))
	resp, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	updateUser := new(User)
	err = json.NewDecoder(resp.Body).Decode(updateUser)
	assert.NoError(err)
	assert.Equal(updateUser.ID, user.ID)
	assert.Equal("jason", updateUser.FirstName)
	assert.Equal(user.LastName, updateUser.LastName)
	assert.Equal(user.Email, updateUser.Email)
}

func TestUsers_WithUsersData(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())	
	defer ts.Close()

	// 등록
	resp, err := http.Post(ts.URL+"/users", "application/json", 
			strings.NewReader(`{"first_name":"junghwa","last_name":"sung","email":"sundor@hanmail.net"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	// 등록
	resp, err = http.Post(ts.URL+"/users", "application/json", 
			strings.NewReader(`{"first_name":"jason","last_name":"park","email":"jason@hanmail.net"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	// get Users
	resp, err = http.Get(ts.URL+"/users")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	
	users := []*User{}
	err = json.NewDecoder(resp.Body).Decode(&users)
	assert.NoError(err)
	assert.Equal(2, len(users))


}