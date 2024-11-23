package campaign

import "testing"

func TestNewCampaign(t *testing.T) {
	name := "Campaign 1"
	content := "Content 1"
	contacts := []Contact{
		{Email: "diego.collares@gmail.com"},
		{Email: "mariane.collares@gmail.com"},
	}

	campaign := NewCampaign(
		name,
		content,
		contacts,
	)

	if campaign.ID != "1" {
		t.Errorf("Expected ID to be 1, but got %s", campaign.ID)
	}

	if campaign.Name != name {
		t.Errorf("Expected Name to be %s, but got %s", name, campaign.Name)
	}

	if campaign.Content != content {
		t.Errorf("Expected Content to be %s, but got %s", content, campaign.Content)
	}

	if len(campaign.Contacts) != 2 {
		t.Errorf("Expected 2 contacts, but got %d", len(campaign.Contacts))
	}

}
