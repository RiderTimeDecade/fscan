package common

var version = "1.8.4"
var Userdict = map[string][]string{
	"ftp":        {"ftp", "anonymous", "admin", "www", "web", "root", "db", "wwwroot", "data", "ftpuser", "upload", "test", "public", "backup", "webadmin", "developer", "www-data", "ftpadmin", "ftpbackup", "webmaster"},
	"mysql":      {"root", "mysql", "admin", "dba", "webapp", "dev", "test", "backup", "dbadmin", "mysqlroot", "mysqluser", "mysqladmin", "system", "database", "webdata", "webdb", "datauser"},
	"mssql":      {"sa", "sql", "admin", "sqlserver", "sqldb", "dbadmin", "mssqladmin", "mssqluser", "webdb", "backup", "test", "dev", "dba", "system", "administrator", "sqlservice", "mssql", "sqlroot"},
	"smb":        {"administrator", "admin", "guest", "root", "system", "backup", "webadmin", "user", "test", "share", "public", "everyone", "domain", "domainadmin", "administrator.domain", "admin.domain", "service", "smbuser", "sharepoint", "network", "office"},
	"rdp":        {"administrator", "admin", "guest", "remote", "support", "helpdesk", "system", "user", "terminal", "rdpuser", "desktop", "operator", "technician", "service", "remote.admin", "remote.support", "administrator.domain", "admin.domain", "domain.admin"},
	"postgresql": {"postgres", "admin", "postgresql", "pgadmin", "pgsql", "postgresadmin", "pguser", "dbadmin", "dba", "system", "webapp", "webdb", "backup", "dev", "test", "postgres.admin", "postgresql.user"},
	"ssh":        {"root", "admin", "system", "user", "shell", "ssh", "sshuser", "remote", "ubuntu", "centos", "debian", "redhat", "oracle", "linux", "unix", "webmaster", "deploy", "devops", "operation", "administrator", "ec2-user", "azureuser", "backup", "git"},
	"mongodb":    {"root", "admin", "mongodb", "mongoadmin", "mongouser", "dbadmin", "system", "backup", "dba", "webapp", "webdb", "dev", "test", "mongo", "mongod", "mongodb.admin", "mongodb.user"},
	"oracle":     {"sys", "system", "admin", "test", "web", "orcl", "oracle", "sysdba", "dba", "scott", "dbsnmp", "backup", "orabackup", "weblogic", "middleware", "oradba", "oraoper", "sysman", "oracle.admin", "oracle.sys", "oracle.system"},
	"redis":      {"redis", "admin", "root", "system", "default", "redisuser", "redisadmin", "cache", "backup", "master", "slave"},
	"web":        {"admin", "administrator", "webadmin", "root", "system", "test", "guest", "user", "web", "developer", "manager", "supervisor", "operation", "ops", "devops", "security", "audit", "monitor", "support", "helpdesk"},
}

