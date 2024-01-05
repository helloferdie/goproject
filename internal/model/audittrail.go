package model

import (
	"database/sql"
	"os"
	"strconv"
	"time"

	"github.com/sony/sonyflake"
)

type AuditTrail struct {
	ID        string       `db:"id" json:"id"`
	ModelName string       `db:"model_name" json:"model_name"`
	ModelKey  string       `db:"model_key" json:"model_key"`
	Action    string       `db:"action" json:"action"`
	Log       string       `db:"log" json:"log"`
	Remark    string       `db:"remark" json:"remark"`
	IPAddress string       `db:"ip_address" json:"ip_address"`
	TokenID   string       `db:"token_id" json:"token_id"`
	CreatedBy int64        `db:"created_by" json:"created_by"`
	CreatedAt sql.NullTime `db:"created_at" json:"created_at" insert:"-"`
}

var initializeAuditSF = false
var auditSF *sonyflake.Sonyflake

func loadAuditSF() {
	if !initializeAuditSF {
		envTZ := os.Getenv("app_timezone")
		loc, err := time.LoadLocation(envTZ)
		if err != nil {
			loc, _ = time.LoadLocation("UTC")
		}

		sfTime, _ := time.ParseInLocation("2006-01-02 15:04:05", "2023-01-01 00:00:00", loc)

		st := sonyflake.Settings{
			StartTime: sfTime,
			MachineID: func() (uint16, error) {
				envID := os.Getenv("app_machine_id")
				if envID == "" {
					envID = "1"
				}
				mID, err := strconv.ParseUint(envID, 10, 16)
				if err != nil {
					mID = 1
				}
				return uint16(mID), nil
			},
		}
		auditSF = sonyflake.NewSonyflake(st)
		initializeAuditSF = true
	}
}

// GenerateID return random sequential ID using Sonyflake
func (m *AuditTrail) GenerateID() {
	loadAuditSF()

	id, err := auditSF.NextID()
	for err != nil {
		id, err = auditSF.NextID()
	}
	m.ID = strconv.FormatUint(id, 10)
}
