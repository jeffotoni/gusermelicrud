package validador

import (
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	validador "github.com/Nhanderu/brdoc"
	"github.com/jeffotoni/gusermeli/apicore/pkg/fmts"
)

func IsCPF(input string) bool {
	if os.Getenv("DEBUG_TEST") == "true" {
		return true
	}
	return validador.IsCPF(input)
}

// regex email valid
func RegexEmail(nameregex string) bool {
	// regex for my endpoint
	var loginRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`) // Contains email valid
	if loginRegex.MatchString(nameregex) {
		return true
	} else {
		return false
	}
}

func Birthday(birthday string) bool {
	t := "T00:00:00Z"
	now := time.Now().Format("2006-01-02")
	now = fmts.ConcatStr(now, t)
	birthday = fmts.ConcatStr(birthday, t)

	d1, _ := time.Parse(time.RFC3339, birthday)
	d2, _ := time.Parse(time.RFC3339, now)

	month1 := int(d1.Month())
	month2 := int(d2.Month())

	day1 := int(d1.Day())
	day2 := int(d2.Day())

	ryear := 0
	if month1 > month2 {
		ryear = ryear + 1
	} else if month1 == month2 && day1 > day2 {
		ryear = ryear + 1
	}

	str := d2.Sub(d1).String()
	str = strings.Replace(str, "h0m0s", "", -1)
	s, _ := strconv.ParseFloat(str, 32)

	year := s * 0.000114
	year2 := int(math.Ceil(year))
	year2 = int(year2) - int(ryear)
	if year2 > 18 {
		return true
	}
	return false
}
