package Tools

import "encoding/json"

//图片处理
func HandelHeadImg(headimg string) []string {
	var img_list []string
	json.Unmarshal([]byte(headimg),&img_list)
	return img_list
}
