package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

type ContactsHandler struct {
	contacts *ContactsStore
}

func (h *ContactsHandler) Redirect(c *fiber.Ctx) error {
	return c.Redirect("/contacts")
}

func (h *ContactsHandler) Index(c *fiber.Ctx) error {
	search := c.Query("q")
	var contacts []Contact
	if search != "" {
		contacts = h.contacts.Search(search)
	} else {
		contacts = h.contacts.All()
	}
	return c.Render("index", fiber.Map{
		"q":        search,
		"contacts": contacts,
	})
}

func (h *ContactsHandler) New(c *fiber.Ctx) error {
	form := ContactForm{
		Errors: make(map[string]string),
	}
	return c.Render("new", fiber.Map{"form": form})
}

func (h *ContactsHandler) Create(c *fiber.Ctx) error {
	form := ContactForm{
		ID:          newID(),
		FirstName:   utils.CopyString(c.FormValue("firstName")),
		LastName:    utils.CopyString(c.FormValue("lastName")),
		PhoneNumber: utils.CopyString(c.FormValue("phoneNumber")),
		Email:       utils.CopyString(c.FormValue("email")),
		Errors:      map[string]string{},
	}
	if !form.IsValid() {
		return c.Render("new", fiber.Map{"form": form})
	}
	contact := form.ToContact()
	h.contacts.Save(contact)
	return c.Redirect("/contacts")
}

func (h *ContactsHandler) Show(c *fiber.Ctx) error {
	id := c.Params("id")
	contact, err := h.contacts.Find(id)
	if err != nil {
		return fiber.ErrNotFound
	}
	return c.Render("show", fiber.Map{"contact": contact})
}

func (h *ContactsHandler) Edit(c *fiber.Ctx) error {
	id := c.Params("id")
	contact, err := h.contacts.Find(id)
	if err != nil {
		return fiber.ErrNotFound
	}
	form := NewContactFormFromContact(contact)
	return c.Render("edit", fiber.Map{"form": form})
}

func (h *ContactsHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	form := ContactForm{
		ID:          id,
		FirstName:   utils.CopyString(c.FormValue("firstName")),
		LastName:    utils.CopyString(c.FormValue("lastName")),
		PhoneNumber: utils.CopyString(c.FormValue("phoneNumber")),
		Email:       utils.CopyString(c.FormValue("email")),
		Errors:      map[string]string{},
	}
	if !form.IsValid() {
		return c.Render("edit", fiber.Map{"form": form})
	}

	contact, err := h.contacts.Find(id)
	if err != nil {
		return fiber.ErrNotFound
	}
	contact.Update(ContactUpdateParams{
		FirstName:   form.FirstName,
		LastName:    form.LastName,
		PhoneNumber: form.PhoneNumber,
		Email:       form.Email,
	})
	h.contacts.Save(contact)
	return c.Redirect("/contacts/" + id)
}

func (h *ContactsHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	contact, err := h.contacts.Find(id)
	if err != nil {
		return fiber.ErrNotFound
	}
	h.contacts.Delete(contact)
	return c.Redirect("/contacts")
}
