// Code generated by entc, DO NOT EDIT.

package template

import (
	"time"
)

const (
	// Label holds the string label denoting the template type in the database.
	Label = "template"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldTemplate holds the string denoting the template field in the database.
	FieldTemplate = "template"
	// FieldVariables holds the string denoting the variables field in the database.
	FieldVariables = "variables"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"

	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeServer holds the string denoting the server edge name in mutations.
	EdgeServer = "server"

	// Table holds the table name of the template in the database.
	Table = "templates"
	// UserTable is the table the holds the user relation/edge. The primary key declared below.
	UserTable = "user_templates"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// ServerTable is the table the holds the server relation/edge. The primary key declared below.
	ServerTable = "server_template_from"
	// ServerInverseTable is the table name for the Server entity.
	// It exists in this package in order to avoid circular dependency with the "server" package.
	ServerInverseTable = "servers"
)

// Columns holds all SQL columns for template fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldTemplate,
	FieldVariables,
	FieldCreatedAt,
	FieldDeletedAt,
}

var (
	// UserPrimaryKey and UserColumn2 are the table columns denoting the
	// primary key for the user relation (M2M).
	UserPrimaryKey = []string{"user_id", "template_id"}
	// ServerPrimaryKey and ServerColumn2 are the table columns denoting the
	// primary key for the server relation (M2M).
	ServerPrimaryKey = []string{"server_id", "template_id"}
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
