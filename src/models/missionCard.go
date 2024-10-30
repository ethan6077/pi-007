package models

import (
	"github.com/google/uuid"
)

type MissionCard struct {
	Status         string    `json:"status"`
	PiDataDetected bool      `json:"piDataDetected"`
	MissionId      uuid.UUID `json:"missionId"`
	MissionKey     string    `json:"missionKey"`
	Body           Form      `json:"body"`
}
