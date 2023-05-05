package register

type CreateEmailVerification struct {
	SUBJECT           string
	EMAIL             string
	VERIFICATION_CODE string
}
