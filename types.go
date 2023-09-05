package main

type ContactUpdateParams struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Email       string
}

type Contact struct {
	ID          string
	FirstName   string
	LastName    string
	PhoneNumber string
	Email       string
}

func (c *Contact) Update(params ContactUpdateParams) {
	c.FirstName = params.FirstName
	c.LastName = params.LastName
	c.PhoneNumber = params.PhoneNumber
	c.Email = params.Email
}

type ContactForm struct {
	ID          string
	FirstName   string
	LastName    string
	PhoneNumber string
	Email       string

	Errors map[string]string
}

func NewContactForm() *ContactForm {
	return &ContactForm{
		Errors: make(map[string]string),
	}
}

func NewContactFormFromContact(c Contact) *ContactForm {
	return &ContactForm{
		ID:          c.ID,
		FirstName:   c.FirstName,
		LastName:    c.LastName,
		PhoneNumber: c.PhoneNumber,
		Email:       c.Email,
		Errors:      make(map[string]string),
	}
}

func (f *ContactForm) IsValid() bool {
	if f.FirstName == "" {
		f.Errors["firstName"] = "required"
	}
	if f.LastName == "" {
		f.Errors["lastName"] = "required"
	}
	if f.PhoneNumber == "" {
		f.Errors["phoneNumber"] = "required"
	}
	if f.Email == "" {
		f.Errors["email"] = "required"
	}
	return len(f.Errors) == 0
}

func (f *ContactForm) ToContact() Contact {
	return Contact{
		ID:          f.ID,
		FirstName:   f.FirstName,
		LastName:    f.LastName,
		PhoneNumber: f.PhoneNumber,
		Email:       f.Email,
	}
}
