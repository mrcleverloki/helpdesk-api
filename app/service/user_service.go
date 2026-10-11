package service

import (
	"strconv"

	"helpdesk-api/app/model"
	"helpdesk-api/app/repository"
	"helpdesk-api/helper"

	"github.com/gofiber/fiber/v2"
)

type UserService struct {
	repo  repository.UserRepository
	perms *helper.PermissionSet
}

func NewUserService(
	repo repository.UserRepository,
	perms *helper.PermissionSet,
) *UserService {
	return &UserService{repo: repo, perms: perms}
}

func (s *UserService) List(c *fiber.Ctx) error {
	var query model.CursorQuery
	if err := c.QueryParser(&query); err != nil {
		return helper.BadRequest("parameter query tidak valid")
	}

	if query.Limit <= 0 {
		query.Limit = 10
	}

	var cursor model.Cursor
	if query.Cursor != "" {
		decoded, err := helper.DecodeCursor(query.Cursor)
		if err != nil {
			return helper.BadRequest("cursor tidak valid")
		}
		cursor = decoded
	}

	users, err := s.repo.FindAllCursor(c.Context(), query, cursor)
	if err != nil {
		return helper.Internal(err)
	}

	var nextCursor string
	hasMore := false

	if len(users) > query.Limit {
		hasMore = true
		users = users[:query.Limit] // Potong kelebihan 1 data
		lastUser := users[len(users)-1]
		nextCursor = helper.EncodeCursor(lastUser.CreatedAt, lastUser.ID)
	}

	return helper.SuccessWithMeta(c, fiber.StatusOK, "daftar user berhasil diambil", users, fiber.Map{
		"next_cursor": nextCursor,
		"has_more":    hasMore,
		"limit":       query.Limit,
	})
}

func (s *UserService) Create(c *fiber.Ctx) error {
	var req model.CreateUserRequest
	if err := helper.ParseAndValidate(c, &req); err != nil {
		return err
	}

	return helper.Success(c, fiber.StatusCreated, "user berhasil dibuat", req)
}

func (s *UserService) Get(c *fiber.Ctx) error {
	currentUser, ok := helper.CurrentUser(c)
	if !ok {
		return helper.Unauthorized("belum terautentikasi")
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return helper.BadRequest("ID user tidak valid")
	}

	if !CanAccessUser(currentUser, id, s.perms, "user:read:any") {
		return helper.Forbidden("akses ditolak ke data user lain")
	}

	user, err := s.repo.FindByID(c.Context(), id)
	if err != nil {
		return helper.NotFound("user tidak ditemukan")
	}

	return helper.Success(c, fiber.StatusOK, "user berhasil diambil", user)
}

func (s *UserService) Replace(c *fiber.Ctx) error {
	currentUser, ok := helper.CurrentUser(c)
	if !ok {
		return helper.Unauthorized("belum terautentikasi")
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return helper.BadRequest("ID user tidak valid")
	}

	if !CanAccessUser(currentUser, id, s.perms, "user:update:any") {
		return helper.Forbidden("akses ditolak mengubah data user lain")
	}

	var req model.ReplaceUserRequest
	if err := helper.ParseAndValidate(c, &req); err != nil {
		return err
	}

	return helper.Success(c, fiber.StatusOK, "user berhasil diperbarui", req)
}

func (s *UserService) Patch(c *fiber.Ctx) error {
	currentUser, ok := helper.CurrentUser(c)
	if !ok {
		return helper.Unauthorized("belum terautentikasi")
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return helper.BadRequest("ID user tidak valid")
	}

	if !CanAccessUser(currentUser, id, s.perms, "user:update:any") {
		return helper.Forbidden("akses ditolak mengubah data user lain")
	}

	var req model.PatchUserRequest
	if err := helper.ParseAndValidate(c, &req); err != nil {
		return err
	}

	return helper.Success(c, fiber.StatusOK, "user berhasil di-patch", req)
}

func (s *UserService) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return helper.BadRequest("ID user tidak valid")
	}

	return helper.Success(c, fiber.StatusOK, "user berhasil dihapus", fiber.Map{"id": id})
}

func (s *UserService) AssignRole(c *fiber.Ctx) error {
	currentUser, ok := helper.CurrentUser(c)
	if !ok {
		return helper.Unauthorized("belum terautentikasi")
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return helper.BadRequest("ID user tidak valid")
	}

	var req model.AssignRoleRequest
	if err := helper.ParseAndValidate(c, &req); err != nil {
		return err
	}

	if errs := ValidateAssignRole(currentUser, id, req, s.perms); len(errs) > 0 {
		return helper.Validation(errs)
	}

	return helper.Success(c, fiber.StatusOK, "role berhasil diubah", fiber.Map{"target_id": id, "role": req.Role})
}
