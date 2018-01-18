package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/markbates/pop"
	"github.com/mattstratton/oathkeeper/models"
	"github.com/pkg/errors"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Talk)
// DB Table: Plural (talks)
// Resource: Plural (Talks)
// Path: Plural (/talks)
// View Template Folder: Plural (/templates/talks/)

// TalksResource is the resource for the Talk model
type TalksResource struct {
	buffalo.Resource
}

// List gets all Talks. This function is mapped to the path
// GET /talks
func (v TalksResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	talks := &models.Talks{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Talks from the DB
	if err := q.All(talks); err != nil {
		return errors.WithStack(err)
	}

	// Make Talks available inside the html template
	c.Set("talks", talks)

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.HTML("talks/index.html"))
}

// Show gets the data for one Talk. This function is mapped to
// the path GET /talks/{talk_id}
func (v TalksResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Talk
	talk := &models.Talk{}

	// To find the Talk the parameter talk_id is used.
	if err := tx.Find(talk, c.Param("talk_id")); err != nil {
		return c.Error(404, err)
	}

	// Make talk available inside the html template
	c.Set("talk", talk)

	return c.Render(200, r.HTML("talks/show.html"))
}

// New renders the form for creating a new Talk.
// This function is mapped to the path GET /talks/new
func (v TalksResource) New(c buffalo.Context) error {
	// Make talk available inside the html template
	c.Set("talk", &models.Talk{})

	return c.Render(200, r.HTML("talks/new.html"))
}

// Create adds a Talk to the DB. This function is mapped to the
// path POST /talks
func (v TalksResource) Create(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	user := c.Value("current_user").(*models.User)
	// Allocate an empty Talk
	talk := &models.Talk{}

	// Bind link to the html form elements
	err := c.Bind(talk)
	if err != nil {
		return errors.WithStack(err)
	}

	talk.UserID = user.ID

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(talk)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make talk available inside the html template
		c.Set("talk", talk)

		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.HTML("talks/new.html"))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Talk was created successfully")

	// and redirect to the talks index page
	return c.Redirect(302, "/talks/%s", talk.ID)
}

// Edit renders a edit form for a Talk. This function is
// mapped to the path GET /talks/{talk_id}/edit
func (v TalksResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Talk
	talk := &models.Talk{}

	if err := tx.Find(talk, c.Param("talk_id")); err != nil {
		return c.Error(404, err)
	}

	// Make talk available inside the html template
	c.Set("talk", talk)
	return c.Render(200, r.HTML("talks/edit.html"))
}

// Update changes a Talk in the DB. This function is mapped to
// the path PUT /talks/{talk_id}
func (v TalksResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Talk
	talk := &models.Talk{}

	if err := tx.Find(talk, c.Param("talk_id")); err != nil {
		return c.Error(404, err)
	}

	// Bind Talk to the html form elements
	if err := c.Bind(talk); err != nil {
		return errors.WithStack(err)
	}

	verrs, err := tx.ValidateAndUpdate(talk)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make talk available inside the html template
		c.Set("talk", talk)

		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.HTML("talks/edit.html"))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Talk was updated successfully")

	// and redirect to the talks index page
	return c.Redirect(302, "/talks/%s", talk.ID)
}

// Destroy deletes a Talk from the DB. This function is mapped
// to the path DELETE /talks/{talk_id}
func (v TalksResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Talk
	talk := &models.Talk{}

	// To find the Talk the parameter talk_id is used.
	if err := tx.Find(talk, c.Param("talk_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(talk); err != nil {
		return errors.WithStack(err)
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", "Talk was destroyed successfully")

	// Redirect to the talks index page
	return c.Redirect(302, "/talks")
}