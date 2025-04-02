package route

import (
	"auth/handlers"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes configure les routes de l'application
func SetupRoutes(app *fiber.App) {
	// Route d'affichage du formulaire de login
	app.Get("/login", handlers.ShowLogin)

	// Soumission du formulaire de login
	app.Post("/login", handlers.Login)

	// Vérification de l'OTP (One-Time Password)
	app.Post("/verifyotp", handlers.VerifyOTP)

	// Dashboard (page protégée, nécessite une vérification de session/authentification)
	app.Get("/dashboard", handlers.Dashboard)

	// Routes pour "Mot de passe oublié"
	// Affiche le formulaire pour demander un lien de réinitialisation de mot de passe
	app.Get("/forgot-password", handlers.ShowForgotPassword)

	// Soumission du formulaire "Mot de passe oublié"
	app.Post("/forgot-password", handlers.ForgotPassword)

	// Affiche le formulaire pour réinitialiser le mot de passe
	app.Get("/reset-password", handlers.ShowResetPassword)

	// Soumission du formulaire pour réinitialiser le mot de passe
	app.Post("/reset-password", handlers.ResetPassword)
}
