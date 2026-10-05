package deployments

import (
	"context"
	"io"

	"github.com/google/uuid"
	"github.com/lkfox993/orkio-core/internal/api/workflows"
	"gorm.io/gorm"
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetByID(ctx context.Context, id int64) (*Deployment, error) {
	return s.repository.GetByID(ctx, id)
}

func (s *Service) GetAll(ctx context.Context) ([]Deployment, error) {
	return s.repository.GetAll(ctx)
}

func (s *Service) Create(ctx context.Context, input CreateDeploymentInput) (*Deployment, error) {

	var tenantID = uuid.New()

	deployment := &Deployment{
		ID:       uuid.New(),
		TenantID: tenantID,
		Name:     input.Name,
	}

	err := s.repository.db.Transaction(func(tx *gorm.DB) error {
		if err := s.repository.Create(ctx, deployment, tx); err != nil {
			return err
		}

		for _, resource := range input.Resources {

			reader, err := resource.Open()

			if err != nil {
				return err
			}

			data, err := io.ReadAll(reader)

			if err != nil {
				return err
			}

			err = reader.Close()

			if err != nil {
				return err
			}

			workflow := &workflows.Workflow{
				ID:       uuid.New(),
				TenantID: tenantID,
				Key:      "workflow-id",
				Name:     resource.Name, // Daily Report
			}

			if err := tx.Create(workflow).Error; err != nil {
				return err
			}

			version := &workflows.WorkflowVersion{
				ID:           uuid.New(),
				DeploymentID: deployment.ID,
				TenantID:     tenantID,
				WorkflowID:   workflow.ID,
				Version:      1,
				ResourceName: resource.Name, // daily-report.bpmn
				Definition:   string(data),
			}

			if err := tx.Create(version).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return deployment, nil
}

func (s *Service) Update(ctx context.Context, id int64, input UpdateDeploymentInput) (*Deployment, error) {
	deployment := &Deployment{
		Name: input.Name,
	}

	return s.repository.Update(ctx, id, deployment)
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	return s.repository.Delete(ctx, id)
}

// <bpmn:process
// id="daily-report"
// name="Daily Report"
// isExecutable="true">
