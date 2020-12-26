package controller

type BaseController struct {
}

func (baseController *BaseController) GetJwtKey() string {
	return "holiday"
}
