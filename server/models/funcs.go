package models

import (
	"bufio"
	"github.com/gin-gonic/gin"
	"io"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
)

func sayhello(s string) {

}
func Check(err error) {
	if err != nil {
		panic(err)
	}
}
func LoadHTML(c *gin.Context, fname string) {
	//加载html页面
	file, err := os.Open(fname)
	Check(err)
	reader := bufio.NewReader(file)
	for {
		var b []byte
		b, err = reader.ReadBytes('\n')
		if err == io.EOF {
			break
		} else {
			c.Writer.Write(b)
		}
	}
	file.Close()
}

func FormID(rely int) string {
	//根据数字生成唯一ID
	var curtime = time.Now().Format("20060102")
	anabytes := []byte(curtime)
	lens := len(anabytes)
	for i := 0; i < lens; i++ {
		anabytes[i] = byte((int(anabytes[i])-int('0')+rely)%(i+2) + int('0'))
	}
	//字节位置交换
	newbyte := anabytes
	newbyte[6] = anabytes[0]
	newbyte[0] = anabytes[6]
	newbyte[5] = anabytes[7]
	newbyte[7] = anabytes[5]
	trans := string(anabytes)
	trans = trans + string(rely%10+'0')
	return trans
}

func GetOutBoundIP() string {
	//获取本机IP
	//未进行公网部署时需要使用
	conn, err := net.Dial("udp", "8.8.8.8:53")
	Check(err)
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip := strings.Split(localAddr.String(), ":")[0]
	return ip
}
func Init_user_dir(_id string) {
	//创建用户文件夹并初始化
	err := os.MkdirAll("./usersdata/"+_id, os.ModePerm)
	Check(err)
	os.MkdirAll("./usersdata/"+_id+"/pic", os.ModePerm)
	os.Create("./usersdata/" + _id + "/log.txt")
	os.Chmod("./usersdata/"+_id+"/log.txt", os.ModePerm)
}
func Write_userlog(_id string, logs string) {
	//书写日志
	file, err := os.OpenFile(".\\usersdata\\"+_id+"\\log.txt", os.O_APPEND, os.ModePerm)
	Check(err)
	time := time.Now().String()
	file.WriteString(time + logs + "\n")
	file.Close()
}

func Form_rannum(top int) int {
	//生成一个随机数
	rand.Seed(time.Now().Unix())
	return rand.Intn(top)
}

func Tran_file(src string, dst string) {
	//进行文件复制
	sfile, err := os.Open(src)
	Check(err)
	_, err = os.Stat(dst)
	if os.IsNotExist(err) {
		os.Create(dst)
	}
	dfile, err := os.OpenFile(dst, os.O_RDWR, os.ModePerm)
	Check(err)
	io.Copy(dfile, sfile)
}

func ReadPic(picname string) []byte {
	file, err := os.Open(picname)
	Check(err)
	PicBlob, _ := io.ReadAll(file)
	return PicBlob
}
