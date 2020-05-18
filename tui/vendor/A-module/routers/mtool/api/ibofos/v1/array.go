package v1

import (
	"A-module/routers/mtool/model"
)

func ListArrayDevice(xrId string, param interface{}) (model.Request, model.Response, error) {
	return Requester{xrId, param}.Get("LISTARRAYDEVICE")
}

func LoadArray(xrId string, param interface{}) (model.Request, model.Response, error) {
	return Requester{xrId, param}.Get("LOADARRAY")
}

func CreateArray(xrId string, param interface{}) (model.Request, model.Response, error) {
	return Requester{xrId, param}.Post("CREATEARRAY")
}

func DeleteArray(xrId string, param interface{}) (model.Request, model.Response, error) {
	return Requester{xrId, param}.Delete("DELETEARRAY")
}
