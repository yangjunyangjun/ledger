package lib

import (
	"crypto/md5"
	"fmt"
	"ledger/sdk"
	"math/rand"
	"net"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

//获取本机ip
func GetIp() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		sdk.Log.Errorf("get ip err:%v", err)
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip = strings.Split(localAddr.String(), ":")[0]
	return
}

//md5加密
func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

//随机生成6位激活码
func RandStr() string {
	rand.Seed(time.Now().UnixNano())
	var strList []string
	for i := 0; i < 6; i++ {
		strList = append(strList, strconv.Itoa(rand.Intn(10)))
	}
	return strings.Join(strList, "")
}

// 执行操作系统命令
func Cmd(cmd string, arg ...string, ) []byte {
	// cmd 需要执行第一条命令
	// arg 后续命令
	out, err := exec.Command(cmd, arg...).Output()
	if err != nil {
		panic("some error found")
	}
	return out

}
