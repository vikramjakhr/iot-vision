package routers

import (
	"gitlab.intelligrape.net/tothenew/iot-vision/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/faceDetection", &controllers.FaceDetectionController{})
	beego.Router("/textReco", &controllers.TextRecoChanController{})
	beego.Router("/createCollection", &controllers.CreateCollectionController{})
	beego.Router("/deleteCollection", &controllers.DeleteCollectionController{})
	beego.Router("/indexFaces", &controllers.IndexFaceController{})
}
