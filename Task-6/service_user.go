package Task_6

type UserService interface {
	GetUser(id string) string
}

func GetUserById(userService UserService, userId string) string {
	return userService.GetUser(userId)
}
