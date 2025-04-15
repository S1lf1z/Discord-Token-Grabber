package utils.go

import "regexp"

var TokenRe = regexp.MustCompile("[\\w-]{24}\\.[\\w-]{6}\\.[\\w-]{27}|mfa\\.[\\w-]{84}")
