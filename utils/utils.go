package utils

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/mitchellh/go-homedir"
)

//将文件中的~/xx 替换为绝对路径
func NormalizePath(path string) string {
	if strings.HasPrefix(path, "~") {
		path, _ = homedir.Expand(path)
	}
	return path
}

//判断文件夹是否存在
func FolderExists(foldername string) bool {
	foldername = NormalizePath(foldername)
	if _, err := os.Stat(foldername); os.IsNotExist(err) {
		return false
	}
	return true
}

//判读文件是否存在
func FileExists(filename string) bool {
	filename = NormalizePath(filename)
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

// 创建一个文件夹
func MakeDir(folder string) {
	folder = NormalizePath(folder)
	os.MkdirAll(folder, 0750)
}

//按行读取一个文件，返回[]string
func FileSlice(path string) ([]string, error) {
	var results []string
	if strings.HasPrefix(path, "~") {
		path, _ = homedir.Expand(path)
	}
	file, err := os.Open(path)
	if err != nil {
		return results, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		results = append(results, line)
	}

	if err := scanner.Err(); err != nil {
		return results, err
	}

	return results, nil
}

//按行读取一个文件，返回map[string]{}{}
func FileSet(path string) (map[string]struct{}, error) {
	results := map[string]struct{}{}
	if strings.HasPrefix(path, "~") {
		path, _ = homedir.Expand(path)
	}
	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return results, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		results[line] = struct{}{}
	}

	if err := scanner.Err(); err != nil {
		return results, err
	}

	return results, nil
}

//匹配域名
func GetSubomains(source string, domain string) []string {
	//strings.Replace(domain,'.', '`\.`')
	reg := `(?:[a-z0-9](?:[a-z0-9\-]{0,61}[a-z0-9])?\.){0,}` + domain
	//results_domains = re.findall(regexp, str(source), re.I)
	var linkFinderRegex = regexp.MustCompile(reg)
	matchs := linkFinderRegex.FindAllString(source, -1)
	//fmt.Println(matchs)
	return matchs
}

//匹配域名不包含本身
func GetSubomainsNot(source string, domain string) []string {
	//strings.Replace(domain,'.', '`\.`')
	reg := `(?:[a-z0-9](?:[a-z0-9\-]{0,61}[a-z0-9])?\.){1,}` + domain
	//results_domains = re.findall(regexp, str(source), re.I)
	var linkFinderRegex = regexp.MustCompile(reg)
	matchs := linkFinderRegex.FindAllString(source, -1)
	//fmt.Println(matchs)
	return matchs
}
