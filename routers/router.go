package routers

import (
	"gitlab.intelligrape.net/tothenew/vision/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/faceDetection", &controllers.FaceDetectionController{})
	beego.Router("/textReco", &controllers.TextRecoChanController{})
	beego.Router("/createCollection", &controllers.CreateCollectionController{})
	beego.Router("/indexFaces", &controllers.IndexFaceController{})
}
