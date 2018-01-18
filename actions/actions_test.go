package actions

import (
	"testing"

	"github.com/gobuffalo/suite"
	"github.com/markbates/pop/nulls"
	"github.com/mattstratton/oathkeeper/models"
)

type ActionSuite struct {
	*suite.Action
}

func Test_ActionSuite(t *testing.T) {
	as := &ActionSuite{suite.NewAction(App())}
	suite.Run(t, as)
}

func (as *ActionSuite) Login() *models.User {
	user := &models.User{
		Name:       "Matt",
		Email:      nulls.NewString("matt@example.com"),
		Provider:   "faux",
		ProviderID: "123",
	}
	as.NoError(as.DB.Create(user))
	as.Session.Set("current_user_id", user.ID)
	return user
}

func (as *ActionSuite) CreateTalk(user *models.User) *models.Talk {
	talk := &models.Talk{
		Title:    "Example Talk",
		Abstract: "Hello World",
		UserID:   user.ID,
	}

	verrs, err := as.DB.ValidateAndCreate(talk)
	as.NoError(err)
	as.False(verrs.HasAny())
	return talk
}
