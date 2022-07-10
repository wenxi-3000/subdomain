package dns

import (
	"bufio"
	"log"
	"os"
	"subdomain/dns/api"
	"subdomain/dns/dns"
	"subdomain/libs"
)

func DnsBrute(input string, opt *libs.Options, dnsServer string) {
	domain := input
	server := dnsServer
	log.Println("DNS Server: ", server)
	dict := opt.DictList
	rate := opt.Thread
	retry := 1
	subDomainsToQuery := mixInAPIDict(domain, dict)
	dns.Configure(domain, server, rate, retry)

	// 输入
	go func() {
		for sub := range subDomainsToQuery {
			dns.Queries <- sub
		}
	}()

	for record := range dns.Records {
		opt.BRwmutex.Lock()
		opt.Bresults[record.Domain] = struct{}{}
		opt.BRwmutex.Unlock()
	}
}

func mixInAPIDict(domain, dict string) <-chan string {
	subDomainsToQuery := make(chan string)
	mix := make(chan string)

	// mix in
	go func() {
		defer close(subDomainsToQuery)

		domains := map[string]struct{}{}
		for sub := range mix {
			domains[sub] = struct{}{}
		}

		for domain := range domains {
			subDomainsToQuery <- domain
		}
	}()

	// API
	// 同步，保证 API 完整执行
	for sub := range api.Query(domain) {
		mix <- sub
	}

	go func() {
		defer close(mix)

		// Domain
		mix <- domain

		// Dict
		file, err := os.Open(dict)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			mix <- scanner.Text() + "." + domain
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}()

	return subDomainsToQuery
}
