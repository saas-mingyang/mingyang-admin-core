package utils

import (
	"context"
	"fmt"
	"strconv"
)

// GetUserIdFromCtx extracts userId from context and converts it to uint64
// It handles different types that might be stored in the context (string, uint64, int64, int, etc.)
func GetUserIdFromCtx(ctx context.Context) uint64 {
	if userIdVal := ctx.Value("userId"); userIdVal != nil {
		switch v := userIdVal.(type) {
		case string:
			if id, err := strconv.ParseUint(v, 10, 64); err == nil {
				return id
			}
		case uint64:
			return v
		case int64:
			return uint64(v)
		case int:
			return uint64(v)
		default:
			// Try to convert to string then parse
			if id, err := strconv.ParseUint(fmt.Sprintf("%v", v), 10, 64); err == nil {
				return id
			}
		}
	}
	return 0
}

// GetRoleIdFromCtx extracts roleId from context and returns it as string
// It handles different types that might be stored in the context
func GetRoleIdFromCtx(ctx context.Context) string {
	if roleIdVal := ctx.Value("roleId"); roleIdVal != nil {
		switch v := roleIdVal.(type) {
		case string:
			return v
		default:
			return fmt.Sprintf("%v", v)
		}
	}
	return ""
}

// GetDeptIdFromCtx extracts deptId from context and converts it to uint64
// It handles different types that might be stored in the context
func GetDeptIdFromCtx(ctx context.Context) uint64 {
	if deptIdVal := ctx.Value("deptId"); deptIdVal != nil {
		switch v := deptIdVal.(type) {
		case string:
			if id, err := strconv.ParseUint(v, 10, 64); err == nil {
				return id
			}
		case uint64:
			return v
		case int64:
			return uint64(v)
		case int:
			return uint64(v)
		default:
			// Try to convert to string then parse
			if id, err := strconv.ParseUint(fmt.Sprintf("%v", v), 10, 64); err == nil {
				return id
			}
		}
	}
	return 0
}
