/**Copyright Blue Medora Inc. 2016**/

package logger

import (
    "os"
    "fmt"
    
    "github.com/cloudfoundry/gosteno"
)

//New logger
func New(logDirectory string) *gosteno.Logger {
    createLogDirectory(logDirectory)
    
    
    loggingConfig := &gosteno.Config {
        Sinks:  []gosteno.Sink{
           gosteno.NewFileSink(fmt.Sprintf("%s/bm_nozzle.log", logDirectory)),  
        },
        Level:      gosteno.LOG_DEBUG,
        Codec:      gosteno.NewJsonCodec(),
        EnableLOC:  true,
    }
    
    gosteno.Init(loggingConfig)
    return gosteno.NewLogger("bm_firehose_nozzle")
}

func createLogDirectory(logDirectory string) {
    if _, err := os.Stat(fmt.Sprintf("%s/", logDirectory)); err == nil {
        os.RemoveAll(fmt.Sprintf("%s", logDirectory))
    }
    
    os.MkdirAll(fmt.Sprintf("%s/", logDirectory), os.ModePerm)
}