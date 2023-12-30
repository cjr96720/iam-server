BEGIN;
-- one application can have many tenants and subject
-- applications.id == tenants.application_id
-- applications.id == subjects.application_id
CREATE TABLE applications (
    id SERIAL PRIMARY KEY,
    created_at BIGINT DEFAULT NULL,
    updated_at BIGINT DEFAULT NULL,
    deleted_at TIMESTAMPTZ DEFAULT NULL,
    name TEXT,
    app_id VARCHAR(200) UNIQUE,
    secret TEXT
);
-- one tenant can have many resources
-- tenants.id == tenant_resources.tenant_id
-- one tenant can have many role
-- tenants_id == tenant_roles.tenant_id
CREATE TABLE tenants (
    id SERIAL PRIMARY KEY,
    created_at BIGINT DEFAULT NULL,
    updated_at BIGINT DEFAULT NULL,
    deleted_at TIMESTAMPTZ DEFAULT NULL,
    application_id BIGINT,
    name TEXT
);
CREATE TABLE subjects (
    id SERIAL PRIMARY KEY,
    created_at BIGINT DEFAULT NULL,
    updated_at BIGINT DEFAULT NULL,
    deleted_at TIMESTAMPTZ DEFAULT NULL,
    application_id BIGINT,
    external_id TEXT,
    name TEXT
);
-- establish a many-to-many relationship between tenants and subjects
-- tenants.id == tenant_subjects.tenant_id
-- subjects_id == tenant_subjects.subject_id
CREATE TABLE tenant_subjects (
    id SERIAL PRIMARY KEY,
    created_at BIGINT DEFAULT NULL,
    updated_at BIGINT DEFAULT NULL,
    deleted_at TIMESTAMPTZ DEFAULT NULL,
    tenant_id BIGINT,
    subject_id BIGINT
);
-- one tenant resource can have many actions
-- tenant_resources.id == tenant_resource_actions.tenant_resource_id
CREATE TABLE tenant_resources (
    id SERIAL PRIMARY KEY,
    created_at BIGINT DEFAULT NULL,
    updated_at BIGINT DEFAULT NULL,
    deleted_at TIMESTAMPTZ DEFAULT NULL,
    name TEXT,
    tenant_id BIGINT
);
-- action can be anything a subject in a tenant can do to a resource
-- e.g., read, write, create, ...etc.
CREATE TABLE tenant_resource_actions (
    id SERIAL PRIMARY KEY,
    created_at BIGINT DEFAULT NULL,
    updated_at BIGINT DEFAULT NULL,
    deleted_at TIMESTAMPTZ DEFAULT NULL,
    tenant_resource_id BIGINT,
    name TEXT
);
CREATE TABLE tenant_roles (
    id SERIAL PRIMARY KEY,
    created_at BIGINT DEFAULT NULL,
    updated_at BIGINT DEFAULT NULL,
    deleted_at TIMESTAMPTZ DEFAULT NULL,
    name TEXT,
    tenant_id BIGINT
);
-- establish a many-to-many relationship between tenant resource actions and roles
-- tenant_role_actions.tenant_resource_action_id == tenant_resource_actions.id 
-- tenant_role_actions.tenant_role_id == tenant_roles.id
CREATE TABLE tenant_role_actions (
    id SERIAL PRIMARY KEY,
    created_at BIGINT DEFAULT NULL,
    updated_at BIGINT DEFAULT NULL,
    deleted_at TIMESTAMPTZ DEFAULT NULL,
    tenant_resource_action_id BIGINT,
    tenant_role_id BIGINT
);
-- establish a many-to-many relationship between tenant subjects and roles
-- tenant_role_subjects.tenant_subject_id == tenant_subjects.id
-- tenant_role_subjects.tenant_role_id == tenant_roles.id
CREATE TABLE tenant_role_subjects (
    id SERIAL PRIMARY KEY,
    created_at BIGINT DEFAULT NULL,
    updated_at BIGINT DEFAULT NULL,
    deleted_at TIMESTAMPTZ DEFAULT NULL,
    tenant_subject_id BIGINT,
    tenant_role_id BIGINT
);
COMMIT;