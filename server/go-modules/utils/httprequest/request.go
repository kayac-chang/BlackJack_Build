package httprequest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func Get(urlpath string, timeout time.Duration) ([]byte, error) {
	client := http.Client{
		Timeout: timeout,
	}

	resp, err := client.Get(urlpath)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", string(body))
	}

	return body, nil
}

func PostForm(urlpath string, timeout time.Duration, data url.Values) ([]byte, error) {
	client := http.Client{
		Timeout: timeout,
	}

	resp, err := client.PostForm(urlpath, data)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", string(body))
	}

	return body, nil
}

func Put(urlpath string, timeout time.Duration, data url.Values) ([]byte, error) {
	client := http.Client{
		Timeout: timeout,
	}

	body := strings.NewReader(data.Encode())

	req, _ := http.NewRequest(http.MethodPut, urlpath, body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", string(response))
	}

	return response, nil
}

func Post(urlpath string, timeout time.Duration, data url.Values) ([]byte, error) {
	client := http.Client{
		Timeout: timeout,
	}

	body := strings.NewReader(data.Encode())
	req, err := http.NewRequest(http.MethodPost, urlpath, body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", string(response))
	}

	return response, nil
}
