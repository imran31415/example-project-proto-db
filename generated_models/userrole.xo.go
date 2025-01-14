package generated_models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"
)

// UserRole represents a row from 'UserRole'.
type UserRole struct {
	UserID     int       `json:"user_id"`     // user_id
	RoleID     int       `json:"role_id"`     // role_id
	AssignedAt time.Time `json:"assigned_at"` // assigned_at
}

// UserRoleKeysetPage retrieves a page of [UserRole] records using keyset pagination with dynamic filtering.
//
// The keyset pagination retrieves results after or before a specific value (`key`)
// for a given column (`column`) with a limit (`limit`) and order (`ASC` or `DESC`).
//
// If `order` is `ASC`, it retrieves records where the value of `column` is greater than `key`.
// If `order` is `DESC`, it retrieves records where the value of `column` is less than `key`.
//
// Filters are dynamically provided via a `filters` map, where keys are column names and values are either single values or slices for `IN` clauses.
func UserRoleKeysetPage(ctx context.Context, db DB, column string, key interface{}, limit int, order string, filters map[string]interface{}) ([]*UserRole, *UserRole, error) {
	if order != "ASC" && order != "DESC" {
		return nil, nil, fmt.Errorf("invalid order: %s", order)
	}

	// Start building the query
	query := fmt.Sprintf(
		`SELECT * FROM UserRole 
         WHERE %s %s ?`,
		column, condition(order),
	)

	// Arguments for the query
	args := []interface{}{key}

	// Dynamically add filters from the `filters` map to the query
	for field, value := range filters {
		switch v := value.(type) {
		case []int:
			if len(v) > 0 {
				placeholders := make([]string, len(v))
				for i := range v {
					placeholders[i] = "?"
					args = append(args, v[i])
				}
				query += fmt.Sprintf(" AND %s IN (%s)", field, strings.Join(placeholders, ", "))
			}
		case []string:
			if len(v) > 0 {
				placeholders := make([]string, len(v))
				for i := range v {
					placeholders[i] = "?"
					args = append(args, v[i])
				}
				query += fmt.Sprintf(" AND %s IN (%s)", field, strings.Join(placeholders, ", "))
			}
		default:
			// Handle NULL and NOT NULL checks
			if value == nil {
				query += fmt.Sprintf(" AND %s IS NULL", field)
			} else if value == "NOT NULL" {
				query += fmt.Sprintf(" AND %s IS NOT NULL", field)
			} else {
				query += fmt.Sprintf(" AND %s = ?", field)
				args = append(args, value)
			}
		}
	}

	// Finalize the query with the order and limit
	query += fmt.Sprintf(" ORDER BY %s %s LIMIT ?", column, order)
	args = append(args, limit)

	// Log the final query for debugging purposes
	log.Printf("Executing query: %s with args: %v", query, args)

	// Execute the query
	rows, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, nil, logerror(err)
	}
	defer rows.Close()

	var results []*UserRole
	var lastItem *UserRole // Variable to store the last item

	for rows.Next() {
		ur := UserRole{}
		if err := rows.Scan(
			&ur.UserID, &ur.RoleID, &ur.AssignedAt,
		); err != nil {
			return nil, nil, logerror(err)
		}
		results = append(results, &ur)
	}

	// Check for errors during row iteration.
	if err := rows.Err(); err != nil {
		return nil, nil, logerror(err)
	}

	// If we have results, set the lastItem to the last element in results.
	if len(results) > 0 {
		lastItem = results[len(results)-1]
	}

	return results, lastItem, nil
}

// UserRoleByRoleID retrieves a row from 'UserRole' as a [UserRole].
//
// Generated from index 'role_id'.
func UserRoleByRoleID(ctx context.Context, db DB, roleID int) ([]*UserRole, error) {
	// query
	const sqlstr = `SELECT ` +
		`user_id, role_id, assigned_at ` +
		`FROM UserRole ` +
		`WHERE role_id = ?`
	// run
	logf(sqlstr, roleID)
	rows, err := db.QueryContext(ctx, sqlstr, roleID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*UserRole
	for rows.Next() {
		ur := UserRole{}
		// scan
		if err := rows.Scan(&ur.UserID, &ur.RoleID, &ur.AssignedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &ur)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// UserRoleByUserID retrieves a row from 'UserRole' as a [UserRole].
//
// Generated from index 'user_id'.
func UserRoleByUserID(ctx context.Context, db DB, userID int) ([]*UserRole, error) {
	// query
	const sqlstr = `SELECT ` +
		`user_id, role_id, assigned_at ` +
		`FROM UserRole ` +
		`WHERE user_id = ?`
	// run
	logf(sqlstr, userID)
	rows, err := db.QueryContext(ctx, sqlstr, userID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*UserRole
	for rows.Next() {
		ur := UserRole{}
		// scan
		if err := rows.Scan(&ur.UserID, &ur.RoleID, &ur.AssignedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &ur)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// User returns the User associated with the [UserRole]'s (UserID).
//
// Generated from foreign key 'userrole_ibfk_1'.
func (ur *UserRole) User(ctx context.Context, db DB) (*User, error) {
	return UserByUserID(ctx, db, ur.UserID)
}

// Role returns the Role associated with the [UserRole]'s (RoleID).
//
// Generated from foreign key 'userrole_ibfk_2'.
func (ur *UserRole) Role(ctx context.Context, db DB) (*Role, error) {
	return RoleByRoleID(ctx, db, ur.RoleID)
}
