package instance

import (
	"errors"
	"fmt"
)

const (
	//PendingResponse -
	PendingResponse = `HaaS Instance not ready. Please try again in a few seconds.`
	//FailureResponse -
	FailureResponse = `HaaS Instance could not be provisioned. Sorry.`
	//InstanceIDVarName --
	InstanceIDVarName = "instance_id"
	//TaskGUIVarName --
	TaskGUIDVarName = "taskguid"
	//TaskStatusComplete ---
	TaskStatusComplete = "complete"
	//AgentTaskStatusFailed ---
	TaskStatusFailed = "failed"
	//CollectionInstanceIDQueryField --
	CollectionInstanceIDQueryField = "instanceid"
	//RequestIDMetadataFieldname -- fieldname for the metadata requestid
	RequestIDMetadataFieldname = "requestid"
	//SuccessGetHandlerBody --
	SuccessGetHandlerBody = `{
		"state": "succeeded",
		"description": "%s"
	}`
	//FailureGetHandlerBody --
	FailureGetHandlerBody = `{
		"state": "failed",
		"description": "%s"
	}`
	//PendingGethandlerBody --
	PendingGethandlerBody = `{
		"state": "in progress",
		"description": "%s"
	}`
)

var (
	//SSOPathPrefix -
	SSOPathPrefix = "/sso"
	//ServiceInstanceDash -
	ServiceInstanceDash = fmt.Sprintf("/dashboard/{%s}", TaskGUIDVarName)
	//HandlerPath - path to normal instance handlers
	HandlerPath = fmt.Sprintf("/service_instances/{%s}", InstanceIDVarName)
	//AsyncHandlerPath - path to async poller
	AsyncHandlerPath     = "/service_instances/{instance_id}/last_operation"
	dashboardUrl         = "https://www.pezapp.io"
	ErrInvalidInstanceID = errors.New("invalid instance id while attempting to get taskid")
	//ErrNoRecordsInResult - empty result set
	ErrNoRecordsInResult = errors.New("no records found in result set")
)
