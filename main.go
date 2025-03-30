package main

import (
	"fmt"
)

// ??????????? ?????? ??? ????????????
const defaultPassword = "changeme"

// ????????? ???? ???????????
type Role struct {
	Name        string
	Permissions []string
}

// ????????? ???????????
type User struct {
	ID       int
	Name     string
	Email    string
	Password string
	Role     Role
}

// ??????? ????????? ??????? ??????????? ?? ??????? ???????
func (u User) HasAccess(permission string) bool {
	for _, p := range u.Role.Permissions {
		if p == permission {
			return true
		}
	}
	return false
}

// ??????? ????????? ???? ?????
func Add(a, b int) int {
	return a + b
}

// ??????? ????? ????? ??????????? (????????? ??????????? ?????? ?????)
func changeName(u *User, newName string) error {
	if newName == "" {
		return fmt.Errorf("name cannot be empty")
	}
	u.Name = newName
	fmt.Println("Name changed successfully")
	return nil
}

// ??????? ????????? ???????????? ?? ??????????? ??????
func setupUsers() (User, User) {
	adminRole := Role{Name: "Admin", Permissions: []string{"create", "read", "update", "delete"}}
	userRole := Role{Name: "User", Permissions: []string{"read"}}

	admin := User{ID: 1, Name: "Maryna", Email: "maryna@example.com", Password: defaultPassword, Role: adminRole}
	user := User{ID: 2, Name: "Ivan", Email: "ivan@example.com", Password: defaultPassword, Role: userRole}

	return admin, user
}

// **???? ???????**: ??????? ??????? ?????? ???????? (??????? main)
func runApp() string {
	admin, user := setupUsers()

	result := fmt.Sprintf(
		"Can user %s create? %v\nCan user %s delete? %v\nCan user %s read? %v\n",
		admin.Name, admin.HasAccess("create"),
		user.Name, user.HasAccess("delete"),
		user.Name, user.HasAccess("read"),
	)

	err := changeName(&user, "Ivan Updated")
	if err != nil {
		result += fmt.Sprintf("Error: %v\n", err)
	} else {
		result += fmt.Sprintf("New user name: %s\n", user.Name)
	}

	return result
}

// ??????? ???????
func main() {
	defer fmt.Println("Program finished")
	fmt.Print(runApp())
}