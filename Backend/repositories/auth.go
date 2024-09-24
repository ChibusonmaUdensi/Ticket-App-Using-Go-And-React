package repositories

import (
    "context"
	"gorm.io/gorm"

    "github.com/ChibusonmaUdensi/Ticket-App-Using-Go-And-React/models"
)


type AuthRepository struct {
    db *gorm.DB
}
func (r *AuthRepository) RegisterUser(ctx context.Context, registerData *models.AuthCredentials)(*models.User, error){
	user := &models.User{
        Email: registerData.Email,
        Password: registerData.Password,
    }

    res := r.db.Create(user)

    if res.Error!= nil {
        return nil, res.Error
    }

    return user, nil
}

func (r *AuthRepository) GetUser(ctx context.Context, query interface{}, args ...interface{})(*models.User, error){
//Using query interface{} and args ...interface{} allows for flexible 
//and dynamic querying in functions, making it 
//easier to build reusable and adaptable database interaction methods.
user := &models.User{}

if res := r.db.Model(user).Where(query,args).First(user); res.Error != nil{
	return nil, res.Error
}
return user, nil
}

func NewAuthRepository(db *gorm.DB) models.AuthRepository {
	return &AuthRepository{
		db: db,
	}
}