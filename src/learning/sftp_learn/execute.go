package main

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
)

func main() {
	addr := "ecom.decta.com:822"

	file, err := ioutil.ReadFile("/Users/vladgoncharenko/Documents/golang/projects/tips/src/learning/sftp_learn/ssh-key-decta.key")

	if err != nil {
		fmt.Println(err)
	}
	key, err := ssh.ParsePrivateKeyWithPassphrase(file, []byte("dKZDNqNvXymFp68K"))

	if err != nil {
		fmt.Println(err)
	}

	config := &ssh.ClientConfig{
		User: "DCT.BETTER",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Config: ssh.Config{
			Ciphers: []string{"3des-cbc", "aes256-cbc", "aes192-cbc", "aes128-cbc", "aes128-ctr", "aes192-ctr", "aes256-ctr"}},
	}

	conn, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		fmt.Println("Failed to dial: " + err.Error())
	}

	client, err := sftp.NewClient(conn)
	if err != nil {
		fmt.Println("Failed to dial: " + err.Error())
	}
	defer client.Close()

	data, err := client.ReadDir("/in_file/reports")

	if err != nil {
		fmt.Println("Failed to dial: " + err.Error())
	}

	for _, v := range data {
		fmt.Println(v.Name())
	}

	//cwd, err := client.Getwd()

	//f, err := ioutil.ReadFile("./in_file/reports/BETTER_chrgbck_20201023.csv")
	//f, err := client.Open("/in_file/reports/BETTER_chrgbck_20201023.csv")
	//if err != nil {
	//	fmt.Println("Failed to dial: " + err.Error())
	//}
	//
	//reader, err := ioutil.ReadAll(f)

	//fmt.Println(string(reader))
	//println("Current working directory:", cwd)
}
