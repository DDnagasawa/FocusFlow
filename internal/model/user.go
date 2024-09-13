package model

import (
	"database/sql"
	"encoding/json"
	"time"
)

type ADHDSeverity string

const (
	Mild     ADHDSeverity = "mild"
	Moderate ADHDSeverity = "moderate"
	Severe   ADHDSeverity = "severe"
)

type User struct {
	ID                  int             `json:"id"`
	Username            string          `json:"username"`
	Email               string          `json:"email"`
	PasswordHash        string          `json:"-"`
	CreatedAt           time.Time       `json:"created_at"`
	LastLogin           sql.NullTime    `json:"last_login"`
	NotificationEnabled bool            `json:"notification_enabled"`
	Theme               string          `json:"theme"`
	FocusModeSettings   json.RawMessage `json:"focus_mode_settings"`
	ADHDDiagnosisDate   sql.NullTime    `json:"adhd_diagnosis_date"`
	ADHDSeverity        ADHDSeverity    `json:"adhd_severity"`
	Medications         sql.NullString  `json:"medications"`
}
