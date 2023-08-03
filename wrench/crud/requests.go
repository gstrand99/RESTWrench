package wrench

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ResponseInfo struct {
  Status string
  Body string
  Headers http.Header
}

func (ri *ResponseInfo) Print() {
	fmt.Println("Status:", ri.Status)
  fmt.Println(" ")
  fmt.Println("Headers:", ri.Headers)
	fmt.Println(" ")
  fmt.Println("Body:", ri.Body)
  fmt.Println(" ")
}

func Get(url string) (*ResponseInfo, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	responseInfo := &ResponseInfo{
		Status:  resp.Status,
		Body:    string(bodyBytes),
		Headers: resp.Header,
	}

	return responseInfo, nil
}

func Post(url string, body interface{}) (*ResponseInfo, error) {
  jsonBody, err := json.Marshal(body)
  if err != nil {
    return nil, err 
  }

  resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
  if err != nil {
    return nil, err 
  }
  defer resp.Body.Close()

  respBody, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err 
  }

  responseInfo := &ResponseInfo{
    Status: resp.Status,
    Body: string(respBody),
    Headers: resp.Header,
  }

  return responseInfo, nil

}

func Put(url string, body interface{}) (*ResponseInfo, error) {
  client := &http.Client{}
  jsonBody, err := json.Marshal(body)
  if err != nil {
    return nil, err 
  }

  req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonBody))
  if err != nil {
    return nil, err 
  }
  
  resp, err := client.Do(req)
  if err != nil {
    return nil, err
  }

  defer resp.Body.Close()

  respBody, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err 
  }

  responseInfo := &ResponseInfo{
    Status: resp.Status,
    Body: string(respBody),
    Headers: resp.Header,
  }

  return responseInfo, nil
}

func Delete(url string) (*ResponseInfo, error) {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	responseInfo := &ResponseInfo{
		Status:  resp.Status,
		Body:    string(bodyBytes),
		Headers: resp.Header,
	}

	return responseInfo, nil
}

