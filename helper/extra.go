package helper

import (
	"github.com/gofiber/fiber/v2"
	"helpdesk-api/app/model"
)

// CalculateMeta menghitung informasi pagination untuk metadata.
func CalculateMeta(page, perPage, totalData int) *model.Meta {
	if page <= 0 {
		page = 1
	}
	if perPage <= 0 {
		perPage = 10
	}

	totalPage := totalData / perPage
	if totalData%perPage != 0 {
		totalPage++
	}

	return &model.Meta{
		Page:      page,
		PerPage:   perPage,
		TotalData: totalData,
		TotalPage: totalPage,
	}
}

// SuccessWithMeta mengirim respons sukses beserta data dan meta pagination.
func SuccessWithMeta(c *fiber.Ctx, status int, message string, data interface{}, meta *model.Meta) error {
	return c.Status(status).JSON(model.WebResponse{
		Success: true,
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}

// ParseAndValidate melakukan parsing body request sekaligus validasi dasar.
func ParseAndValidate(c *fiber.Ctx, out interface{}) error {
	if err := c.BodyParser(out); err != nil {
		return Fail(c, fiber.StatusBadRequest, "format body request tidak valid")
	}
	return nil
}

// FailWithErrors mengirim respons gagal beserta peta error validasi.
func FailWithErrors(c *fiber.Ctx, status int, message string, errors map[string]string) error {
	return c.Status(status).JSON(model.WebResponse{
		Success: false,
		Message: message,
		Errors:  errors,
	})
}
