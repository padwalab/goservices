package responsehandler

// ResponseLevel defines the kind of response to write
type ResponseLevel string

// const for different response status
const (
	ERROR ResponseLevel = "ERROR"
	OK    ResponseLevel = "OK"
)

// Response ...
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