var Passwords = []string{
	// 常见简单密码
	"1", "123", "123456", "123123", "123321", "12345678", "123456789", "1234567890",
	"666666", "888888", "000000", "111111", "11111111",

	// 常见用户名相关
	"{user}", "{user}123", "{user}123..", "{user}@123", "{user}#123", "{user}_123",
	"{user}123!", "{user}@123#", "{user}@2023", "{user}@2024",
	"{user}1234", "{user}12345", "{user}123456",

	// 管理员相关
	"admin", "admin123", "admin@123", "admin123!@#", "Admin@123",
	"administrator", "administrator123", "superadmin",
	"sysadmin", "system", "system123",

	// 服务相关
	"root", "root123", "root@123", "toor",
	"mysql", "mysql123", "mysql@123",
	"oracle", "oracle123", "oracle@123",
	"sqlserver", "sa123456", "sa@123",
	"postgres", "postgres123",
	"mongodb", "mongo123",
	"redis123", "redis@123",

	// 常见弱密码组合
	"password", "Password", "password123", "Password123",
	"p@ssw0rd", "P@ssw0rd", "P@ssword", "P@ssword123",
	"Passw0rd", "Passw0rd!", "P@ssw0rd!", "P@$$w0rd",

	// 键盘组合
	"1qaz2wsx", "1qaz@WSX", "1qaz!QAZ",
	"qwerty", "qwerty123", "qwertyuiop",
	"!QAZ2wsx", "!QAZ@WSX", "2wsx@WSX",
	"zxcvbnm", "asdfghjkl",

	// 常见英文+数字组合
	"abc123", "abc123456", "abc@123",
	"test", "test123", "test@123",
	"guest", "guest123", "guest@123",

	// 复杂密码组合
	"Admin@2023", "Admin@2024",
	"P@ssw0rd2023", "P@ssw0rd2024",
	"Qwerty@123", "Qwerty@2023",
	"Welcome@123", "Welcome@2023",
	"Complex@123", "Complex@2023",

	// 特殊字符组合
	"P@ssw0rd123!@#",
	"Admin123!@#$",
	"Root123!@#",
	"Test123!@#",

	// 中文拼音常见组合
	"woaini", "woaini123",
	"iloveyou", "iloveyou123",
	"aini123", "aini1314",

	// 日期相关
	"123456789a",
	"12345678a",
	"1234567a",
	"123456a",
	"12345a",
	"1234a",

	// 空密码
	"",

	// 服务默认密码
	"changeme",
	"changeme123",
	"default",
	"default123",
	"admin888",
	"admin666",
	"manager",
	"manager123",

	// 常见变种
	"Admin123456",
	"admin123456",
	"root123456",
	"123456root",
	"123456admin",
	"administrator123456",

	// 特殊服务密码
	"tomcat", "tomcat123",
	"weblogic", "weblogic123",
	"websphere", "websphere123",
	"jboss", "jboss123",

	// 云服务默认密码
	"Aliyun@123",
	"Tencent@123",
	"Huawei@123",
	"Cloud@123",

	// 季节性密码
	"Spring@2023", "Spring@2024",
	"Summer@2023", "Summer@2024",
	"Autumn@2023", "Autumn@2024",
	"Winter@2023", "Winter@2024",
}
var PORTList = map[string]int{
	"ftp":         21,
	"ssh":         22,
	"findnet":     135,
	"netbios":     139,
	"smb":         445,
	"mssql":       1433,
	"oracle":      1521,
	"mysql":       3306,
	"rdp":         3389,
	"psql":        5432,
	"redis":       6379,
	"fcgi":        9000,
	"mem":         11211,
	"mgo":         27017,
	"ms17010":     1000001,
	"cve20200796": 1000002,
	"web":         1000003,
	"webonly":     1000003,
	"webpoc":      1000003,
	"smb2":        1000004,
	"wmiexec":     1000005,
	"all":         0,
	"portscan":    0,
	"icmp":        0,
	"main":        0,
}
var PortGroup = map[string]string{
	"ftp":         "21",
	"ssh":         "22",
	"findnet":     "135",
	"netbios":     "139",
	"smb":         "445",
	"mssql":       "1433",
	"oracle":      "1521",
	"mysql":       "3306",
	"rdp":         "3389",
	"psql":        "5432",
	"redis":       "6379",
	"fcgi":        "9000",
	"mem":         "11211",
	"mgo":         "27017",
	"ms17010":     "445",
	"cve20200796": "445",
	"service":     "21,22,135,139,445,1433,1521,3306,3389,5432,6379,9000,11211,27017",
	"db":          "1433,1521,3306,5432,6379,11211,27017",
	"web":         "80,81,82,83,84,85,86,87,88,89,90,91,92,98,99,443,800,801,808,880,888,889,1000,1010,1080,1081,1082,1099,1118,1888,2008,2020,2100,2375,2379,3000,3008,3128,3505,5555,6080,6648,6868,7000,7001,7002,7003,7004,7005,7007,7008,7070,7071,7074,7078,7080,7088,7200,7680,7687,7688,7777,7890,8000,8001,8002,8003,8004,8006,8008,8009,8010,8011,8012,8016,8018,8020,8028,8030,8038,8042,8044,8046,8048,8053,8060,8069,8070,8080,8081,8082,8083,8084,8085,8086,8087,8088,8089,8090,8091,8092,8093,8094,8095,8096,8097,8098,8099,8100,8101,8108,8118,8161,8172,8180,8181,8200,8222,8244,8258,8280,8288,8300,8360,8443,8448,8484,8800,8834,8838,8848,8858,8868,8879,8880,8881,8888,8899,8983,8989,9000,9001,9002,9008,9010,9043,9060,9080,9081,9082,9083,9084,9085,9086,9087,9088,9089,9090,9091,9092,9093,9094,9095,9096,9097,9098,9099,9100,9200,9443,9448,9800,9981,9986,9988,9998,9999,10000,10001,10002,10004,10008,10010,10250,12018,12443,14000,16080,18000,18001,18002,18004,18008,18080,18082,18088,18090,18098,19001,20000,20720,21000,21501,21502,28018,20880",
	"all":         "1-65535",
	"main":        "21,22,80,81,135,139,443,445,1433,1521,3306,5432,6379,7001,8000,8080,8089,9000,9200,11211,27017",
}
var Outputfile = "result.txt"
var IsSave = true
var Webport = "80,81,82,83,84,85,86,87,88,89,90,91,92,98,99,443,800,801,808,880,888,889,1000,1010,1080,1081,1082,1099,1118,1888,2008,2020,2100,2375,2379,3000,3008,3128,3505,5555,6080,6648,6868,7000,7001,7002,7003,7004,7005,7007,7008,7070,7071,7074,7078,7080,7088,7200,7680,7687,7688,7777,7890,8000,8001,8002,8003,8004,8006,8008,8009,8010,8011,8012,8016,8018,8020,8028,8030,8038,8042,8044,8046,8048,8053,8060,8069,8070,8080,8081,8082,8083,8084,8085,8086,8087,8088,8089,8090,8091,8092,8093,8094,8095,8096,8097,8098,8099,8100,8101,8108,8118,8161,8172,8180,8181,8200,8222,8244,8258,8280,8288,8300,8360,8443,8448,8484,8800,8834,8838,8848,8858,8868,8879,8880,8881,8888,8899,8983,8989,9000,9001,9002,9008,9010,9043,9060,9080,9081,9082,9083,9084,9085,9086,9087,9088,9089,9090,9091,9092,9093,9094,9095,9096,9097,9098,9099,9100,9200,9443,9448,9800,9981,9986,9988,9998,9999,10000,10001,10002,10004,10008,10010,10250,12018,12443,14000,16080,18000,18001,18002,18004,18008,18080,18082,18088,18090,18098,19001,20000,20720,21000,21501,21502,28018,20880"
var DefaultPorts = "21,22,80,81,135,139,443,445,1433,1521,3306,5432,6379,7001,8000,8080,8089,9000,9200,11211,27017"

