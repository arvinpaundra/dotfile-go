package request

import "time"

type (
	Common struct {
		Keyword   string `form:"keyword"`
		Offset    int    `form:"offset"`
		Limit     int    `form:"limit"`
		StartDate string `form:"start_date"`
		EndDate   string `form:"end_date"`
	}
)

func (c Common) GetLimit() int {
	if c.Limit < 10 {
		return 10
	}
	return c.Limit
}

func (c Common) GetOffset() int {
	if c.Offset < 0 {
		return 0
	}
	return c.Offset
}

func (c Common) ValidateStartDate() error {
	_, err := time.Parse(time.DateOnly, c.StartDate)

	return err
}

func (c Common) ValidateEndDate() error {
	_, err := time.Parse(time.DateOnly, c.EndDate)

	return err
}
