package brute

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"subdomain/dns"
	"subdomain/libs"
	"subdomain/utils"
	"time"
)

func Brute(options *libs.Options) {
	for input := range options.Inputs {
		startBrute(input, options)
	}
}

//爆破一个域名的子域名
func startBrute(input string, opt *libs.Options) {
	log.Println("开始爆破：", input)
	start := time.Now()

	opt.Bresults = make(map[string]struct{})
	//每个域名跑opt.Retry次，每跑一次修改一次DNS Server
	if opt.DnsServer == "" {
		for i, j := 0, 0; i < opt.Retry; i, j = i+1, j+1 {
			if j >= len(libs.DnsServers) {
				j = 0
			}
			dnsServer := libs.DnsServers[j]
			dns.DnsBrute(input, opt, dnsServer)
		}
	} else {
		//如果通过命令行指定了DNS Server则直接使用
		for i := 0; i < opt.Retry; i++ {
			dnsServer := opt.DnsServer
			dns.DnsBrute(input, opt, dnsServer)
		}
	}

	//输出
	//判断结果文件是否存在，存在就删除
	resultFile := path.Join(opt.Paths.Result, input+"-brute"+".txt")
	log.Println("dns爆破结果保存在: ", resultFile)
	if utils.FileExists(resultFile) {
		os.Remove(utils.NormalizePath(resultFile))
	}
	file, err := os.OpenFile(resultFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	w := bufio.NewWriter(file)
	counter := 0
	log.Println("brute: ")
	for item := range opt.Bresults {
		counter++
		fmt.Println(item)
		w.Write([]byte(item + "\n"))
	}
	w.Flush()

	log.Println(time.Since(start).Seconds())
	log.Println("records: ", counter)
}
