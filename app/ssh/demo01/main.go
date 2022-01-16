package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"learn-go/common/funcs"
	"log"
	"os"
	"os/signal"
	"path"
	"sync"
	"time"

	"github.com/mitchellh/go-homedir"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type sshInfo struct {
	Host    string
	User    string
	passwd  string
	Type    string // password or key
	keyPath string
	port    int
	Client  *ssh.Client
	Session *ssh.Session
}

var sshList map[string]*sshInfo
var wg sync.WaitGroup
var timeOut = 60

func init() {
	sshList = make(map[string]*sshInfo)
	sshList = map[string]*sshInfo{
		"ys": {
			Host:    "193.8.83.217",
			User:    "root",
			passwd:  "Lewaimai_123",
			Type:    "password",
			keyPath: "",
			port:    22,
		},
		"qf_h": {
			Host:    "43.155.76.164",
			User:    "root",
			passwd:  "Lewaimai_123",
			Type:    "password",
			keyPath: "",
			port:    22,
		},
		"qf_a": {
			Host:    "43.130.117.95",
			User:    "root",
			passwd:  "Lewaimai_123",
			Type:    "password",
			keyPath: "",
			port:    22,
		},
		"90_h": {
			Host:    "27.124.40.40",
			User:    "root",
			passwd:  "Lewaimai_123",
			Type:    "password",
			keyPath: "",
			port:    22,
		},
		// "90_a": {
		// 	Host:    "154.38.228.74",
		// 	User:    "root",
		// 	passwd:  "Lewaimai_123",
		// 	Type:    "password",
		// 	keyPath: "",
		// 	port:    22,
		// },
	}
}

func (sInfo *sshInfo) getClient() (*ssh.Client, error) {
	//创建sshp登陆配置
	config := &ssh.ClientConfig{
		Timeout:         time.Second * 3, //ssh 连接time out 时间一秒钟, 如果ssh验证错误 会在一秒内返回
		User:            sInfo.User,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //这个可以， 但是不够安全
		//HostKeyCallback: hostKeyCallBackFunc(h.Host),
	}
	if sInfo.Type == "password" {
		config.Auth = []ssh.AuthMethod{ssh.Password(sInfo.passwd)}
	} else {
		config.Auth = []ssh.AuthMethod{publicKeyAuthFunc(sInfo.keyPath)}
	}

	//dial 获取ssh client
	addr := fmt.Sprintf("%s:%d", sInfo.Host, sInfo.port)
	sshClient, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		return nil, err
	}
	return sshClient, nil
}

