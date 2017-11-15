package etl

import (
	"time"
	"strings"
	"bytes"
	"os"
	"os/exec"
)

func Exec_tracerpt(){
	yesterday := get_yesterday()
	filename := "E:\\percatch\\" + yesterday + "\\DataCollector01.etl"
	idcshare := "\\\\idcshare.op.internal.gridsumdissector.com\\idcshare\\wangyiqi\\xml\\" + yesterday + ".xml"
	c := exec.Command("tracerpt", filename , "-o", idcshare, "-lr", "-of", "XML")
	c.Stdout = os.Stdout
	c.Run()

}

func get_yesterday() string {
	timestamp := time.Now().Unix()-86400
	now := time.Unix(timestamp, 0)
	date := strings.Split(strings.Split(now.String(), " ")[0], "-")
	var buffer bytes.Buffer
	for _,str := range date{
		buffer.WriteString(str)
	}
	yesterday := buffer.String()
	return yesterday
}
