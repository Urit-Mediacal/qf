package launcher

import (
	"fmt"
	"github.com/kardianos/service"
	"log"
	"os"
	"path"
	"path/filepath"
	"qf/util/io"
	"runtime"
	"strings"
	"sync"
	"time"
)

//func main() {
//	// 启动
//	launcher.Run(start, stop)
//}
//
//func start() {
//	// 进行自己项目的相关初始化
//	// 如初始化api，初始化web服务等
//}
//
//func stop() {
//	// 进行自己项目的相关释放
//	// 如停止web服务
//}

var (
	stopChan = make(chan struct{}, 1)
	wg       = &sync.WaitGroup{}
)

// Run 运行服务
func Run(start func(), stop func()) {
	// 初始化系统服务
	pm := &program{
		start: start,
		stop:  stop,
	}
	// 获取当前程序所在路径
	p, err := io.GetCurrentDirectory()
	p = strings.Replace(p, "\\", "/", -1)
	n1 := strings.Split(path.Dir(p), "/")
	n2 := strings.TrimSuffix(path.Base(p), path.Ext(p))
	serv, err := service.New(pm, &service.Config{
		// 统一使用 目录名_文件名 作为服务名
		Name: n1[len(n1)-1] + "_" + n2,
	})
	if err != nil {
		log.Fatalln(err)
		return
	}
	// 修改当前工作目录为exe所在目录
	// 如果不执行该操作，注册生成服务后，程序路径会默认在系统盘
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	err = os.Chdir(dir)
	if err != nil {
		log.Fatalln(err)
		return
	}

	// 如果是linux系统且未安装改服务时
	if runtime.GOOS == "linux" {
		st, se := serv.Status()
		if st == service.StatusUnknown && se.Error() == "the service is not installed" {
			// 如果有对应的服务部署文件
			if io.PathExists(fmt.Sprintf("/lib/systemd/system/%s.service", serv.String())) {
				err = serv.Install()
				if err != nil {
					log.Println(fmt.Sprintf("[%s] Installed Error, %s", serv.String(), err))
				} else {
					log.Println(fmt.Sprintf("[%s] Installed OK", serv.String()))
				}
			}
		}
	}
	// 运行
	err = serv.Run()
	if err != nil {
		log.Println(err)
		return
	}
}

// Exit 退出服务
func Exit() {
	go func() {
		time.Sleep(time.Millisecond * 100)
		close(stopChan)
		wg.Wait()
		os.Exit(0)
	}()
}

type program struct {
	start func()
	stop  func()
}

func (p *program) Start(s service.Service) error {
	// 执行外层启动
	if p.start != nil {
		p.start()
	}

	// 启动成功
	log.Println(fmt.Sprintf("[%s] Started OK", s.String()))

	wg.Add(1)
	stopChan = make(chan struct{})
	go p.run()

	return nil
}

func (p *program) run() {
	for {
		select {
		case <-stopChan:
			// 执行外层停止
			if p.stop != nil {
				p.stop()
			}
			// 全部退出完成
			wg.Done()
			return
		}
	}
}

func (p *program) Stop(s service.Service) error {
	// 启动成功
	log.Println(fmt.Sprintf("[%s] Stoped OK", s.String()))
	close(stopChan)
	wg.Wait()
	return nil
}
