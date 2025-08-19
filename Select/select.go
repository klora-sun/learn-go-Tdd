package selectCode

import (
	"net/http"
	"time"
)


func Racer(slowUrl,FstUrl string) string{
	startA := time.Now()
	http.Get(slowUrl )
	aDuration := time.Since(startA)

	startB := time.Now()
    http.Get(FstUrl)
    bDuration := time.Since(startB)


    if aDuration < bDuration {
        return slowUrl
    }

    return FstUrl
}