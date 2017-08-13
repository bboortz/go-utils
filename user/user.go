package user

import (
	"github.com/bboortz/go-utils/stringutil"
	osuser "os/user"
	"strconv"
)

// User is the structure of a user
type User struct {
	UID       int
	Gid       int
	Username  string
	Groupname string
	Groupids  []int
	HomeDir   string
}

// GetCurrentUser retrieves the current
func GetCurrentUser() (User, error) {
	var user User
	var err error

	theUser, err := osuser.Current()
	// if str, ok := val.(string); ok {
	// if theUser, err := osuser.Current(); err != nil {
	if err != nil {
		return user, err
	}
	theGroup, err := osuser.LookupGroupId(theUser.Gid)
	if err != nil {
		return user, err
	}
	uid, err := strconv.Atoi(theUser.Uid)
	if err != nil {
		return user, err
	}
	gid, err := strconv.Atoi(theUser.Gid)
	if err != nil {
		return user, err
	}
	username := theUser.Username
	groupname := theGroup.Name
	groupidsTmp, err := theUser.GroupIds()
	if err != nil {
		return user, err
	}
	groupids, err := stringutil.ConvertStringArrayToIntArray(groupidsTmp)
	if err != nil {
		return user, err
	}
	homeDir := theUser.HomeDir
	user = User{
		UID:       uid,
		Gid:       gid,
		Username:  username,
		Groupname: groupname,
		Groupids:  groupids,
		HomeDir:   homeDir,
	}
	return user, err
}
