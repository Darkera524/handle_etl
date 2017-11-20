package etl

import (
	"time"
	"strings"
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"fmt"
)

func Exec_tracerpt() {
	var interval int64 = 300
	var ticker= time.NewTicker(time.Duration(interval) * time.Second)

	for {
		exec_inter()
		<-ticker.C
	}
}

func exec_inter(){
	dirname := dir_name()
	if (dirname == ""){

	} else{

		filename := dirname+"\\DataCollector01.etl"
		path_list := strings.Split(dirname,"\\")
		single_dir_name := path_list[len(path_list)-1]
		//"E:\\percatch\\" + yesterday + "\\DataCollector01.etl"
		idcshare := "\\\\idcshare.op.internal.gridsumdissector.com\\idcshare\\wangyiqi\\xml\\" + single_dir_name + ".xml"
		c := exec.Command("tracerpt", filename, "-o", idcshare, "-lr", "-of", "XML")
		c.Stdout = os.Stdout
		c.Run()
	}

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

func dir_name() string {
	old_ins := ""
	new_ins := ""
	path := "C:\\PerfLogs\\Admin\\domain-collector\\"
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if ( f == nil ) {return err}
		if !f.IsDir() {return nil}
		old_ins = new_ins
		new_ins = path
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}

	return old_ins
}
