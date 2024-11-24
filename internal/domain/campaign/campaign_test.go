package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	name    = "Campaign 1"
	content = "Content 1"
)

var contacts = []Contact{
	{Email: "diego.collares@gmail.com"},
	{Email: "mariane.collares@gmail.com"},
}

func Test_NewCampaign(t *testing.T) {
	assert := assert.New(t)

	campaign := NewCampaign(
		name,
		content,
		contacts,
	)

	println(campaign.ID)
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))

}

func Test_NewCampaign_Id_NotNil(t *testing.T) {
	assert := assert.New(t)

	campaign := NewCampaign(
		name,
		content,
		contacts,
	)

	assert.NotNil(campaign.ID)
}

func Test_NewCampaign_CreatedAt_MostBeNow(t *testing.T) {
	assert := assert.New(t)

	now := time.Now().Add(-time.Minute)

	campaign := NewCampaign(
		name,
		content,
		contacts,
	)

	assert.Greater(campaign.CreatedAt, now)
}
