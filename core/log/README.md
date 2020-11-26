## 日志配置


#### 配置文件格式
	sfgo:
	  log:
	    # 可选: stdout, stderr, /path/to/log/file
	    output: stdout
	    # 可选: logfmt, json
	    formater: logfmt
	    # 可选: debug, info, warn, error, fatal, panic
	    level: info
	    # 100M
	    # 时间格式
	    timeformat: 2006-01-02T15:04:05.000Z07:00
	    maxsize: 100
	    # 保留备份日志文件数
	    maxbackups: 3
	    # 保留天数
	    maxage: 30
	    # 启动 level server
	    levelserver: false
