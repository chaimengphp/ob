package controllers

type WebController struct {
	BaseController
}


func (w *WebController) Index() {
	w.TplName = "index.html"
}
