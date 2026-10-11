package service

import (
	"time"

	"helpdesk-api/app/repository"
	"helpdesk-api/helper"

	"github.com/gofiber/fiber/v2"
)

type AuthService struct {
	userRepo   repository.UserRepository
	tokenRepo  repository.TokenRepository
	jwtManager *helper.JWTManager
	perms      *helper.PermissionSet
	refreshTTL time.Duration
}

func NewAuthService(
	userRepo repository.UserRepository,
	tokenRepo repository.TokenRepository,
	jwtManager *helper.JWTManager,
	perms *helper.PermissionSet,
	refreshTTL time.Duration,
) *AuthService {
	return &AuthService{
		userRepo:   userRepo,
		tokenRepo:  tokenRepo,
		jwtManager: jwtManager,
		perms:      perms,
		refreshTTL: refreshTTL,
	}
}

func (s *AuthService) Register(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "register endpoint"})
}

func (s *AuthService) Login(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "login endpoint"})
}

func (s *AuthService) Refresh(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "refresh endpoint"})
}

func (s *AuthService) Logout(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "logout endpoint"})
}

// Me menampilkan profil user beserta daftar permission-nya (Langkah 8)[cite: 1]
func (s *AuthService) Me(c *fiber.Ctx) error {
	user, ok := helper.CurrentUser(c)
	if !ok {
		return helper.Fail(c, fiber.StatusUnauthorized, "belum terautentikasi")
	}

	return helper.Success(c, fiber.StatusOK, "profil berhasil diambil", fiber.Map{
		"user":        user,
		"permissions": s.perms.PermissionsOf(user.Role),
	})
}
