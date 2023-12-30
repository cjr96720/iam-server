package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// gorm.Model definition
// type Model struct {
// 	ID        uint           `gorm:"primaryKey"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt gorm.DeletedAt `gorm:"index"`
// }

type Application struct {
	gorm.Model
	Name      string
	AppID     uuid.UUID `gorm:"unique"`
	Secret    string
	Tenants   []Tenant  // Application has many Tenants, Application.ID is the foreign key
	Subjects  []Subject // Application has many Subjects, Application.ID is the foriegn key
	CreatedAt int
	UpdatedAt int
}

type Tenant struct {
	gorm.Model
	ApplicationID   uint
	Name            string
	Subjects        []Subject        `gorm:"many2many:tenant_subjects"` // Tenant has many Subject, `tenant_subjects` is the join table
	TenantResources []TenantResource // Tenant has many TenantResources, Tenant.ID is the foriegn key
	TenantRoles     []TenantRole     // Tenant has many TenantRoles, Tenant.ID is the foriegn key
	CreatedAt       int
	UpdatedAt       int
}

type Subject struct {
	gorm.Model
	ApplicationID uint
	ExternalID    string
	Name          string
	Tenants       []Tenant `gorm:"many2many:tenant_subjects"` // Subject has many Tenants, `tenant_subjects` is the join table
	CreatedAt     int
	UpdatedAt     int
}

type TenantSubject struct {
	gorm.Model
	TenantID    uint
	Tenant      Tenant // TenantSubject belongs to Tenant, `TenantID` is the foreign key
	SubjectID   uint
	Subject     Subject      // TenantSubject belongs to Subject, `SubjectID is the foreign key`
	TenantRoles []TenantRole `gorm:"many2many:tenant_role_subjects"` // TenantSubject has many TenantRoles, `tenant_role_subjects` is the join table
	CreatedAt   int
	UpdatedAt   int
}

type TenantRole struct {
	gorm.Model
	Name                  string
	TenantID              uint
	TenantSubjects        []TenantSubject        `gorm:"many2many:tenant_role_subjects"` // TenantRole has many TenantSubjects, `tenant_role_subjects` is the join table
	TenantResourceActions []TenantResourceAction `gorm:"many2many:tenant_role_actions"`  // TenantRole has many TenantResourceActions, `tenant_role_actions` is the join table
	CreatedAt             int
	UpdatedAt             int
}

type TenantResource struct {
	gorm.Model
	Name                  string
	TenantID              uint
	TenantResourceActions []TenantResourceAction // TenantResource has many TenantResourceActions, TenantID is the foreign key
	CreatedAt             int
	UpdatedAt             int
}

type TenantResourceAction struct {
	gorm.Model
	TenantResourceID uint
	Name             string
	TenantResource   TenantResource
	TenantRoles      []TenantRole `gorm:"many2many:tenant_role_actions"` // TenantResourceAction has many TenantRoles, `tenant_role_actions` is the join table
	CreatedAt        int
	UpdatedAt        int
}

type TenantRoleAction struct {
	gorm.Model
	TenantResourceActionID uint
	TenantRoleID           uint
	CreatedAt              int
	UpdatedAt              int
}

type TenantRoleSubject struct {
	gorm.Model
	TenantSubjectID uint
	TenantRoleID    uint
	CreatedAt       int
	UpdatedAt       int
}