func main() {
	// 监听控制台退出信号
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// 执行操作
	go func() {

		defer func() {
			fmt.Println("exit .........")
		}()

		// 建立ssh连接
		for name, sInfo := range sshList {
			var err error
			sInfo.Client, err = sInfo.getClient()
			if err != nil {
				log.Printf("%s create client error: %s \n", name, err)
			}
			log.Printf("%s create client success .... \n", name)
			defer sInfo.Client.Close()
		}

		LastTime := time.Now()

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("请输入远程执行命令: ")
		for scanner.Scan() {
			cmd := scanner.Text()
			if cmd == "quit" {
				fmt.Println("Quitting...")
				break
			}

			// 判断是否超过了最大链接时常
			if time.Since(LastTime).Seconds() > float64(timeOut) {
				log.Fatalln("连接超时， 长时间未输入命令...")
			}

			for name, sInfo := range sshList {
				wg.Add(1)
				go func(sInfo *sshInfo, name string) {
					defer wg.Done()
					// UploadFile(sInfo, "D:/dev.txt", "/root")

					switch cmd {
					case "upload":
						if name == "ys" {
							return
						}
						if name == "90_h" || name == "90_a" {
							doCmd(sInfo, name, `rm -rf /home/wwwroot/default/16.zip`)
							doCmd(sInfo, name, `rm -rf /home/wwwroot/default/renren.zip`)
							UploadFile(sInfo, "D:/app/16.zip", "/home/wwwroot/default")
							cmd = "ls -a /home/wwwroot/default"
							doCmd(sInfo, name, cmd)
						}
						if name == "qf_h" || name == "qf_a" {
							doCmd(sInfo, name, `rm -rf /home/wwwroot/default/69.zip`)
							doCmd(sInfo, name, `rm -rf /home/wwwroot/default/renren.zip`)
							UploadFile(sInfo, "D:/app/69.zip", "/home/wwwroot/default")
							cmd = "ls -a /home/wwwroot/default"
							doCmd(sInfo, name, cmd)
						}

					case "test":
						// if name != "ys" {
						// 	sInfo.Session.Output("rm -rf /home/wwwroot/default/renren/.env.php")
						// 	UploadFile(sInfo, "D:/app/.env.php", "/home/wwwroot/default/renren")
						// }

						// if name != "ys" && name != "qf_h" {
						// 	//执行远程命令
						// 	cmd = `mysql -uroot -plewaimai123 && create user "admin"@"%" identified by "lewaimai123"; && grant all privileges on *.* to 'admin'@'%'; && FLUSH PRIVILEGES; && exit;`
						// 	combo, err := sInfo.Session.CombinedOutput(cmd)
						// 	if err != nil {
						// 		log.Printf("%s 远程执行 (%s) 失败: %s \n", name, cmd, err)
						// 		return
						// 	}
						// 	log.Printf("%s 命令输出:\n %s \n", name, string(combo))
						// }

						if !funcs.In_array(name, []string{"qf_a", "qf_h"}) {
							// UploadFile(sInfo, "D:/default.gif", "/home/wwwroot/default/renren/public/template/pc/images")
							// sInfo.Session.Output(`sed -i "s/{\$item.img}/{\$item.img|default='\/public\/template\/pc\/images\/default.gif'}/g" /home/wwwroot/default/renren/public/template/pc/index.html`)
							// sInfo.Session.Output(`sed -i "s/{\$item.img}/{\$item.img|default='\/public\/template\/pc\/images\/default.gif'}/g" /home/wwwroot/default/renren/public/template/pc/node.html`)
							// sInfo.Session.Output(`sed -i "s/{\$item.img}/{\$item.img|default='\/public\/template\/pc\/images\/default.gif'}/g" /home/wwwroot/default/renren/public/template/pc/cat.html`)
							doCmd(sInfo, name, `sed -i "s/{\$item.img}/{\$item.img|default='\/public\/template\/pc\/images\/default.gif'}/g" /home/wwwroot/default/renren/public/template/pc/cat.html`)
							doCmd(sInfo, name, `sed -i "s/{\$item.img}/{\$item.img|default='\/public\/template\/pc\/images\/default.gif'}/g" /home/wwwroot/default/renren/public/template/pc/node.html`)
							// sInfo.Session.Output(`sed -i "s/{\$item.img}/{\$item.img|default='\/public\/template\/pc\/images\/default.gif'}/g" /home/wwwroot/default/renren/public/template/pc/search.html`)
						}

					case "baidu":
						bdMap := map[string]string{
							"qf_a": "baidu_verify_code-F2i8BLSgMY.html",
							"qf_h": "baidu_verify_code-BJKq9fD0h7.html",
							"90_h": "baidu_verify_code-P3T2z8phzI.html",
						}
						if file, ok := bdMap[name]; ok {
							UploadFile(sInfo, "D:/app/"+file, "/home/wwwroot/default/renren")
						} else {
							log.Println(name, " not found baidu stat file")
						}

					case "init": // 初始化
						// 文件权限处理
						if name == "ys" {
							return
						}
						file := "16.zip"
						if name == "qf_h" || name == "qf_a" {
							file = "69.zip"
						}
						cmd = "cd /home/wwwroot/default && unzip -o " + file + " -d renren && chmod -R 777 renren/runtime && chmod -R 777 renren/public/setup && chmod -R 777 renren/public/addons"
						doCmd(sInfo, name, cmd)

						// 2. 配置nginx 并重启
						upNignx(sInfo, name)
						cmd = "lnmp nginx restart"
						doCmd(sInfo, name, cmd)

						// 3. mysql防火墙处理 单独处理，每个服务器都不一样
						// centos 6:

						// （1）通过vi /etc/sysconfig/iptables
						// 进入编辑增添一条-A INPUT -p tcp -m tcp --dport 3306 -j ACCEPT 即可

						// （2）执行 /etc/init.d/iptables restart 命令将iptables服务重启

						// （3）保存 /etc/rc.d/init.d/iptables save

						// 也可   iptables -I INPUT -p tcp --dport 3306 -j ACCEPT

						// // centos 7

						// firewall-cmd --zone=public --add-port=3306/tcp --permanent
						// firewall-cmd --reload
						// # systemctl stop firewalld.service #停止
						// # systemctl disable firewalld.service #禁用

					default:
						//执行远程命令
						doCmd(sInfo, name, cmd)
					}
				}(sInfo, name)

			}
			wg.Wait()
			fmt.Printf("请输入远程执行命令: ")
			LastTime = time.Now()
		}
	}()

	// 退出 这里可以等待命令执行完成
	s := <-c
	fmt.Println("Got signal:", s)
}

