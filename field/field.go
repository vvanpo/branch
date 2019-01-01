package field

import (
	"sync"
)

// A Field is used to categorize and format information associated with
// contacts.
type Field struct {
	lock        *sync.RWMutex
	repo        Fields
	name        label
	description text
	definition  Type
}

// NewField
func NewField(name, description string, definition Type, repository Fields) (*Field, error) {
	field := &Field{
		lock: new(sync.RWMutex),
		repo: repository,
	}

	if err := field.name.Set(name); err != nil {
		return nil, err
	}

	if err := field.description.Set(description); err != nil {
		return nil, err
	}

	if err := definition.Validate(); err != nil {
		return nil, err
	}

	field.definition = definition
	return field, nil
}

// Name
func (f Field) Name() string {
	f.lock.RLock()
	defer f.lock.RUnlock()
	return string(f.name)
}

// Description
func (f Field) Description() string {
	f.lock.RLock()
	defer f.lock.RUnlock()
	return string(f.description)
}

// SetDescription
func (f *Field) SetDescription(description string) error {
	f.lock.Lock()
	defer f.lock.Unlock()

	if err := f.description.Set(description); err != nil {
		return err
	}

	f.repo.Persist(f)
	return nil
}
