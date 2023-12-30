package repository

import (
	"fmt"
	"strconv"

	"github.com/casbin/casbin/v2"
	"gorm.io/gorm"

	"iam-service/internal/domain"
	"iam-service/internal/model"
)

type IAMRepository struct {
	db       *gorm.DB
	enforcer *casbin.Enforcer
}

func NewIAMRepository(db *gorm.DB, enforcer *casbin.Enforcer) domain.IAMRepositoryInterface {
	iam := &IAMRepository{
		db:       db,
		enforcer: enforcer,
	}

	iam.InitPolicies()
	iam.InitRoleAssign()

	return iam
}

// InitPolicies initializes access control policies.
// It retrieves all TenantRoles from the database and their associated TenantResourceActions,
// and adds corresponding policies to the Casbin enforcer.
func (ir *IAMRepository) InitPolicies() error {
	allRoles := []model.TenantRole{}
	result := ir.db.Preload("TenantResourceActions").Find(&allRoles)
	if result.Error != nil {
		return result.Error
	}

	for _, role := range allRoles {
		for _, action := range role.TenantResourceActions {
			if _, err := ir.enforcer.AddPolicy(
				fmt.Sprintf(fmt.Sprintf("r_%d", role.ID)),               // role
				strconv.FormatUint(uint64(role.TenantID), 10),           // domain
				strconv.FormatUint(uint64(action.TenantResourceID), 10), // resource
				strconv.FormatUint(uint64(action.ID), 10),               // action
			); err != nil {
				return err
			}
		}
	}
	return nil
}

// InitRoleAssign initializes role assignments.
// It retrieves all TenantSubjects from the database along with their associated TenantRoles,
// and establishes role assignments using the Casbin enforcer.
func (ir *IAMRepository) InitRoleAssign() error {
	allTenantSubjects := []model.TenantSubject{}
	result := ir.db.Preload("TenantRoles").Find(&allTenantSubjects)
	if result.Error != nil {
		return result.Error
	}

	for _, subject := range allTenantSubjects {
		for _, role := range subject.TenantRoles {
			if _, err := ir.enforcer.AddRoleForUserInDomain(
				strconv.FormatUint(uint64(subject.ID), 10),       // subject
				fmt.Sprintf("r_%d", role.ID),                     // role
				strconv.FormatUint(uint64(subject.TenantID), 10), // tenant
			); err != nil {
				return err
			}
		}
	}
	return nil
}
