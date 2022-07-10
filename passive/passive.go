package passive

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"subdomain/libs"
	"subdomain/passive/resources/censys"
	"subdomain/passive/resources/fofa"
	"subdomain/passive/resources/qianxun"
	"subdomain/passive/resources/securitytrails"
	"subdomain/passive/resources/virustotal"
	"subdomain/utils"
	"sync"
)

func Task(domain string, opt *libs.Options, wg *sync.WaitGroup) {
	for _, source := range opt.Source {
		log.Println("source: ", source)
		wg.Add(1)
		go doTask(source, domain, opt, wg)
	}
	wg.Wait()
}

func Passive(options *libs.Options) {
	for input := range options.Inputs {
		var wg sync.WaitGroup
		//开始收集input的子域名
		options.Presults = make(map[string]struct{})
		log.Println("开始被动收集域名: ", input)
		Task(input, options, &wg)

		//判断结果文件是否存在，存在就删除
		reutltFile := path.Join(options.Paths.Result, input+"-passive"+".txt")
		if utils.FileExists(reutltFile) {
			os.Remove(utils.NormalizePath(reutltFile))
		}
		//结果处理
		file, err := os.OpenFile(reutltFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		defer file.Close()
		w := bufio.NewWriter(file)
		for domain := range options.Presults {
			fmt.Println(domain)
			w.Write([]byte(domain + "\n"))
		}
		w.Flush()
	}
	// tasks := []string{
	// 	// "alienvault",
	// 	"fofa",
	// 	// "sensys",
	// 	// "qianxun",
	// 	// "securitytrails",
	// 	// "virustotal",
	// }

	// //并发查询
	// for _, task := range tasks {
	// 	wg.Add(1)
	// 	go doTask(task, options)
	// }
	// wg.Wait()

	// // //读取结果文件，去重，清洗
	// // results := GetPassiveResult(options.Domain, options.TmpPath)
	// // for _, result := range results {
	// // 	fmt.Println(result)
	// // }

	// // //保存结果
	// // utils.SaveResult(results, "passive_sudomain.txt", options.JobPath)
	// // log.Println("被动收集的域名已完成,结果保存在: ", options.JobPath)

	// os.Exit(0)
}

func doTask(task string, domain string, opt *libs.Options, wg *sync.WaitGroup) {
	switch {
	// case task == "alienvault":
	// 	alienvault.Alienvault(options)
	// 	wg.Done()

	case task == "fofa":
		fofa.Fofa(domain, opt)
		wg.Done()

	case task == "censys":
		censys.Censys(domain, opt)
		wg.Done()

	case task == "qianxun":
		qianxun.Qianxun(domain, opt)
		wg.Done()

	case task == "securitytrails":
		securitytrails.Securitytrails(domain, opt)
		wg.Done()

	case task == "virustotal":
		virustotal.Virustotal(domain, opt)
		wg.Done()
	}
	// }
}
