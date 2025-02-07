package utils

import (
	"golang.org/x/crypto/bcrypt"
	"os"
	"strconv"
)

func GenerateInviteLink(authorLink int64) (inviteLink string) {
	link, _ := bcrypt.GenerateFromPassword([]byte(strconv.Itoa(int(authorLink))), bcrypt.DefaultCost)
	inviteLink = os.Getenv("LINK_TEMPLATE") + string(link)

	return inviteLink
}
