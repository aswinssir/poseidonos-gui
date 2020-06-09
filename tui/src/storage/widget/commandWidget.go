package widget

import (
	iBoFOS "a-module/routers/m9k/api/ibofos"
	"a-module/routers/m9k/model"
	"fmt"
	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/google/uuid"
	"github.com/tidwall/gjson"
)

type commandWidget struct {
	Widget   *widgets.List
	List     []string
	Position int
}

func (cm *commandWidget) Input(key string) int {
	switch key {
	case "j", "J":
		cm.Position++
	case "k", "K":
		cm.Position--
		if cm.Position < 0 {
			cm.Position = 0
		}
	case "<Enter>":
		excuteCommand(cm.Position)
	}
	cm.Widget.Rows = cm.List[cm.Position%len(commandList):]
	return cm.Position
}

func excuteCommand(position int) {
	newUUID, _ := uuid.NewUUID()
	xrId := newUUID.String()

	switch position {
	case 0:
		iBoFRequest, res, err := iBoFOS.RuniBoFOS(xrId, nil)
		showResult(iBoFRequest, res, err)
	case 1:
		iBoFRequest, res, err := iBoFOS.ExitiBoFOS(xrId, nil)
		showResult(iBoFRequest, res, err)
	case 2:
		iBoFRequest, res, err := iBoFOS.ScanDevice(xrId, nil)
		showResult(iBoFRequest, res, err)
	case 3:
		iBoFRequest, res, err := iBoFOS.ListDevice(xrId, nil)
		showResult(iBoFRequest, res, err)

		if err != nil && res.Result.Data != nil {
			value := gjson.Get(res.Result.Data.(string), "devicelist[0].name")
			Device.Widget.Text = value.String()
		}
	case 4:
		deviceName := "unvme-ns-0"
		param := model.DeviceParam{Name: deviceName}
		iBoFRequest, res, err := iBoFOS.GetSMART(xrId, param)
		showResult(iBoFRequest, res, err)

		if err != nil && res.Result.Data != nil {
			value := gjson.Get(res.Result.Data.(string), "read_only")
			Volume.Widget.Text = value.String()
		}
	}
}

func showResult(iBoFRequest model.Request, res model.Response, err error) {
	if err != nil {
		errStr := fmt.Sprintf("%+v", err)
		Info.Widget.Text = iBoFRequest.Command + " Error : " + errStr
	} else {
		dataStr := fmt.Sprintf("%+v", res.Result.Data)
		Info.Widget.Text = iBoFRequest.Command + " Success\n" + res.Result.Status.Description + "\n" + dataStr
	}
}

var Command = commandWidget{
	Widget:   initCommand(),
	List:     commandList,
	Position: 0,
}

var commandList = []string{
	"[0] Run POS",
	"[1] Exit POS",
	"[2] Scan Device",
	"[3] List Device",
	"[4] SMART",
	"[5] Load Array (TBD)",
	"[6] Delete Array (TBD)",
	"[7] State Array (TBD)",
	"[8] Mount POS (TBD)",
	"[9] Unmount POS (TBD)",
}

func initCommand() *widgets.List {
	list := widgets.NewList()
	list.Title = "Command (J & K)"
	list.Rows = commandList
	list.SetRect(0, 6, 25, 15)
	list.TextStyle.Fg = termui.ColorYellow

	return list
}
