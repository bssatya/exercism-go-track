// Package leap impliments a function to check the leap year
package leap

// IsLeapYear take year as input argument and retuns a bool confirmed whether the year is leap or not.
func IsLeapYear(year int) bool {
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		return true
	}
	return false
}
