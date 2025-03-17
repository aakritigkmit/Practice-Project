package utils

import "log"

func PANIC(message string, err error) {
    if err != nil {
        log.Panic(message, err)
    }
}
//panic is used to immediately stop program execution when an unexpected error occurs. 
func INFO(message string, data interface{}) {

    log.Print(message, data)

}

