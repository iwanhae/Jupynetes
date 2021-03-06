// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/iwanhae/Jupynetes/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID string `json:"user_id,omitempty"`
	// UserPw holds the value of the "user_pw" field.
	UserPw string `json:"user_pw,omitempty"`
	// QuotaInstance holds the value of the "quota_instance" field.
	QuotaInstance int `json:"quota_instance,omitempty"`
	// QuotaCPU holds the value of the "quota_cpu" field.
	QuotaCPU int `json:"quota_cpu,omitempty"`
	// QuotaMemory holds the value of the "quota_memory" field.
	QuotaMemory int `json:"quota_memory,omitempty"`
	// QuotaNvidiaGpu holds the value of the "quota_nvidia_gpu" field.
	QuotaNvidiaGpu int `json:"quota_nvidia_gpu,omitempty"`
	// QuotaStorage holds the value of the "quota_storage" field.
	QuotaStorage int `json:"quota_storage,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Events holds the value of the events edge.
	Events []*Event
	// Servers holds the value of the servers edge.
	Servers []*Server
	// Templates holds the value of the templates edge.
	Templates []*Template
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// EventsOrErr returns the Events value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) EventsOrErr() ([]*Event, error) {
	if e.loadedTypes[0] {
		return e.Events, nil
	}
	return nil, &NotLoadedError{edge: "events"}
}

// ServersOrErr returns the Servers value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ServersOrErr() ([]*Server, error) {
	if e.loadedTypes[1] {
		return e.Servers, nil
	}
	return nil, &NotLoadedError{edge: "servers"}
}

// TemplatesOrErr returns the Templates value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) TemplatesOrErr() ([]*Template, error) {
	if e.loadedTypes[2] {
		return e.Templates, nil
	}
	return nil, &NotLoadedError{edge: "templates"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // user_id
		&sql.NullString{}, // user_pw
		&sql.NullInt64{},  // quota_instance
		&sql.NullInt64{},  // quota_cpu
		&sql.NullInt64{},  // quota_memory
		&sql.NullInt64{},  // quota_nvidia_gpu
		&sql.NullInt64{},  // quota_storage
		&sql.NullTime{},   // created_at
		&sql.NullTime{},   // deleted_at
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(values ...interface{}) error {
	if m, n := len(values), len(user.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	u.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field user_id", values[0])
	} else if value.Valid {
		u.UserID = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field user_pw", values[1])
	} else if value.Valid {
		u.UserPw = value.String
	}
	if value, ok := values[2].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field quota_instance", values[2])
	} else if value.Valid {
		u.QuotaInstance = int(value.Int64)
	}
	if value, ok := values[3].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field quota_cpu", values[3])
	} else if value.Valid {
		u.QuotaCPU = int(value.Int64)
	}
	if value, ok := values[4].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field quota_memory", values[4])
	} else if value.Valid {
		u.QuotaMemory = int(value.Int64)
	}
	if value, ok := values[5].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field quota_nvidia_gpu", values[5])
	} else if value.Valid {
		u.QuotaNvidiaGpu = int(value.Int64)
	}
	if value, ok := values[6].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field quota_storage", values[6])
	} else if value.Valid {
		u.QuotaStorage = int(value.Int64)
	}
	if value, ok := values[7].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[7])
	} else if value.Valid {
		u.CreatedAt = value.Time
	}
	if value, ok := values[8].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field deleted_at", values[8])
	} else if value.Valid {
		u.DeletedAt = value.Time
	}
	return nil
}

// QueryEvents queries the events edge of the User.
func (u *User) QueryEvents() *EventQuery {
	return (&UserClient{config: u.config}).QueryEvents(u)
}

// QueryServers queries the servers edge of the User.
func (u *User) QueryServers() *ServerQuery {
	return (&UserClient{config: u.config}).QueryServers(u)
}

// QueryTemplates queries the templates edge of the User.
func (u *User) QueryTemplates() *TemplateQuery {
	return (&UserClient{config: u.config}).QueryTemplates(u)
}

// Update returns a builder for updating this User.
// Note that, you need to call User.Unwrap() before calling this method, if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", user_id=")
	builder.WriteString(u.UserID)
	builder.WriteString(", user_pw=")
	builder.WriteString(u.UserPw)
	builder.WriteString(", quota_instance=")
	builder.WriteString(fmt.Sprintf("%v", u.QuotaInstance))
	builder.WriteString(", quota_cpu=")
	builder.WriteString(fmt.Sprintf("%v", u.QuotaCPU))
	builder.WriteString(", quota_memory=")
	builder.WriteString(fmt.Sprintf("%v", u.QuotaMemory))
	builder.WriteString(", quota_nvidia_gpu=")
	builder.WriteString(fmt.Sprintf("%v", u.QuotaNvidiaGpu))
	builder.WriteString(", quota_storage=")
	builder.WriteString(fmt.Sprintf("%v", u.QuotaStorage))
	builder.WriteString(", created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", deleted_at=")
	builder.WriteString(u.DeletedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
