package main

import (
	"strings"
	"testing"
)

func TestHasAccess(t *testing.T) {
	adminRole := Role{Name: "Admin", Permissions: []string{"create", "read", "update", "delete"}}
	userRole := Role{Name: "User", Permissions: []string{"read"}}

	admin := User{ID: 1, Name: "AdminUser", Email: "admin@example.com", Password: "adminpass", Role: adminRole}
	user := User{ID: 2, Name: "RegularUser", Email: "user@example.com", Password: "userpass", Role: userRole}

	tests := []struct {
		user       User
		permission string
		want       bool
	}{
		{admin, "create", true},
		{admin, "delete", true},
		{user, "read", true},
		{user, "update", false},
		{user, "delete", false},
	}

	for _, tt := range tests {
		got := tt.user.HasAccess(tt.permission)
		if got != tt.want {
			t.Errorf("HasAccess(%s, %s) = %v; want %v", tt.user.Name, tt.permission, got, tt.want)
		}
	}
}

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Add(2,3) = %d; want %d", result, expected)
	}
}

func TestChangeName(t *testing.T) {
	user := User{ID: 1, Name: "OldName", Email: "test@example.com", Password: "password", Role: Role{Name: "User"}}
	newName := "NewName"

	err := changeName(&user, newName)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if user.Name != newName {
		t.Errorf("Expected user.Name to be updated to %s, but got %s", newName, user.Name)
	}

	err = changeName(&user, "")
	if err == nil {
		t.Errorf("Expected error for empty name, but got nil")
	}
}

func TestSetupUsers(t *testing.T) {
	admin, user := setupUsers()

	if admin.Name != "Maryna" || admin.Role.Name != "Admin" {
		t.Errorf("setupUsers() admin user incorrect: %+v", admin)
	}

	if user.Name != "Ivan" || user.Role.Name != "User" {
		t.Errorf("setupUsers() regular user incorrect: %+v", user)
	}
}

// **????? ????** ??? `runApp()`
func TestRunApp(t *testing.T) {
	output := runApp()

	if !strings.Contains(output, "Can user Maryna create? true") {
		t.Errorf("Expected admin to have 'create' permission")
	}
	if !strings.Contains(output, "Can user Ivan delete? false") {
		t.Errorf("Expected user to NOT have 'delete' permission")
	}
	if !strings.Contains(output, "New user name: Ivan Updated") {
		t.Errorf("Expected user name to be changed to 'Ivan Updated'")
	}
}