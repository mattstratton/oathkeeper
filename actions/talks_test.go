package actions

import "github.com/mattstratton/oathkeeper/models"

func (as *ActionSuite) Test_TalksResource_List() {
	user := as.Login()
	talk := as.CreateTalk(user)

	res := as.HTML("/talks").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), talk.Title)
}

func (as *ActionSuite) Test_TalksResource_Show() {
	user := as.Login()
	talk := as.CreateTalk(user)

	res := as.HTML("/talks/%s", talk.ID).Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), talk.Title)
}

func (as *ActionSuite) Test_TalksResource_New() {
	as.Login()

	res := as.HTML("/talks/new").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "New Talk")
}

func (as *ActionSuite) Test_TalksResource_Create() {
	user := as.Login()
	talk := &models.Talk{
		Title:    "Test Talk",
		Abstract: "My abstract",
	}

	res := as.HTML("/talks").Post(talk)
	as.Equal(302, res.Code)

	t := &models.Talk{}
	as.NoError(as.DB.First(t))
	as.Equal(talk.Title, t.Title)
	as.Equal(user.ID, t.UserID)
}

func (as *ActionSuite) Test_TalksResource_Edit() {
	user := as.Login()
	talk := as.CreateTalk(user)

	res := as.HTML("/talks/%s/edit", talk.ID).Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "Edit Talk")
}

func (as *ActionSuite) Test_TalksResource_Update() {
	user := as.Login()
	talk := as.CreateTalk(user)

	pt := talk.Title

	talk.Title = "Another talk"

	res := as.HTML("/talks/%s", talk.ID).Put(talk)
	as.Equal(302, res.Code)

	t := &models.Talk{}
	as.NoError(as.DB.First(t))
	as.NotEqual(pt, t.Title)
	as.Equal(user.ID, t.UserID)

}

func (as *ActionSuite) Test_TalksResource_Destroy() {
	user := as.Login()
	talk := as.CreateTalk(user)

	res := as.HTML("/talks/%s", talk.ID).Delete()
	as.Equal(302, res.Code)
	as.Equal("/talks", res.Location())
}
