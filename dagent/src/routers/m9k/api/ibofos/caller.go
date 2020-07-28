package ibofos

import (
	"a-module/src/routers/m9k/model"
	"bytes"
	"dagent/src/routers/m9k/api"
	"dagent/src/routers/m9k/header"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"io"
	"log"
)

func CalliBoFOS(ctx *gin.Context, f func(string, interface{}) (model.Request, model.Response, error)) {
	req := model.Request{}
	ctx.ShouldBindBodyWith(&req, binding.JSON)
	_, res, err := f(header.XrId(ctx), req.Param)
	api.HttpResponse(ctx, res, err)
}

func CalliBoFOSwithParam(ctx *gin.Context, f func(string, interface{}) (model.Request, model.Response, error), param interface{}) {
	req := model.Request{}
	err := ctx.ShouldBindBodyWith(&req, binding.JSON)

	if err == io.EOF {
		req = model.Request{}
	}

	mergedParam := merge(param, req.Param)
	_, res, err := f(header.XrId(ctx), mergedParam)
	api.HttpResponse(ctx, res, err)
}

func merge(src interface{}, tar interface{}) interface{} {
	var m map[string]interface{}

	ja, _ := json.Marshal(src)
	json.Unmarshal(ja, &m)

	jb, _ := json.Marshal(tar)
	json.Unmarshal(jb, &m)

	jm, _ := json.Marshal(m)

	log.Println("aaaaaaaaaa src", src)
	log.Println("aaaaaaaaaa ja", ja)
	log.Println("aaaaaaaaaa tar", tar)
	log.Println("aaaaaaaaaa jb", jb)
	log.Println("aaaaaaaaaa jb", m)

	var param interface{}

	d := json.NewDecoder(bytes.NewBuffer(jm))
	d.UseNumber()

	if err := d.Decode(&param); err != nil {
		log.Fatal(err)
	}

	return param
}
