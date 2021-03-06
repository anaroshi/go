package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexPage(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Equal("Hello World", string(data))

}

func TestDecoHandler(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	buf := &bytes.Buffer{}
	log.SetOutput(buf)

	resp, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	r := bufio.NewReader(buf) // 버퍼에 저장
	line, _, err := r.ReadLine() // 버퍼에서 한줄씩 읽겠당
	assert.NoError(err)
	assert.Contains(string(line),"[LOGGER2] Started")

	line, _, err = r.ReadLine()
	assert.NoError(err)
	assert.Contains(string(line), "[LOGGER1] Started")

	line, _, err = r.ReadLine()
	assert.NoError(err)
	assert.Contains(string(line),"[LOGGER1] Completed")

	line, _, err = r.ReadLine()
	assert.NoError(err)
	assert.Contains(string(line),"[LOGGER2] Completed")
}