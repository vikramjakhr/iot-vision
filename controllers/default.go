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

type CreateCollectionController struct {
	beego.Controller
}

type DeleteCollectionController struct {
	beego.Controller
}

type IndexFaceController struct {
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

func (c *CreateCollectionController) Get() {
	name := c.GetString("collection")
	c.Data["json"] = services.CreateCollection(name)
	c.ServeJSON();
}

func (c *DeleteCollectionController) Get() {
	name := c.GetString("collection")
	c.Data["json"] = services.DeleteCollection(name)
	c.ServeJSON();
}

func (c *IndexFaceController) Get() {
	collName := c.GetString("collection")
	imageId := c.GetString("imageId")
	img := c.GetString("image")
	c.Data["json"] = services.IndexFaces(collName, imageId, img)
	c.ServeJSON();
}
