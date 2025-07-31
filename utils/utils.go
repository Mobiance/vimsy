package utils

import (
	"fmt"
	"os"
	"golang.org/x/sys/windows"
)

func IsAdmin() bool {
	var sid *windows.SID
	sid, err := windows.CreateWellKnownSid(windows.WinBuiltinAdministratorsSid)
	if err != nil {
		fmt.Println("Error creating well-known SID:", err)
		return false
	}
	token := windows.Token(0)
	isMember, err := token.IsMember(sid)
	if err != nil {
		fmt.Println("Error checking token membership:", err)
		return false
	}
	return isMember
}

func RemoveIfExists(path string) error {
	if _, err := os.Stat(path); err == nil {
		if err := os.RemoveAll(path); err != nil {
			return fmt.Errorf("failed to remove existing file or symlink: %w", err)
		}
	}
	return nil
}
