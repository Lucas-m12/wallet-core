package entity_test

import (
	"testing"

	entity "github.com/lucas-m12/wallet-core/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := entity.NewClient("John Doe", "johnDoe@example.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John Doe", client.Name)
	assert.Equal(t, "johnDoe@example.com", client.Email)
	assert.NotEmpty(t, client.ID)
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := entity.NewClient("", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}
