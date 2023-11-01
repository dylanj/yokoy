package yokoy

import "time"

func parseLongTime(s *string) *time.Time {
	if s == nil {
		return nil
	}

	t, err := time.Parse("Mon, 02 Jan 2006 03:04:05 MST", *s)

	if err != nil {
		return nil
	}

	return &t
}

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

func parseDate(s *string) *time.Time {
	if s == nil {
		return nil
	}

	t, err := time.Parse("2006-01-02", *s)

	if err != nil {
		return nil
	}

	return &t
}
