package v1

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"ibofdagent/server/handler"
	"ibofdagent/server/routers/mtool/api"
	"ibofdagent/server/routers/mtool/model"
	"log"
	"net/http"
	"sync/atomic"
)

const (
	stateUnlocked uint32 = iota
	stateLocked
)

var (
	locker    = stateUnlocked
	errLocked = errors.New("Locked out buddy")
)

func sendIBoF(ctx *gin.Context, iBoFRequest model.Request) {
	if !atomic.CompareAndSwapUint32(&locker, stateUnlocked, stateLocked) {
		log.Printf("sendIBoF : %+v", iBoFRequest)
		api.MakeBadRequest(ctx, 12000)
		return
	}
	defer atomic.StoreUint32(&locker, stateUnlocked)

	log.Printf("sendIBoF : %+v", iBoFRequest)
	marshaled, _ := json.Marshal(iBoFRequest)
	err := handler.WriteToIBoFSocket(marshaled)

	if err != nil {
		log.Printf("sendIBoF : %v", err)
		api.MakeFailResponse(ctx, 400, error.Error(err), 19002)
		return
	}

	for {
		temp := handler.GetIBoFResponse()
		log.Printf("Response From iBoF : %s", string(temp))

		response := model.Response{}
		err := json.Unmarshal(temp, &response)

		if response.Rid != "timeout" && iBoFRequest.Rid != response.Rid {
			log.Printf("Previous request's response, Wait again")
			continue
		}

		if err != nil {
			log.Printf("Response Unmarshal Error : %v", err)
			api.MakeFailResponse(ctx, 400, error.Error(err), 12310)
			return
		} else if response.Result.Status.Code != 0 {
			api.MakeBadRequest(ctx, response.Result.Status.Code)

			return
		} else {
			ctx.JSON(http.StatusOK, &response)
			return
		}
	}
}

func makeRequest(ctx *gin.Context, command string) model.Request {
	xrId := ctx.GetHeader("X-request-Id")
	request := model.Request{
		Command: command,
		Rid:     xrId,
	}
	return request
}
