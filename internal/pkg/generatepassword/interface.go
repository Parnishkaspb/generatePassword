package generatepassword

type GeneratePassword interface {
	GeneratePassword(bool, bool, bool, int, int) string
}
