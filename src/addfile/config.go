package addfile

	var (
		usingDNS = true
		usingHTTP = true
		launchFakeFile = true  
		usingRegistry  = true  
		usingAutorun  = true                        
		debugMode      = false                         
		attemptCount   = 1
		dnsAddress     = "your-site.com"
		winHttpTimeout = []int{0, 60000, 30000, 30000}
		firstInterface = "HTTP"
		rewriteExe = true

		urls = []string{"127.0.0.1","your-site.com"}

		ports = []int{80,8090}
	)

	func GetUsingDns() bool {
		return usingDNS;
	}

	func GetUsingHttp() bool {
		return usingHTTP;
	}


	func GetRewriteExe() bool{
		return rewriteExe;
	}

	func GetLaunchFakeFile() bool {
		return launchFakeFile;
	}

	func GetDebugMode() bool {
		return debugMode;
	}

	func GetUsingRegistry() bool{
		return usingRegistry;
	}

	func GetUsingAutorun() bool{
		return usingAutorun;
	}

	func GetAttempCount() int{
		return attemptCount;
	}

	func GetDnsAddress() string{
		return dnsAddress;
	}

	func GetWinHttpTimeout() []int{
		return winHttpTimeout;
	}

	func GetServerUrls() []string{
		return urls;
	}

	func GetServerPorts() []int{
		return ports;
	}

	func GetFirstInterface() string{
		return firstInterface
	}		