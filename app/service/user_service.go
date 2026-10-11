package service

import (
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

// Placeholder method sementara jika belum ada
func (s *UserService) List(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "list users"})
}

func (s *UserService) Create(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "create user"})
}

func (s *UserService) Delete(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "delete user"})
}

func (s *UserService) AssignRole(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "assign role"})
}

func (s *UserService) Get(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "get user by id"})
}

func (s *UserService) Replace(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "replace user"})
}

func (s *UserService) Patch(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "patch user"})
}
