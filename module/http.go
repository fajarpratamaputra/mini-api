package module

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.elastic.co/apm/module/apmhttp"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func FormatBody(data interface{}) (body *bytes.Buffer, err error) {
	postBody, err_ := json.Marshal(data)
	if err_ != nil {
		err = err_
		return
	}
	body = bytes.NewBuffer(postBody)
	return
}

func RequestAPIJson(ctx context.Context, url string, data interface{}, method string, headers http.Header) (respBody []byte, err error) {
	// APM tracing
	client := &http.Client{}
	body, errorFormatBody := FormatBody(data)
	if errorFormatBody != nil {
		err = errorFormatBody
		return
	}

	req, errNewRequest := http.NewRequest(method, url, body)
	if errNewRequest != nil {
		err = errNewRequest
		return
	}
	req = req.WithContext(ctx)

	req.Header = headers

	client = apmhttp.WrapClient(client)

	resp, errRequest := client.Do(req)
	if errRequest != nil {
		err = errRequest
		return
	}

	// close body after function returns
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	bodyBytes, errReadBody := ioutil.ReadAll(resp.Body)
	if errReadBody != nil {
		err = errReadBody
		return
	}
	respBody = bodyBytes
	return
}

func GetPermalink(ctx context.Context, actionType string, actionValue string, token string) (string, error) {
	actionType = strings.ToUpper(actionType)
	isError := true
	permalink := ""

	if actionType == "" && actionValue == "" {
		return permalink, errors.New("invalid type")
	}

	response := map[string]interface{}{}

	getDataUrl := os.Getenv(fmt.Sprintf("%s_URL", actionType))
	getDataMethod := os.Getenv(fmt.Sprintf("%s_METHOD", actionType))
	linkField := os.Getenv(fmt.Sprintf("%s_FIELD", actionType))

	if linkField == "" {
		linkField = "permalink"
	}
	if getDataMethod == "" {
		getDataMethod = "GET"
	}

	if getDataUrl != "" {
		getDataUrl = strings.ReplaceAll(getDataUrl, ":id", actionValue)
		headers := http.Header{
			"Authorization": []string{token},
		}
		if respBody, err := RequestAPIJson(ctx, getDataUrl, nil, getDataMethod, headers); err == nil {
			if err_ := json.Unmarshal(respBody, &response); err_ == nil {
				data := response["data"]
				if indexedData, ok := data.(map[string]interface{}); ok {
					if val, ok_ := indexedData[linkField]; ok_ {
						if valStr, convertOk := val.(string); convertOk {
							permalink = valStr
							isError = false
						}
					}
				}
			}
		}
	}

	if isError {
		return "", errors.New("something went wrong")
	}

	return permalink, nil
}
