package Tools

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

//json图片处理
func HandelHeadImg(img string) []string {
	var img_list []string
	json.Unmarshal([]byte(img),&img_list)
	var img_list_domain []string
	for _,img := range img_list {
		img_url := fmt.Sprintf("http://img.wepethome.com/%s",img)
		img_list_domain = append(img_list_domain,img_url)
	}
	return img_list_domain
}

//单张图片处理
func HandelImg(img string) string {
	img_url := fmt.Sprintf("http://img.wepethome.com/%s",img)
	return img_url
}

//图片上传
func UpImg(imgBase64 string) (imgpath string,err error) {
	imgs, err := base64.StdEncoding.DecodeString(imgBase64)
	if err != nil {
		return "",errors.New("base64 decode error")
	}
	timenow := time.Now()

	//定义图片目录
	dir := "/data/image/"+timenow.Format("20060102")
	if _,err := os.Stat(dir);err !=nil {
		os.MkdirAll(dir,0777)
	}
	img_path := "/"+strconv.FormatInt(timenow.Unix(), 10)+".jpg"
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
	return timenow.Format("20060102")+img_path,nil
}
