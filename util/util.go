package util

func StrPtr(s string) *string {
	return &s
}

func NonEmptyStringPtrOr(s *string, alt string) string {
	if s == nil || *s == "" {
		return alt
	}
	return *s
}
