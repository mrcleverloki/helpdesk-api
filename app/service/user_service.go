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
	var query model.ListQuery
	if err := c.QueryParser(&query); err != nil {
		return helper.Fail(c, fiber.StatusBadRequest, "parameter query tidak valid")
	}

	users, total, err := s.repo.FindAll(c.Context(), query)
	if err != nil {
		return helper.Fail(c, fiber.StatusInternalServerError, "gagal mengambil daftar user")
	}

	meta := helper.CalculateMeta(query.Page, query.PerPage, int(total))
	return helper.SuccessWithMeta(c, fiber.StatusOK, "daftar user berhasil diambil", users, meta)
}

func (s *UserService) Create(c *fiber.Ctx) error {
	var req model.CreateUserRequest
	if err := helper.ParseAndValidate(c, &req); err != nil {
		return err
	}

	// Logic create user...
	return helper.Success(c, fiber.StatusCreated, "user berhasil dibuat", req)
}

func (s *UserService) Get(c *fiber.Ctx) error {
	currentUser, ok := helper.CurrentUser(c)
	if !ok {
		return helper.Fail(c, fiber.StatusUnauthorized, "belum terautentikasi")
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return helper.Fail(c, fiber.StatusBadRequest, "ID user tidak valid")
	}

	// Validasi kepemilikan atau hak akses menyeluruh (user:read:any)
	if !CanAccessUser(currentUser, id, s.perms, "user:read:any") {
		return helper.Fail(c, fiber.StatusForbidden, "akses ditolak ke data user lain")
	}

	user, err := s.repo.FindByID(c.Context(), id)
	if err != nil {
		return helper.Fail(c, fiber.StatusNotFound, "user tidak ditemukan")
	}

	return helper.Success(c, fiber.StatusOK, "user berhasil diambil", user)
}

func (s *UserService) Replace(c *fiber.Ctx) error {
	currentUser, ok := helper.CurrentUser(c)
	if !ok {
		return helper.Fail(c, fiber.StatusUnauthorized, "belum terautentikasi")
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return helper.Fail(c, fiber.StatusBadRequest, "ID user tidak valid")
	}

	if !CanAccessUser(currentUser, id, s.perms, "user:update:any") {
		return helper.Fail(c, fiber.StatusForbidden, "akses ditolak mengubah data user lain")
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
		return helper.Fail(c, fiber.StatusUnauthorized, "belum terautentikasi")
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return helper.Fail(c, fiber.StatusBadRequest, "ID user tidak valid")
	}

	if !CanAccessUser(currentUser, id, s.perms, "user:update:any") {
		return helper.Fail(c, fiber.StatusForbidden, "akses ditolak mengubah data user lain")
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
		return helper.Fail(c, fiber.StatusBadRequest, "ID user tidak valid")
	}

	return helper.Success(c, fiber.StatusOK, "user berhasil dihapus", fiber.Map{"id": id})
}

func (s *UserService) AssignRole(c *fiber.Ctx) error {
	currentUser, ok := helper.CurrentUser(c)
	if !ok {
		return helper.Fail(c, fiber.StatusUnauthorized, "belum terautentikasi")
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return helper.Fail(c, fiber.StatusBadRequest, "ID user tidak valid")
	}

	var req model.AssignRoleRequest
	if err := helper.ParseAndValidate(c, &req); err != nil {
		return err
	}

	// Validasi aturan role assignment (tidak boleh ubah role sendiri & role harus valid)
	if errs := ValidateAssignRole(currentUser, id, req, s.perms); len(errs) > 0 {
		return helper.FailWithErrors(c, fiber.StatusUnprocessableEntity, "validasi gagal", errs)
	}

	return helper.Success(c, fiber.StatusOK, "role berhasil diubah", fiber.Map{"target_id": id, "role": req.Role})
}