func upNignx(sInfo *sshInfo, name string) {
	remoteConfDir := "/usr/local/nginx/conf/"
	remoteConfPath := remoteConfDir + "nginx.conf"
	localConfPath := "D:/nginx.conf"
	ok, err := funcs.PathExists(localConfPath)
	if err != nil {
		log.Fatalln("local nginx error : ", err)
	}
	if !ok {
		// 下载
		DownLoad(sshList["ys"], "D:", remoteConfPath)
		b, err := ioutil.ReadFile(localConfPath)
		if err != nil {
			log.Fatalln("read local nginx error : ", err)
		}
		if len(b) == 0 {
			log.Fatalln("read empty local nginx error")
		}
	}

	UploadFile(sInfo, localConfPath, remoteConfDir)
}

func publicKeyAuthFunc(kPath string) ssh.AuthMethod {
	keyPath, err := homedir.Expand(kPath)
	if err != nil {
		log.Fatal("find key's home dir failed", err)
	}
	key, err := ioutil.ReadFile(keyPath)
	if err != nil {
		log.Fatal("ssh key file read failed", err)
	}
	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatal("ssh key signer failed", err)
	}
	return ssh.PublicKeys(signer)
}

//获取ftp连接
func getftpclient(client *ssh.Client) *sftp.Client {
	ftpclient, err := sftp.NewClient(client)
	if err != nil {
		fmt.Println("创建ftp客户端失败", err)
		panic(err)
	}
	return ftpclient
}

//上传文件
func UploadFile(sInfo *sshInfo, localpath, remotepath string) {
	ftpclient := getftpclient(sInfo.Client)
	defer ftpclient.Close()

	remoteFileName := path.Base(localpath)
	fmt.Println(localpath, remoteFileName)
	srcFile, err := os.Open(localpath)
	if err != nil {
		fmt.Println("打开文件失败", err)
		panic(err)
	}
	defer srcFile.Close()

	dstFile, e := ftpclient.Create(path.Join(remotepath, remoteFileName))
	if e != nil {
		fmt.Println("创建文件失败", e)
		panic(e)
	}
	defer dstFile.Close()
	// TODO: 这里控制传输速度 根据实际情况来定
	buffer := make([]byte, 1024*128)
	for {
		n, err := srcFile.Read(buffer)
		if err != nil {
			if err == io.EOF {
				// fmt.Println("已读取到文件末尾")
				break
			} else {
				fmt.Println("读取文件出错", err)
				panic(err)
			}
		}
		dstFile.Write(buffer[:n])
		//注意，由于文件大小不定，不可直接使用buffer，否则会在文件末尾重复写入，以填充1024的整数倍
	}
	fmt.Println("文件上传成功")
}

func doCmd(sInfo *sshInfo, name, cmd string) {
	var err error
	sInfo.Session, err = sInfo.Client.NewSession()
	if err != nil {
		log.Printf("%s create seesion error: %s \n", name, err)
		return
	}
	defer sInfo.Session.Close()

	//执行远程命令
	combo, err := sInfo.Session.CombinedOutput(cmd)
	if err != nil {
		log.Printf("%s 远程执行 (%.10s) 失败: %s \n", name, cmd, err)
		return
	}
	log.Printf("%s 命令输出:\n %s \n", name, string(combo))
}

//文件下载
func DownLoad(sInfo *sshInfo, localpath, remotepath string) {
	ftpClient := getftpclient(sInfo.Client)
	defer ftpClient.Close()

	srcFile, err := ftpClient.Open(remotepath)
	if err != nil {
		fmt.Println("文件读取失败", err)
		panic(err)
	}
	defer srcFile.Close()
	localFilename := path.Base(remotepath)
	dstFile, e := os.Create(path.Join(localpath, localFilename))
	if e != nil {
		fmt.Println("文件创建失败", e)
		panic(e)
	}
	defer dstFile.Close()
	if _, err1 := srcFile.WriteTo(dstFile); err1 != nil {
		fmt.Println("文件写入失败", err1)
		panic(err1)
	}
	fmt.Println("文件下载成功")
}

// 远程执行脚本
func Exec_Task(sInfo *sshInfo, localpath, remotepath string) int {
	UploadFile(sInfo, localpath, remotepath)
	session, err := sInfo.Client.NewSession()
	if err != nil {
		fmt.Println("创建会话失败", err)
		return 0
	}
	defer session.Close()
	remoteFileName := path.Base(localpath)
	dstFile := path.Join(remotepath, remoteFileName)
	err1 := session.Run(fmt.Sprintf("/usr/bin/sh %s", dstFile))
	if err1 != nil {
		fmt.Println("远程执行脚本失败", err1)
		return 2
	} else {
		fmt.Println("远程执行脚本成功")
		return 1
	}
}
