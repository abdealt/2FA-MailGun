package store

import (
	"auth/database"
	"auth/model"
	"log"
)

// GetUserByEmail : Récupère un utilisateur depuis la base de données en fonction de son email
func GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	// Recherche un utilisateur avec l'email fourni
	result := database.DB.Where("email = ?", email).First(&user)
	return &user, result.Error
}

// CreateUser : Ajoute un nouvel utilisateur dans la base de données
func CreateUser(user *model.User) error {
	// Insère un utilisateur dans la base de données
	return database.DB.Create(user).Error
}

// CreateUsers : Ajoute un utilisateur avec un email et un mot de passe (non sécurisé)
func CreateUsers(email, password string) (*model.User, error) {
	user := &model.User{
		Email:    email,
		Password: password, // Stocke le mot de passe en clair (non recommandé pour la production)
	}

	// Sauvegarde l'utilisateur dans la base de données
	result := database.DB.Create(user)
	if result.Error != nil {
		log.Println("Erreur lors de la création de l'utilisateur :", result.Error)
		return nil, result.Error
	}

	return user, nil
}

// UpdateUserPassword : Met à jour le mot de passe d'un utilisateur (non sécurisé)
func UpdateUserPassword(userID uint, newPassword string) error {
	var user model.User

	// Récupère l'utilisateur depuis la base de données
	if err := database.DB.First(&user, userID).Error; err != nil {
		return err
	}

	// Met à jour le mot de passe (en clair)
	user.Password = newPassword
	return database.DB.Save(&user).Error
}
