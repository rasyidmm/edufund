package enum

type TracingName string

const (
	StartController TracingName = "Start Controller"
	StartService    TracingName = "Start Service"
	Error           TracingName = "Error"
	Response        TracingName = "Response"
	Request         TracingName = "Request"
)
