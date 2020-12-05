package Tools

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"time"
)

//图片处理
func HandelHeadImg(headimg string) []string {
	var img_list []string
	json.Unmarshal([]byte(headimg),&img_list)
	return img_list
}

//图片上传
func UpImg(imgBase64 string) (imgpath string,err error) {
	imgs, err := base64.StdEncoding.DecodeString(imgBase64)
	if err != nil {
		return "",errors.New("base64 decode error")
	}
	timenow := time.Now()

	//定义图片目录
	dir := "D:/go/goWork/src/obapi/image/"
	img_path := timenow.Format("20060102")+strconv.FormatInt(timenow.Unix(), 10)+".jpg"
	save_path := dir+img_path
	file, err := os.OpenFile(save_path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return "",errors.New("create file error")
	}
	defer file.Close()
	w:=bufio.NewWriter(file)
	_, err3 := w.WriteString(string(imgs))
	if err3 != nil {
		return "",errors.New("write error")
	}
	w.Flush()
	return img_path,nil
}
