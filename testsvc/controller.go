package testsvc

import (
	"net/http"
	"time"

	"github.com/padwalab/goservices/logger"
	rqh "github.com/padwalab/goservices/requesthandler"
	rph "github.com/padwalab/goservices/responsehandler"
)

// SampleService is a sample service for demo
func SampleService(w http.ResponseWriter, r *http.Request) {

	defer logger.TimeLogger("TestService", time.Now())

	reqData := rqh.ExtractRequestDetails(r)

	result, err := SomeOperation(reqData)
	if err != nil {
		rph.WriteResponse(w, rph.ERROR, err, "failed to perform operation")
		return
	}

	rph.WriteResponse(w, rph.OK, result, "Operation successful")
	return
}
