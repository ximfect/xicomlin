package xicomlin

import "strings"

// PadLeft pads the given string to the left with spaces. Output is n
// characters long.
func PadLeft(s string, n int) string {
	if len(s) >= n {
		return s
	}

	for len(s) != n {
		s = s + " "
	}

	return s
}

// PadRight pads the given string to the right with spaces. Output is n
// characters long.
func PadRight(s string, n int) string {
	if len(s) >= n {
		return s
	}

	for len(s) != n {
		s = " " + s
	}

	return s
}

// More makes the given string n characters long, adding ellipsis if needed.
func More(s string, n int) string {
	if len(s) >= n {
		return s[:n-3] + "..."
	}
	return s
}

// Linebreak format the given string into lines at most n characters long.
func Linebreak(s string, n int) string {
	ctx := ""
	lines := []string{}

	for _, c := range s {
		ctx += string(c)
		if len(strings.TrimSpace(ctx)) == n-1 {
			lines = append(lines, strings.TrimSpace(ctx))
			ctx = ""
		}
	}
	lines = append(lines, ctx)

	return strings.Join(lines, "\n")
}
