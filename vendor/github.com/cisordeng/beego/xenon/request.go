package xenon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/cisordeng/beego"
)

func Get(service string, resource string, data Map, apiUrls ...string) Map {
	return request("GET", service, resource, data, apiUrls...)
}

func Put(service string, resource string, data Map, apiUrls ...string) Map {
	return request("PUT", service, resource, data, apiUrls...)
}

func Post(service string, resource string, data Map, apiUrls ...string) Map {
	return request("POST", service, resource, data, apiUrls...)
}

func Delete(service string, resource string, data Map, apiUrls ...string) Map {
	return request("DELETE", service, resource, data, apiUrls...)
}


func GetUrl(url string, data Map, header... Map) Map {
	return requestUrl("GET", url, data, header...)
}

func PutUrl(url string, data Map, header... Map) Map {
	return requestUrl("PUT", url, data, header...)
}

func PostUrl(url string, data Map, header... Map) Map {
	return requestUrl("POST", url, data, header...)
}

func DeleteUrl(url string, data Map, header... Map) Map {
	return requestUrl("DELETE", url, data, header...)
}

func requestUrl(method string, requestUrl string, data Map, header... Map) Map {

	params := url.Values{"__source_service": {beego.AppConfig.String("appname")}}
	for k, v := range data {
		value := ""
		switch t := v.(type) {
		case int:
			value = fmt.Sprintf("%d", v)
		case bool:
			value = fmt.Sprintf("%t", v)
		case string:
			value = v.(string)
		case float64:
			value = fmt.Sprintf("%f", v)
		default:
			beego.Notice("json marshal type: ", t)
			bytes, err := json.Marshal(v)
			PanicNotNilError(err)
			value = string(bytes)
		}
		params.Set(k, value)
	}

	requestUrl = fmt.Sprintf("%s?%s", requestUrl, params.Encode())
	beego.Notice(fmt.Sprintf("request url: %s %s", requestUrl, method))

	request, err := http.NewRequest(method, requestUrl, nil)
	PanicNotNilError(err)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.117 Safari/537.36")
	if len(header) > 0 {
		for key, value := range header[0] {
			request.Header.Set(key, value.(string))
		}
	}
	response, err := http.DefaultClient.Do(request)
	PanicNotNilError(err)
	bytes, err := ioutil.ReadAll(response.Body)
	PanicNotNilError(err)
	resMap := Map{}
	err = json.Unmarshal(bytes, &resMap)
	PanicNotNilError(err)
	return resMap
}

func request(method string, service string, resource string, data Map, apiUrls ...string) Map {
	apiUrl := beego.AppConfig.String("api::apiUrl")
	if len(apiUrls) > 0 {
		apiUrl = apiUrls[0]
	}
	if apiUrl[len(apiUrl) - 1] != '/' {
		apiUrl = apiUrl + "/"
	}
	params := url.Values{"__source_service": {beego.AppConfig.String("appname")}}
	for k, v := range data {
		value := ""
		switch t := v.(type) {
		case int:
			value = fmt.Sprintf("%d", v)
		case bool:
			value = fmt.Sprintf("%t", v)
		case string:
			value = v.(string)
		case float64:
			value = fmt.Sprintf("%f", v)
		default:
			beego.Notice("json marshal type: ", t)
			bytes, err := json.Marshal(v)
			PanicNotNilError(err)
			value = string(bytes)
		}
		params.Set(k, value)
	}
	timestamp := time.Now().Unix()
	params.Set("timestamp", fmt.Sprintf("%d", timestamp))
	var signSecret = beego.AppConfig.String("api::signSecret")
	sign := strings.ToLower(EncodeMD5(signSecret + params.Encode()))
	params.Set("sign", sign)

	requestUrl := fmt.Sprintf("%s%s/%s/?%s", apiUrl, service, strings.Replace(resource, ".", "/", -1), params.Encode())
	beego.Notice(fmt.Sprintf("request url: %s %s", requestUrl, method))

	request, err := http.NewRequest(method, requestUrl, nil)
	PanicNotNilError(err)
	request.Header.Set("Content-Type", "application/json")
	response, err := http.DefaultClient.Do(request)
	PanicNotNilError(err)
	bytes, err := ioutil.ReadAll(response.Body)
	PanicNotNilError(err)
	resMap := Map{}
	err = json.Unmarshal(bytes, &resMap)
	PanicNotNilError(err)
	if int(resMap["code"].(float64)) != 200 {
		RaiseException(resMap["errCode"].(string), resMap["errMsg"].(string))
	}
	return resMap
}