package yokoy

import "time"

func parseTime(s *string) *time.Time {
	if s == nil {
		return nil
	}

	t, err := time.Parse("2006-01-02T15:04", *s)

	if err != nil {
		return nil
	}

	return &t
}
