package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"path"
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
		"90_a": {
			Host:    "154.38.228.74",
			User:    "root",
			passwd:  "Lewaimai_123",
			Type:    "password",
			keyPath: "",
			port:    22,
		},
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

	// 执行操作
	go func() {

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("请输入远程执行命令: ")
		for scanner.Scan() {
			cmd := scanner.Text()
			if cmd == "quit" {
				fmt.Println("Quitting...")
				break
			}
			for name, sInfo := range sshList {
				var err error
				sInfo.Session, err = sInfo.Client.NewSession()
				if err != nil {
					log.Printf("%s create seesion error: %s \n", name, err)
				}
				//执行远程命令
				combo, err := sInfo.Session.CombinedOutput(cmd)
				sInfo.Session.Close()
				if err != nil {
					log.Printf("%s 远程执行 (%s) 失败: %s \n", name, cmd, err)
					continue
				}
				log.Printf("%s 命令输出: %s \n", name, string(combo))
			}
			fmt.Printf("请输入远程执行命令: ")
		}
	}()

	// 退出 这里可以等待命令执行完成
	s := <-c
	fmt.Println("Got signal:", s)
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
	buffer := make([]byte, 1024)
	for {
		n, err := srcFile.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println("已读取到文件末尾")
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
