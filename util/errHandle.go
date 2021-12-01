package util

import (
	. "gin/log"
	"os"
)

func HandleErr(msg string, err error, args ...string) {
	if err != nil {

		fMap := make(map[string]interface{}, 0)
		fMap["err"] = err
		InfoWithFields(msg+": ", fMap)
		if args[0] == "exit1" {
			os.Exit(1)
		} else if args[0] == "return" {
			return
		}
	}
}
