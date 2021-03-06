// Code generated by entc, DO NOT EDIT.

package server

import (
	"time"
)

const (
	// Label holds the string label denoting the server type in the database.
	Label = "server"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldTemplate holds the string denoting the template field in the database.
	FieldTemplate = "template"
	// FieldVariables holds the string denoting the variables field in the database.
	FieldVariables = "variables"
	// FieldIP holds the string denoting the ip field in the database.
	FieldIP = "ip"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldCPU holds the string denoting the cpu field in the database.
	FieldCPU = "cpu"
	// FieldMemory holds the string denoting the memory field in the database.
	FieldMemory = "memory"
	// FieldNvidiaGpu holds the string denoting the nvidia_gpu field in the database.
	FieldNvidiaGpu = "nvidia_gpu"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"

	// EdgeOwners holds the string denoting the owners edge name in mutations.
	EdgeOwners = "owners"
	// EdgeEvent holds the string denoting the event edge name in mutations.
	EdgeEvent = "event"
	// EdgeTemplateFrom holds the string denoting the template_from edge name in mutations.
	EdgeTemplateFrom = "template_from"

	// Table holds the table name of the server in the database.
	Table = "servers"
	// OwnersTable is the table the holds the owners relation/edge. The primary key declared below.
	OwnersTable = "user_servers"
	// OwnersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnersInverseTable = "users"
	// EventTable is the table the holds the event relation/edge. The primary key declared below.
	EventTable = "event_server"
	// EventInverseTable is the table name for the Event entity.
	// It exists in this package in order to avoid circular dependency with the "event" package.
	EventInverseTable = "events"
	// TemplateFromTable is the table the holds the template_from relation/edge. The primary key declared below.
	TemplateFromTable = "server_template_from"
	// TemplateFromInverseTable is the table name for the Template entity.
	// It exists in this package in order to avoid circular dependency with the "template" package.
	TemplateFromInverseTable = "templates"
)

// Columns holds all SQL columns for server fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldTemplate,
	FieldVariables,
	FieldIP,
	FieldDescription,
	FieldCPU,
	FieldMemory,
	FieldNvidiaGpu,
	FieldCreatedAt,
	FieldDeletedAt,
}

var (
	// OwnersPrimaryKey and OwnersColumn2 are the table columns denoting the
	// primary key for the owners relation (M2M).
	OwnersPrimaryKey = []string{"user_id", "server_id"}
	// EventPrimaryKey and EventColumn2 are the table columns denoting the
	// primary key for the event relation (M2M).
	EventPrimaryKey = []string{"event_id", "server_id"}
	// TemplateFromPrimaryKey and TemplateFromColumn2 are the table columns denoting the
	// primary key for the template_from relation (M2M).
	TemplateFromPrimaryKey = []string{"server_id", "template_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt func() time.Time
)
