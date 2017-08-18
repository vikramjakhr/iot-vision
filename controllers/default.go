package controllers

import (
	"github.com/astaxie/beego"
	"log"
	"gitlab.intelligrape.net/tothenew/vision/services"
)

type MainController struct {
	beego.Controller
}

type FaceDetectionController struct {
	beego.Controller
}

type TextRecoChanController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.tpl"
}

func (c *FaceDetectionController) Get() {
	c.TplName = "face.tpl"
}

func (c *TextRecoChanController) Get() {
	select {
	case message, ok := <-services.TextRecoChan:
		if ok {
			c.Data["json"] = message
		} else {
			log.Println("Notification Channel is closed!")
		}
	default:
		c.Data["json"] = services.TextReco{}
	}
	c.ServeJSON();
}
