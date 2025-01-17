// FILE: alertmanager/alertmanager.go
package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/AlertFlow/runner/config"
	"github.com/AlertFlow/runner/pkg/models"
	"github.com/AlertFlow/runner/pkg/payloads"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Receiver struct {
	Receiver string `json:"receiver"`
}

type AlertmanagerPayloadPlugin struct{}

func (p *AlertmanagerPayloadPlugin) Init() models.Plugin {
	return models.Plugin{
		Name:    "Alertmanager",
		Type:    "payload_endpoint",
		Version: "1.0.10",
		Creator: "JustNZ",
	}
}

func (p *AlertmanagerPayloadPlugin) Details() models.PluginDetails {
	return models.PluginDetails{
		Payload: models.PayloadEndpoint{
			Name:     "Alertmanager",
			Type:     "alertmanager",
			Endpoint: "/alertmanager",
		},
	}
}

func (p *AlertmanagerPayloadPlugin) Execute(execution models.Execution, flow models.Flows, payload models.Payload, steps []models.ExecutionSteps, step models.ExecutionSteps, action models.Actions) (data map[string]interface{}, finished bool, canceled bool, no_pattern_match bool, failed bool) {
	return nil, false, false, false, false
}

func (p *AlertmanagerPayloadPlugin) Handle(context *gin.Context) {
	log.Info("Received Alertmanager Payload")
	incPayload, err := io.ReadAll(context.Request.Body)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request body",
		})
		return
	}

	receiver := Receiver{}
	json.Unmarshal(incPayload, &receiver)

	payloadData := models.Payload{
		Payload:  incPayload,
		FlowID:   receiver.Receiver,
		RunnerID: config.Config.Alertflow.RunnerID,
		Endpoint: "alertmanager",
	}

	payloads.SendPayload(payloadData)
}

var Plugin AlertmanagerPayloadPlugin
