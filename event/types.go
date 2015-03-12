package event

import "github.com/concourse/atc"

const (
	EventTypeInvalid atc.EventType = ""

	// build log (e.g. from input or build execution)
	EventTypeLog atc.EventType = "log"

	// build status change (e.g. 'started', 'succeeded')
	EventTypeStatus atc.EventType = "status"

	// build initializing (all inputs fetched; fetching image)
	EventTypeInitializeTask atc.EventType = "initialize-task"

	// build execution started
	EventTypeStartTask atc.EventType = "start-task"

	// build execution finished
	EventTypeFinishTask atc.EventType = "finish-task"

	// finished getting something
	EventTypeFinishGet atc.EventType = "finish-get"

	// finished putting something
	EventTypeFinishPut atc.EventType = "finish-put"

	// error occurred
	EventTypeError atc.EventType = "error"
)
