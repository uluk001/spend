package postgres

import (
	"fmt"

	"github.com/uluk001/spend/internal/model"
)

type UserRepo struct {
    db *Database
}

func NewUserRepository(db *Database) *UserRepo {
    return &UserRepo{
        db: db,
    }
}

func (r *UserRepo) Create(user *model.User) error {
    return r.db.GetDB().Create(user).Error
}

func (r *UserRepo) GetByID(id int) (*model.User, error) {
    var user model.User
    err := r.db.GetDB().First(&user, id).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *UserRepo) Update(id uint, user *model.User) (*model.User, error) {
    var existingUser model.User

    if err := r.db.GetDB().First(&existingUser, id).Error; err != nil {
        return nil, fmt.Errorf("User with %d id doesn't found", id);
    };
    updates := map[string]interface{} {
        "telegram_id": user.TelegramID,
        "username":    user.Username,
    }

    if err := r.db.GetDB().Model(&existingUser).Updates(updates).Error; err != nil {
        return nil, fmt.Errorf("failed to update user: %w", err)
    }

    var updatedUser model.User
    if err := r.db.GetDB().First(&updatedUser, id).Error; err != nil {
        return nil, fmt.Errorf("failed to fetch updated user: %w", err)
    }

    return &updatedUser, nil
}
