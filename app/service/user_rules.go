package service

import (
	"strings"

	"helpdesk-api/app/model"
)

func ValidateCreate(req model.CreateUserRequest) map[string]string {
	errs := map[string]string{}
	if strings.TrimSpace(req.Username) == "" {
		errs["username"] = "wajib diisi"
	}
	if !isValidEmail(req.Email) {
		errs["email"] = "format email tidak valid"
	}
	if len(req.Password) < 8 {
		errs["password"] = "minimal 8 karakter"
	}
	return errs
}

func ValidateReplace(req model.ReplaceUserRequest) map[string]string {
	errs := map[string]string{}
	if strings.TrimSpace(req.Username) == "" {
		errs["username"] = "wajib diisi"
	}
	if !isValidEmail(req.Email) {
		errs["email"] = "format email tidak valid"
	}
	return errs
}

func IsEmptyPatch(req model.PatchUserRequest) bool {
	return req.Username == nil && req.Email == nil && req.IsActive == nil
}

func ApplyPatch(u model.User, req model.PatchUserRequest) (model.User, map[string]string) {
	errs := map[string]string{}

	if req.Username != nil {
		username := strings.TrimSpace(*req.Username)
		if username == "" {
			errs["username"] = "wajib diisi"
		} else {
			u.Username = username
		}
	}

	if req.Email != nil {
		email := strings.TrimSpace(*req.Email)
		if !isValidEmail(email) {
			errs["email"] = "format email tidak valid"
		} else {
			u.Email = email
		}
	}

	if req.IsActive != nil {
		u.IsActive = *req.IsActive
	}

	return u, errs
}

func CountTotalPages(total, limit int) int {
	if limit <= 0 {
		return 0
	}
	pages := total / limit
	if total%limit != 0 {
		pages++
	}
	return pages
}
