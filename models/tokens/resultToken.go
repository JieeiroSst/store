package tokens

type ResultToken struct {
	Token            string
	PayLoad          PayLoad
	RefreshToken     string
	RefreshExpressIn string
}

type PayLoad struct {
	Id int
	Username string
}
