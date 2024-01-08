package service

import (
	"context"
	"encoding/json"
	"fmt"
	"spun/internal/model"
	"spun/internal/repository"
	"spun/pkg/liberror"
	"spun/pkg/libsession"
	"strconv"
)

// AuditTrailService provides methods to interact with the audit trail repository
type AuditTrailService struct {
	repo repository.AuditTrailRepository
}

// NewAuditTrailService creates a new instance of AuditTrailService
func NewAuditTrailService(repo repository.AuditTrailRepository) *AuditTrailService {
	return &AuditTrailService{
		repo: repo,
	}
}

// CreateAuditTrail
func (s *AuditTrailService) CreateAuditTrail(ctx context.Context, modelName string, modelKey interface{},
	action string, logData interface{}, remark string) *liberror.Error {
	// Prepare from param
	audit := new(model.AuditTrail)
	audit.Action = action
	audit.Remark = remark
	audit.ModelName = modelName

	// Prepare model key from string or int64
	var modelKeyStr string
	switch key := modelKey.(type) {
	case int64:
		modelKeyStr = strconv.FormatInt(key, 10)
	case string:
		modelKeyStr = key
	default:
		modelKeyStr = fmt.Sprintf("%v", modelKey)
	}
	audit.ModelKey = modelKeyStr

	// Convert log data to json string
	if logData != nil {
		b, _ := json.Marshal(logData)
		audit.Log = string(b)
	}

	// Prepare from session
	session, _ := libsession.FromContext(ctx)
	audit.CreatedBy = session.UserID
	audit.TokenID = strconv.FormatInt(session.TokenID, 10)
	audit.IPAddress = session.IPAddress

	audit.GenerateID()
	err := s.repo.Create(audit)
	if err != nil {
		return liberror.NewServerError(err.Error())
	}
	return nil
}

// logAuditTrail wrapper for CreateAuditTrail
func logAuditTrail(svcAudit *AuditTrailService, ctx context.Context, modelName string, modelKey interface{}, action string, logData interface{}, remark string) {
	if svcAudit != nil {
		go func() {
			svcAudit.CreateAuditTrail(ctx, modelName, modelKey, action, logData, remark)
		}()
	}
}
