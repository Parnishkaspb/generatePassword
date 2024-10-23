package generatepassword

import (
	"Project1/internal/pkg/charts"
	"math/rand"
	"strings"
)

type GeneratePasswordImpl struct {
}

func New() *GeneratePasswordImpl {
	return &GeneratePasswordImpl{}
}

func (g *GeneratePasswordImpl) GeneratePassword(includeUpperCase bool, includeNumbers bool, includeSpecialChars bool, length int, repeats int) string {
	mainString := "qwertyuiopasdfghjklzxcvbnm"
	chartSet := charts.New()

	mainString += chartSet.IncludeNumbers(includeNumbers)
	mainString += chartSet.IncludeUpWord(includeUpperCase)
	mainString += chartSet.IncludeSpecialCharts(includeSpecialChars)

	chars := []rune(mainString)
	var rep []string

	for j := 0; j < repeats; j++ {
		var b strings.Builder
		for i := 0; i < length; i++ {
			b.WriteRune(chars[rand.Intn(len(chars))])
		}
		rep = append(rep, b.String())
	}

	return strings.Join(rep, "\n")
}
