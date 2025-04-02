package model

import "gorm.io/gorm"

// Modèle User : Représente un utilisateur dans la base de données
type User struct {
	gorm.Model        // Fournit des champs intégrés comme ID, CreatedAt, UpdatedAt, DeletedAt
	Email      string `gorm:"uniqueIndex"` // L'email de l'utilisateur, doit être unique dans la base de données
	Password   string // Le mot de passe de l'utilisateur (doit être stocké de manière sécurisée, par exemple, haché)
}
