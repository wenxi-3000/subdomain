package securitytrails

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

func Securitytrails(domain string, opt *libs.Options) {
	log.Println("==Securitytrails==")
	keys := opt.Keys["securitytrails"]
	//随机获取key列表里的值
	rand.Seed(time.Now().Unix())
	key := keys[rand.Intn(len(keys))]

	requestUrl := "https://api.securitytrails.com/v1/domain/" + domain + "/subdomains?apikey=" + key
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
		result := record + "." + domain
		log.Println("securitytrails: ", result)
		opt.PRwmutex.Lock()
		opt.Presults[result] = struct{}{}
		opt.PRwmutex.Unlock()
	}

}
