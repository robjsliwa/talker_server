package main

// MockStore - fake store
type MockStore struct {
	RoomStore []Room
	UserStore []User
}

var mainStore *MockStore

// NewMockStore - create new mock store
func NewMockStore() (*MockStore, error) {
	mockStore := &MockStore{
		RoomStore: make([]Room, 0),
		UserStore: make([]User, 0),
	}

	return mockStore, nil
}

// FindUser - find if the given user already exists
func (ms *MockStore) FindUser(user User) (*User, bool) {
	for _, storedUser := range ms.UserStore {
		if storedUser.Name == user.Name {
			return &storedUser, true
		}
	}

	return nil, false
}

// AddUser - add new user to the store
func (ms *MockStore) AddUser(user User) {
	ms.UserStore = append(ms.UserStore, user)
}
