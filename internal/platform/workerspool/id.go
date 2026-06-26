package workerspool

import (
	"github.com/google/uuid"
)

type IdGenerator interface {
	Generate() string
	Validate(string) bool
}

type GoogleIdGenerator struct {
}

func NewGoogleIdGenerator() *GoogleIdGenerator {
	return &GoogleIdGenerator{}
}

func (g GoogleIdGenerator) Generate() string {
	id, _ := uuid.NewUUID()
	//TODO: HANDLE ERROR

	return id.String()

}

func (g GoogleIdGenerator) Validate(id string) bool {
	err := uuid.Validate(id)
	return err == nil
}
