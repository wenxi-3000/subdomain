package libs

import (
	"log"
	"path"
	"runtime"
	"subdomain/utils"
	"sync"
)

//全局配置
type Options struct {
	//定义["passive", "brute"]
	Module     []string
	Domain     string
	Domains    []string
	Inputs     map[string]struct{}
	InputFile  string
	ConfigFile string
	DictList   string
	Retry      int
	Thread     int
	Paths      Paths
	DnsServer  string
	Source     []string
	Keys       map[string][]string
	PRwmutex   sync.RWMutex
	BRwmutex   sync.RWMutex
	Presults   map[string]struct{}
	Bresults   map[string]struct{}
}

//路径配置
type Paths struct {
	//结果目录
	Result string

	//当前项目根路径
	Root string
}

func InitOptions(opt *Options) *Options {
	//设置日志格式
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	//初始化输入
	opt.Inputs = initInput(opt)
	//初始化路径
	paths := initPath()
	opt.Paths = paths

	//初始化配置文件
	err := InitConfig(opt)
	if err != nil {
		log.Println(err)
	}

	return opt
}

func initInput(options *Options) map[string]struct{} {
	var inputs map[string]struct{}
	if options.InputFile != "" {
		var err error
		if inputs, err = utils.FileSet(options.InputFile); err != nil {
			log.Println(err)
		}
	}

	if options.Domain != "" {
		inputs[options.Domain] = struct{}{}
	}
	if len(options.Domains) != 0 {
		for _, domain := range options.Domains {
			inputs[domain] = struct{}{}
		}
	}
	return inputs
}

func initPath() Paths {
	var paths Paths
	//当前文件路径
	_, currentFile, _, _ := runtime.Caller(0)
	//当前文件目录
	currentPath := path.Dir(currentFile)
	rootPath := path.Dir(currentPath)
	paths.Root = rootPath

	//结果文件路径
	paths.Result = path.Join(rootPath, "results")
	//不存在结果文件夹就创建
	if !utils.FolderExists(paths.Result) {
		utils.MakeDir(paths.Result)
	}
	return paths
}
