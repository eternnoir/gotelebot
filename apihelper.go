package gotelebot

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/eternnoir/gotelebot/types"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

func sendGetRequest(method string, token string, params url.Values) ([]byte, error) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/%s?%s",
		token, method, params.Encode())

	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	return checkResult(resp)
}

func sendRequest(method, token, name, path string, params url.Values) ([]byte, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	if path != "" {
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		defer file.Close()
		part, err := writer.CreateFormFile(name, filepath.Base(path))
		if err != nil {
			return nil, err
		}

		if _, err = io.Copy(part, file); err != nil {
			return nil, err
		}
	}
	for field, values := range params {
		if len(values) > 0 {
			writer.WriteField(field, values[0])
		}
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("https://api.telegram.org/bot%s/%s", token, method)
	req, err := http.NewRequest("POST", url, body)
	fmt.Println(body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return checkResult(resp)

}

func checkResult(resp *http.Response) ([]byte, error) {
	if resp.StatusCode != 200 {
		con, _ := ioutil.ReadAll(resp.Body)
		return nil, errors.New(fmt.Sprintf("gotelebot error:%s-%s", resp.Status, con))
	}
	jsonStr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	err = json.Unmarshal(jsonStr, &result)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("gotelebot The server returned an invalid JSON response. %s-%s",
			resp.Status, resp.Body))
	}
	if result["ok"] != true {
		return nil, errors.New(fmt.Sprintf("gotelebot: Error.ErrorCode: %s-Description%s",
			result["errorCode"], result["description"]))
	}
	str, errs := json.Marshal(result["result"])
	if errs != nil {
		fmt.Println("Error encoding JSON")
		return nil, errors.New(fmt.Sprintln("gotelebot"))
	}
	return []byte(str), nil
}

func makeRequest(method, token, name, filepath string, params url.Values) ([]byte, error) {
	return sendRequest(method, token, name, filepath, params)
}

func getMe(token string) (*types.User, error) {
	jsonStr, err := sendGetRequest("getMe", token, url.Values{})
	if err != nil {
		return nil, err
	}
	var user types.User
	err = json.Unmarshal(jsonStr, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func getUpdates(token, offset, limit, timeout string) ([]*types.Update, error) {
	payload := url.Values{}
	if offset != "" {
		payload.Add("offset", offset)
	}
	if limit != "" {
		payload.Add("limit", limit)
	}
	if timeout != "" {
		payload.Add("timeout", timeout)
	}
	jsonStr, err := sendGetRequest("getUpdates", token, payload)
	if err != nil {
		return nil, err
	}
	var result []*types.Update
	err = json.Unmarshal(jsonStr, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func sendMessage(token, chat_id, text, disable_web_page_preview, reply_to_message_id, reply_markup string) (*types.Message, error) {
	payload := url.Values{}
	payload.Add("chat_id", chat_id)
	payload.Add("text", text)
	jsonStr, err := makeRequest("sendMessage", token, "", "", payload)
	var msg types.Message
	err = json.Unmarshal(jsonStr, &msg)
	if err != nil {
		return nil, err
	}
	return &msg, nil
}
