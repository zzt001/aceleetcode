package crawler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	BASE_URL = `http://leetcode.com/api/problems/`
	DATA_URL = BASE_URL + `algorithms/`
)

func CrawlData() (*Dto, error) {
	resp, err := http.Get(DATA_URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var data *Dto
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
