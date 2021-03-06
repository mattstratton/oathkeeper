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
// Model: Singular (Speaker)
// DB Table: Plural (speakers)
// Resource: Plural (Speakers)
// Path: Plural (/speakers)
// View Template Folder: Plural (/templates/speakers/)

// SpeakersResource is the resource for the Speaker model
type SpeakersResource struct {
	buffalo.Resource
}

// List gets all Speakers. This function is mapped to the path
// GET /speakers
func (v SpeakersResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	speakers := &models.Speakers{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Speakers from the DB
	if err := q.All(speakers); err != nil {
		return errors.WithStack(err)
	}

	// Make Speakers available inside the html template
	c.Set("speakers", speakers)

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.HTML("speakers/index.html"))
}

// Show gets the data for one Speaker. This function is mapped to
// the path GET /speakers/{speaker_id}
func (v SpeakersResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Speaker
	speaker := &models.Speaker{}

	// To find the Speaker the parameter speaker_id is used.
	if err := tx.Find(speaker, c.Param("speaker_id")); err != nil {
		return c.Error(404, err)
	}

	// Make speaker available inside the html template
	c.Set("speaker", speaker)

	return c.Render(200, r.HTML("speakers/show.html"))
}

// New renders the form for creating a new Speaker.
// This function is mapped to the path GET /speakers/new
func (v SpeakersResource) New(c buffalo.Context) error {
	// Make speaker available inside the html template
	c.Set("speaker", &models.Speaker{})

	return c.Render(200, r.HTML("speakers/new.html"))
}

// Create adds a Speaker to the DB. This function is mapped to the
// path POST /speakers
func (v SpeakersResource) Create(c buffalo.Context) error {
	// Allocate an empty Speaker
	speaker := &models.Speaker{}

	// Bind speaker to the html form elements
	if err := c.Bind(speaker); err != nil {
		return errors.WithStack(err)
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(speaker)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make speaker available inside the html template
		c.Set("speaker", speaker)

		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.HTML("speakers/new.html"))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Speaker was created successfully")

	// and redirect to the speakers index page
	return c.Redirect(302, "/speakers/%s", speaker.ID)
}

// Edit renders a edit form for a Speaker. This function is
// mapped to the path GET /speakers/{speaker_id}/edit
func (v SpeakersResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Speaker
	speaker := &models.Speaker{}

	if err := tx.Find(speaker, c.Param("speaker_id")); err != nil {
		return c.Error(404, err)
	}

	// Make speaker available inside the html template
	c.Set("speaker", speaker)
	return c.Render(200, r.HTML("speakers/edit.html"))
}

// Update changes a Speaker in the DB. This function is mapped to
// the path PUT /speakers/{speaker_id}
func (v SpeakersResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Speaker
	speaker := &models.Speaker{}

	if err := tx.Find(speaker, c.Param("speaker_id")); err != nil {
		return c.Error(404, err)
	}

	// Bind Speaker to the html form elements
	if err := c.Bind(speaker); err != nil {
		return errors.WithStack(err)
	}

	verrs, err := tx.ValidateAndUpdate(speaker)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make speaker available inside the html template
		c.Set("speaker", speaker)

		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.HTML("speakers/edit.html"))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Speaker was updated successfully")

	// and redirect to the speakers index page
	return c.Redirect(302, "/speakers/%s", speaker.ID)
}

// Destroy deletes a Speaker from the DB. This function is mapped
// to the path DELETE /speakers/{speaker_id}
func (v SpeakersResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Speaker
	speaker := &models.Speaker{}

	// To find the Speaker the parameter speaker_id is used.
	if err := tx.Find(speaker, c.Param("speaker_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(speaker); err != nil {
		return errors.WithStack(err)
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", "Speaker was destroyed successfully")

	// Redirect to the speakers index page
	return c.Redirect(302, "/speakers")
}
