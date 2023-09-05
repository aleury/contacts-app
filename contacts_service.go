package main

import (
	"errors"
	"strings"
	"sync"

	"github.com/segmentio/ksuid"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

type ContactsStore struct {
	mux      sync.RWMutex
	contacts map[string]Contact
}

func NewContactsStore() *ContactsStore {
	contacts := []Contact{
		{newID(), "John", "Smith", "123-456-7890", "john@example.com"},
		{newID(), "Dana", "Crandith", "123-456-7890", "dcran@example.com"},
		{newID(), "Edith", "Neutvaar", "123-456-7890", "en@example.com"},
	}
	return &ContactsStore{
		contacts: toMap(contacts),
	}
}

func (s *ContactsStore) Save(contact Contact) error {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.contacts[contact.ID] = contact
	return nil
}

func (s *ContactsStore) Search(search string) []Contact {
	s.mux.RLock()
	defer s.mux.RUnlock()
	results := []Contact{}
	for _, c := range s.contacts {
		if matchesContact(c, search) {
			results = append(results, c)
		}
	}
	return results
}

func (s *ContactsStore) All() []Contact {
	s.mux.RLock()
	defer s.mux.RUnlock()
	contacts := maps.Values(s.contacts)
	slices.SortFunc(contacts, func(a, b Contact) int {
		aID, _ := ksuid.Parse(a.ID)
		bID, _ := ksuid.Parse(b.ID)
		return ksuid.Compare(aID, bID)
	})
	return contacts
}

func (s *ContactsStore) Find(id string) (Contact, error) {
	s.mux.RLock()
	defer s.mux.RUnlock()
	if c, ok := s.contacts[id]; ok {
		return c, nil
	}
	return Contact{}, errors.New("contact not found")
}

func (s *ContactsStore) Delete(contact Contact) error {
	s.mux.Lock()
	defer s.mux.Unlock()
	delete(s.contacts, contact.ID)
	return nil
}

func matchesContact(contact Contact, search string) bool {
	items := []string{
		strings.ToLower(contact.FirstName),
		strings.ToLower(contact.LastName),
		strings.ToLower(strings.Trim(contact.PhoneNumber, "-")),
		strings.ToLower(contact.Email),
	}
	itemsStr := strings.Join(items, "")
	return strings.Contains(itemsStr, strings.ToLower(search))
}

func newID() string {
	return ksuid.New().String()
}

func toMap(contacts []Contact) map[string]Contact {
	m := map[string]Contact{}
	for _, c := range contacts {
		m[c.ID] = c
	}
	return m
}
