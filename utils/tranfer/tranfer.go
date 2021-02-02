package tranfer

import "time"

func DeferString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

func DeferInterface(s *interface{}) interface{}{
	if s != nil {
		return *s
	}
	return nil
}

func DeferInt(n *int) int {
	if n != nil {
		return *n
	}
	return 0
}

func DeferBool(n *bool) bool {
	if n!= nil {
		return *n
	}
	return false
}

func DeferBoolPonter(n bool) *bool {
	return &n
}

func DeferTimeToString(time *time.Time) *string {
	strTime := time.String()

	return &strTime
}