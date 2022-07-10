package fofa

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"subdomain/libs"
	"subdomain/utils"
	"time"
)

type fofaResponse struct {
	Results []string `json:"results"`
}

func Fofa(domain string, opt *libs.Options) {
	log.Println("==Fofa==")
	keys := opt.Keys["fofa"]
	rand.Seed(time.Now().Unix())
	key := keys[rand.Intn(len(keys))]
	keyx := strings.Split(key, ":")
	user := keyx[0]
	token := keyx[1]
	page := 1
	qbase64 := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("domain=\"%s\"", domain)))

	//fmt.Println(url)

	var response fofaResponse
	for {
		//最多只能查看10000条数据，如需获取 更多数据，请在搜索界面点击"下载数据"链接进行下载。
		if page >= 1001 {
			break
		}

		url := "https://fofa.info/api/v1/search/all?full=true&fields=host&page=" + strconv.Itoa(page) + "&size=10&email=" + user + "&key=" + token + "&qbase64=" + qbase64
		// fmt.Println(url)
		resp, err := http.Get(url)
		if err != nil {
			log.Println(err)
		}
		// body, err := ioutil.ReadAll(resp.Body)
		// fmt.Println(string(body))
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			log.Println(err)
		}
		//fmt.Println(response.Results)
		if len(response.Results) == 0 {
			break
		}

		for _, records := range response.Results {
			lines := utils.GetSubomains(records, domain)
			for _, line := range lines {
				log.Println("fofa: ", line)
				opt.PRwmutex.Lock()
				opt.Presults[line] = struct{}{}
				opt.PRwmutex.Unlock()
			}

		}

		resp.Body.Close()
		//fmt.Println("page: ", page)
		page++
		time.Sleep(1 * time.Second)

	}

}
