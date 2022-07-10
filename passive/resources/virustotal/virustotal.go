package virustotal

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"subdomain/libs"
)

type subdomains struct {
	Resresults []string `json:"subdomains"`
}

func Virustotal(domain string, opt *libs.Options) {
	log.Println("==Virustotal==")
	keys := opt.Keys["virustotal"]
	//随机获取key列表里的值
	rand.Seed(time.Now().Unix())
	key := keys[rand.Intn(len(keys))]
	requestUrl := "https://www.virustotal.com/vtapi/v2/domain/report?apikey=" + key + "&domain=" + domain
	resp, err := http.Get(requestUrl)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))
	var respSubdomain subdomains
	err = json.NewDecoder(resp.Body).Decode(&respSubdomain)
	if err != nil {
		log.Println(err)
	}
	//fmt.Println(respSubdomain.Resresults)
	for _, record := range respSubdomain.Resresults {
		//result := record + "." + domain
		log.Println("virustotal: ", record)
		opt.PRwmutex.Lock()
		opt.Presults[record] = struct{}{}
		opt.PRwmutex.Unlock()
	}

}
