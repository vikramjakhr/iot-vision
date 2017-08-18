package routers

import (
	"gitlab.intelligrape.net/tothenew/ttn-iot/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
