package dal

import (
	"github.com/c0dect/basic-rest-service/models"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"log"
)

type UserManager interface {
	AddUser(user models.User) (newUser models.User, err error)
	CheckIfUsernameExists(username string) bool
}

const entity_user = "User"

type userDAL struct {
	context context.Context
}

func NewUserDAL(context context.Context) *userDAL {
	userDal := new(userDAL)
	userDal.context = context
	return userDal
}

func (userDal *userDAL) AddUser(user models.User) (newUser models.User, err error) {

	userKey := datastore.NewIncompleteKey(userDal.context, entity_user, nil)
	userKey, err = datastore.Put(userDal.context, userKey, &user)
	user.UserId = userKey.IntID()
	userKey, err = datastore.Put(userDal.context, userKey, &user)

	return user, err
}

func (userDal *userDAL) CheckIfUsernameExists(username string) bool {
	log.Println(username)

	query := datastore.NewQuery(entity_product).Filter("Username =", username)

	var users []models.User
	_, err := query.GetAll(userDal.context, &users)
	log.Println(users)
	log.Println(len(users))

	if len(users) > 0 || err != nil {
		return true
	}
	return false
}
