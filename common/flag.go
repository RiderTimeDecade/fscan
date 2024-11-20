package common

import (
	"flag"
	"fmt"
)

func Banner() {
	banner := `
    ______                        
   / ____/______________ ___  ___ 
  / /_  / ___/ ___/ __ '__ \/ _ \
 / __/ (__  ) /__/ / / / / /  __/
/_/   /____/\___/_/ /_/ /_/\___/ 
                                
    A fast all-in-one intranet scanning tool.
    
    Version: %s
    Author:  shadow1ng
    Github:  https://github.com/shadow1ng/fscan

    Use -h for help, or visit: https://github.com/shadow1ng/fscan#readme
    
`
	if !Nocolor {
		banner = fmt.Sprintf("\033[36m%s\033[0m", banner) // 使用青色
	}
	fmt.Printf(banner, version)
}

func Flag(Info *HostInfo) {
	Banner()

	// 将参数分组，使其更有条理
	// 基本选项
	flag.StringVar(&Info.Host, "h", "", "Target host(s). Supports: IP, CIDR, Range.\n\tExample: 192.168.1.1, 192.168.1.1/24, 192.168.1.1-255")
	flag.StringVar(&NoHosts, "hn", "", "Hosts to exclude from scan\n\tExample: -hn 192.168.1.1/24")
	flag.StringVar(&Ports, "p", DefaultPorts, "Target port(s). Supports: Port, Range, List\n\tExample: 22, 1-65535, 22,80,3306")
	flag.StringVar(&PortAdd, "pa", "", "Add additional ports to default port list\n\tExample: -pa 3389")

	// 认证选项
	flag.StringVar(&Username, "user", "", "Username for authentication")
	flag.StringVar(&Password, "pwd", "", "Password for authentication")
	flag.StringVar(&UserAdd, "usera", "", "Add additional username to default list")
	flag.StringVar(&PassAdd, "pwda", "", "Add additional password to default list")
	flag.StringVar(&Domain, "domain", "", "Domain for SMB authentication")
	flag.StringVar(&Hash, "hash", "", "Hash for authentication")

	// 扫描控制
	flag.StringVar(&Scantype, "m", "all", "Scan module selection\n\tExample: -m ssh,smb,web")
	flag.IntVar(&Threads, "t", 600, "Number of concurrent threads")
	flag.Int64Var(&Timeout, "time", 3, "Timeout in seconds")
	flag.Int64Var(&WebTimeout, "wt", 5, "Web timeout in seconds")
	flag.BoolVar(&NoPing, "np", false, "Skip ping detection")
	flag.BoolVar(&Ping, "ping", false, "Use ping instead of ICMP")

	// 文件操作
	flag.StringVar(&HostFile, "hf", "", "Import targets from file")
	flag.StringVar(&Userfile, "userf", "", "Import usernames from file")
	flag.StringVar(&Passfile, "pwdf", "", "Import passwords from file")
	flag.StringVar(&Outputfile, "o", "result.txt", "Output file path")

	// 代理设置
	flag.StringVar(&Proxy, "proxy", "", "HTTP proxy for POC scanning\n\tExample: http://127.0.0.1:8080")
	flag.StringVar(&Socks5Proxy, "socks5", "", "Socks5 proxy for TCP connections")

	// 特殊功能
	flag.StringVar(&Command, "c", "", "Command to execute (ssh|wmiexec)")
	flag.StringVar(&RedisFile, "rf", "", "Redis public key file\n\tExample: -rf id_rsa.pub")
	flag.StringVar(&RedisShell, "rs", "", "Redis shell command\n\tExample: -rs 192.168.1.1:6666")

	// 输出控制
	flag.BoolVar(&Silent, "silent", false, "Silent mode, no banner")
	flag.BoolVar(&Nocolor, "nocolor", false, "Disable colored output")
	flag.BoolVar(&JsonOutput, "json", false, "Output in JSON format")

	// POC相关
	flag.BoolVar(&NoPoc, "nopoc", false, "Skip web vulnerability scanning")
	flag.StringVar(&PocPath, "pocpath", "", "Custom POC file path")
	flag.StringVar(&Pocinfo.PocName, "pocname", "", "Specify POC by name\n\tExample: -pocname weblogic")

	flag.Parse()
}
