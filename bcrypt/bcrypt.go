package bcrypt

const (
	MIN_COST     = 4
	DEFAULT_COST = 10
	MAX_COST     = 31
)

// https://en.wanweibaike.com/wiki-BCrypt
func HashPassword(password string, salt string) string {

}

// https://en.wanweibaike.com/wiki-BCrypt
func CheckPassword() {

}

func GenDefaultSalt() string {
	return GenSalt(DEFAULT_COST)
}
func GenSalt(cost int) string {
	return ""
}