type HostInfo struct {
	Host    string
	Ports   string
	Url     string
	Infostr []string
}

type PocInfo struct {
	Target  string
	PocName string
}

var (
	Ports       string
	Path        string
	Scantype    string
	Command     string
	SshKey      string
	Domain      string
	Username    string
	Password    string
	Proxy       string
	Timeout     int64 = 3
	WebTimeout  int64 = 5
	TmpSave     bool
	NoPing      bool
	Ping        bool
	Pocinfo     PocInfo
	NoPoc       bool
	IsBrute     bool
	RedisFile   string
	RedisShell  string
	Userfile    string
	Passfile    string
	Hashfile    string
	HostFile    string
	PortFile    string
	PocPath     string
	Threads     int
	URL         string
	UrlFile     string
	Urls        []string
	NoPorts     string
	NoHosts     string
	SC          string
	PortAdd     string
	UserAdd     string
	PassAdd     string
	BruteThread int
	LiveTop     int
	Socks5Proxy string
	Hash        string
	Hashs       []string
	HashBytes   [][]byte
	HostPort    []string
	IsWmi       bool
	Noredistest bool
)

var (
	UserAgent  = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36"
	Accept     = "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
	DnsLog     bool
	PocNum     int
	PocFull    bool
	CeyeDomain string
	ApiKey     string
	Cookie     string
)
