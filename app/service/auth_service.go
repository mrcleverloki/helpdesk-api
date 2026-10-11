package service

import (
	"time"

	"helpdesk-api/app/model"
	"helpdesk-api/app/repository"
	"helpdesk-api/helper"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
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
	var req model.CreateUserRequest
	if err := helper.ParseAndValidate(c, &req); err != nil {
		return err
	}
	return helper.Success(c, fiber.StatusCreated, "registrasi berhasil", req)
}

func (s *AuthService) Login(c *fiber.Ctx) error {
	var req model.LoginRequest
	if err := helper.ParseAndValidate(c, &req); err != nil {
		return err
	}

	user, err := s.userRepo.FindByUsername(c.Context(), req.Username)
	if err != nil {
		return helper.Unauthorized("username atau password salah")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return helper.Unauthorized("username atau password salah")
	}

	accessToken, err := s.jwtManager.GenerateAccess(user)
	if err != nil {
		return helper.Internal(err)
	}

	return helper.Success(c, fiber.StatusOK, "login berhasil", fiber.Map{
		"access_token": accessToken,
		"token_type":   "Bearer",
		"user":         user,
	})
}

func (s *AuthService) Refresh(c *fiber.Ctx) error {
	return helper.Success(c, fiber.StatusOK, "refresh token endpoint", nil)
}

func (s *AuthService) Logout(c *fiber.Ctx) error {
	return helper.Success(c, fiber.StatusOK, "logout berhasil", nil)
}

func (s *AuthService) Me(c *fiber.Ctx) error {
	user, ok := helper.CurrentUser(c)
	if !ok {
		return helper.Unauthorized("belum terautentikasi")
	}

	return helper.Success(c, fiber.StatusOK, "profil berhasil diambil", fiber.Map{
		"user":        user,
		"permissions": s.perms.PermissionsOf(user.Role),
	})
}
