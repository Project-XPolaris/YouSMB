package smb

import (
	"bufio"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

var (
	DefaultUserManager UserManager = UserManager{
		Users: []*User{},
	}
	UserNotFoundError = errors.New("target user not found")
)

type UserManager struct {
	Users []*User
}

func (m *UserManager) Create(username string, password string) error {
	script := fmt.Sprintf("(echo \"%s\"; echo \"%s\") | smbpasswd -s -a %s\n", password, password, username)
	cmd := exec.Command("/bin/sh", "-c", script)
	output, err := cmd.Output()
	if err != nil {
		return err
	}
	fmt.Println(string(output))
	err = m.LoadUser()
	if err != nil {
		return err
	}
	return nil
}

func (m *UserManager) LoadUser() error {
	cmd := exec.Command("pdbedit", "-L")
	output, err := cmd.Output()
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ":")
		m.Users = append(m.Users, &User{Username: parts[0]})
	}
	return nil
}

func (m *UserManager) GetUserByName(username string) *User {
	for _, user := range m.Users {
		if user.Username == username {
			return user
		}
	}
	return nil
}
func (m *UserManager) RemoveUser(username string) error {
	user := m.GetUserByName(username)
	if user == nil {
		return UserNotFoundError
	}
	err := user.Remove()
	if err != nil {
		return err
	}
	err = m.LoadUser()
	if err != nil {
		return err
	}
	return nil
}

type User struct {
	Username string
}

func (u *User) Remove() error {
	cmd := exec.Command("pdbedit", "-x", u.Username)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
