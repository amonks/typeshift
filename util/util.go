package util

func NonEmptyStringPtrOr(s *string, alt string) string {
	if s == nil || *s == "" {
		return alt
	}
	return *s
}
