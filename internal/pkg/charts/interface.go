package charts

type Charts interface {
	IncludeNumbers(bool) string
	IncludeUpWord(bool) string
	IncludeSpecialCharts(bool) string
}
