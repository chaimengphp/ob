package controllers
//
//import (
//	"encoding/json"
//	"github.com/qiniu/api.v7/auth/qbox"
//	"github.com/qiniu/api.v7/storage"
//	"io/ioutil"
//	"net/http"
//)
//
//type QiniuController struct {
//	BaseController
//}
//
//func (qn *QiniuController) GetUploadToken() {
//	putPolicy := storage.PutPolicy{
//		Scope:"petfamily",
//		CallbackURL:"http://api.example.com/upload/upload/callback",
//		CallbackBody:`{"key":"${key}","fsize":$(fsize),"bucket":"$(bucket)"}`,
//		CallbackBodyType: "application/json",
//	}
//	mac := qbox.NewMac("Ka2EtsaaLRf-lA29pk75xGYCBNWIfVHooBS-gQw_", "6nmYIfESBnqZYfBOpEKD0tS98-YPQb6SnB9Ff1g_")
//	upToken := putPolicy.UploadToken(mac)
//	result := struct {
//		Token string
//	}{
//		upToken,
//	}
//	qn.ResponseData(0,"suss",result)
//}
//
//func (qn *QiniuController) Callback(r *http.Request) {
//	type MyPutRet struct {
//		Key    string
//		Hash   string
//		Fsize  int
//		Bucket string
//	}
//	mac := qbox.NewMac("Ka2EtsaaLRf-lA29pk75xGYCBNWIfVHooBS-gQw_","6nmYIfESBnqZYfBOpEKD0tS98-YPQb6SnB9Ff1g_")
//	//httpRequest,_ := http.NewRequest("post","http://api.example.com/upload/upload/callback",reader)
//	if is_verfiy,_ := qbox.VerifyCallback(mac, r);is_verfiy == false {
//		return
//	}
//	body, err := ioutil.ReadAll(r.Body)
//	if err !=nil {
//		return
//	}
//	MyPutReto := new(MyPutRet)
//	errj := json.Unmarshal(body,&MyPutReto)
//	if errj !=nil {
//		return
//	}
//
//	result := make(map[string]string)
//	result["video_url"] = MyPutReto.Key
//	qn.ResponseData(0,"suss",result)
//}