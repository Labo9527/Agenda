package service

import (
    "os"
	"log"
)


var log_file *os.File

func Record_output (output string){
	log_file,_ := os.OpenFile("./service/agenda.log",os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	debugLog := log.New(log_file,"[Outputs]",log.LstdFlags)
	debugLog.Println(output)
}

func Record_input (input string){
	log_file,_ := os.OpenFile("./service/agenda.log",os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	debugLog := log.New(log_file,"[Commands]",log.LstdFlags)
	debugLog.Println(input)
}