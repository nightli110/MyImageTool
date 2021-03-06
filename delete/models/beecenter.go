package models

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"myimagetool/common"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

type CenterMsg struct {
	imageID   string
	imagecode string
}

func (msg *CenterMsg) ImagetoBase64(imagepath string) (bool, error) {
	if !common.IsExists(imagepath) {
		return false, errors.New("Image not exit")
	}
	ImageFile, err := os.Open(imagepath)

	sourcebuffer := make([]byte, 500000)

	if err != nil {
		return false, err
	}
	n, err := ImageFile.Read(sourcebuffer)
	if err != nil {
		return false, err
	}

	msg.imagecode = base64.StdEncoding.EncodeToString(sourcebuffer[:n])

	return true, nil

}

func (msg *CenterMsg) CenterMsgTOJSON() ([]byte, error) {
	Msgjson, err := json.Marshal(msg)
	if err == nil {
		return nil, err
	}
	return Msgjson, nil
}

func (msg *CenterMsg) MsgPost(postjson []byte) ([]byte, error) {
	CenterHost := beego.AppConfig.String("centerhost")
	CenterPort := beego.AppConfig.String("centerport")
	CenterUrl := CenterHost + ":" + CenterPort
	Req := httplib.Post(CenterUrl)
	Req.JSONBody(msg)

}

func (msg *CenterMsg) CnterPost(imagepath string) (string, error) {

}
