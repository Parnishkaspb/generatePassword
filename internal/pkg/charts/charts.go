package charts

type ChartsImpl struct {
}

func New() *ChartsImpl {
	return &ChartsImpl{}
}

func (c *ChartsImpl) IncludeNumbers(flag bool) string {
	if flag {
		return "0123456789"
	}
	return ""
}

func (c *ChartsImpl) IncludeUpWord(flag bool) string {
	if flag {
		return "QWERTYUIOPASDFGHJKLZXCVBNM"
	}
	return ""
}

func (c *ChartsImpl) IncludeSpecialCharts(flag bool) string {
	if flag {
		return "!@#$%^&*()"
	}
	return ""
}
