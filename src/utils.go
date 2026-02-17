package utils
import (
	"fmt"
	"log"
	"net/http"
)


//***************Logging*******************************//
//message to log and code is the severity of the message
//codes= {1:Console Log,2:Warning log,3:Error log,4:Crash happend}
func log(message:string, code:int){}

func consoleLog(message:string){}

func logToFile(message:string){}



//*************Configuration files*********************
func readConfigFile(filePath:string){}






