package censys

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"subdomain/libs"

	"github.com/corpix/uarand"
)

type resultsq struct {
	Data  []string `json:"parsed.extensions.subject_alt_name.dns_names"`
	Data1 []string `json:"parsed.names"`
}

type response struct {
	Results  []resultsq `json:"results"`
	Metadata struct {
		Pages int `json:"pages"`
	} `json:"metadata"`
}

type BasicAuth struct {
	Username string
	Password string
}

func Censys(domain string, opt *libs.Options) {
	log.Println("==Censys==")
	keys := opt.Keys["censys"]
	//随机获取key列表里的值
	rand.Seed(time.Now().Unix())
	key := keys[rand.Intn(len(keys))]
	keyx := strings.Split(key, ":")
	UID := keyx[0]
	SECRET := keyx[1]
	delay := 3
	maxPage := 10
	page := 1
	var censysResponse response
	for {
		var data = []byte(`{"query":"` + domain + `", "page":` + strconv.Itoa(page) + `, "fields":["parsed.names","parsed.extensions.subject_alt_name.dns_names"], "flatten":true}`)
		time.Sleep(time.Second * time.Duration(delay))
		resp, err := HTTPRequest(
			"POST",
			"https://search.censys.io/api/v1/search/certificates",
			"",
			map[string]string{"Content-Type": "application/json", "Accept": "application/json"},
			bytes.NewReader(data),
			BasicAuth{Username: UID, Password: SECRET},
		)
		if err != nil {

			log.Println(err)
			break
		}

		if resp.Status == "200 OK" {
			// body, err := ioutil.ReadAll(resp.Body)
			// if err != nil {
			// 	log.Println(err)
			// }
			// respbody := string(body)
			// //fmt.Println(respbody)
			// result := utils.GetSubomainsNot(respbody, domain)
			// //fmt.Println(result)
			// if result == nil {
			// 	break
			// }

			err = json.NewDecoder(resp.Body).Decode(&censysResponse)
			if err != nil {
				log.Println(err)
			}

			//fmt.Println(censysResponse.Results)
			// fmt.Println("page: ", censysResponse.Metadata.Page)
			if len(censysResponse.Results) == 0 {
				//fmt.Println(len(censysResponse.Results))
				break
			}
			// // strconv.
			for _, record := range censysResponse.Results {
				for _, i := range record.Data {
					log.Println("censys: ", i)
					opt.PRwmutex.Lock()
					opt.Presults[i] = struct{}{}
					opt.PRwmutex.Unlock()
				}

				// results = append(results, record.Data[i])
				// utils.SaveTmp(results, "censys_domain.txt")
			}
			// utils.SaveTmp(results, "censys_domain.txt")
			//fmt.Println(json.NewDecoder(resp.Body))

			//fmt.Println(len(response.Results))
			resp.Body.Close()
			//fmt.Println(page)
			if page > maxPage {
				break
			}
			page++
		} else {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println(err)
			}
			log.Println(string(body))
			break
		}

	}
}

func HTTPRequest(method, requestURL, cookies string, headers map[string]string, body io.Reader, basicAuth BasicAuth) (*http.Response, error) {
	req, err := http.NewRequest(method, requestURL, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", uarand.GetRandom())
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en")
	req.Header.Set("Connection", "close")

	if basicAuth.Username != "" || basicAuth.Password != "" {
		req.SetBasicAuth(basicAuth.Username, basicAuth.Password)
	}

	if cookies != "" {
		req.Header.Set("Cookie", cookies)
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// proxy, _ := url.Parse("http://127.0.0.1:8080")
	// tr := &http.Transport{
	// 	Proxy:           http.ProxyURL(proxy),
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// }

	client := &http.Client{
		Timeout: time.Second * 10, //超时时间
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
