package translate

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"fmt"
	"strings"
)

type Translate struct {
}

var DefaultTranslate = &Translate{}

func (t *Translate) doSend(method, urlStr string, body []byte, isJson bool) ([]byte, error) {
	var r io.Reader
	if body != nil && len(body) > 0 {
		r = bytes.NewReader(body)
	}

	req, err := http.NewRequest(method, urlStr, r)
	if err != nil {
		return nil, err
	}

	if isJson {
		req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET code=%d, msg=%s", resp.StatusCode, resp.Status)
	}

	if isJson {
		if !strings.Contains(strings.ToLower(resp.Header.Get("Content-Type")), "application/json") {
			return nil, fmt.Errorf("返回数据格式包非json格式。")
		}
	}

	return ioutil.ReadAll(resp.Body)
}

func (t *Translate) Send(method, urlStr string, body []byte) ([]byte, error) {
	return t.doSend(method, urlStr, body, false)
}

func (t *Translate) SendJsonData(method, urlStr string, body []byte) ([]byte, error) {
	return t.doSend(method, urlStr, body, true)
}

func (t *Translate) SendWithJson(method, urlStr string, body interface{}, result interface{}) error {
	var err error
	var bs []byte
	if body != nil {
		bs, err = json.Marshal(body)
		if err != nil {
			return err
		}
	}

	bs, err = t.doSend(method, urlStr, bs, true)
	if err != nil {
		return err
	}

	if result != nil {
		return json.Unmarshal(bs, result)
	}

	return nil
}

// func (t *Translate) PatchJsonData(urlStr string, body []byte) ([]byte, error) {
// 	return t.Send("PATCH", urlStr, body)
// }
// func (t *Translate) PutJsonData(urlStr string, body []byte) ([]byte, error) {
// 	return t.SendJsonData("PUT", urlStr, body, true)
// }

// func (t *Translate) PatchWithJson(urlStr string, body interface{}, result interface{}) error {
// 	return t.SendWithJson("PATCH", urlStr, body, result)
// }

// func (t *Translate) PutWithJson(urlStr string, body interface{}, result interface{}) error {
// 	return t.SendWithJson("PUT", urlStr, body, result)
// }

// func (t *Translate) GetWithJson(urlStr string, result interface{}) error {
// 	return t.SendWithJson("GET", urlStr, nil, result)
// }

// func (t *Translate) PostWithJson(urlStr string, body interface{}, result interface{}) error {
// 	return t.SendWithJson("POST", urlStr, body, result)
// }

// func (t *Translate) Get(urlStr string) ([]byte, error) {
// 	return t.Send("GET", urlStr, nil)
// }

// func (t *Translate) PostJsonData(urlStr string, data []byte) ([]byte, error) {
// 	resp, err := http.Post(urlStr, "application/json", bytes.NewReader(data))
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return nil, fmt.Errorf("POST code=%d, msg=%s", resp.StatusCode, resp.Status)
// 	}

// 	return t.readBodyFromResponse(resp)
// }
