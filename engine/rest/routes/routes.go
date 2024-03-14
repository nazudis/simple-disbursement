package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/nazudis/disbursement/engine/rest/middleware"
	"github.com/nazudis/disbursement/src"
)

func AppRoutes(r chi.Router) {
	r.Route("/payments", func(pr chi.Router) {

		// Public routes
		pr.Group(func(pubR chi.Router) {
			authController, _ := src.InitializeAuthController()
			transactionController, _ := src.InitializeTransactionController()

			pubR.Get("/auth/token", authController.GetAccessToken)
			pubR.Post("/disbursements/cb", transactionController.PostCallbackDisbursement)
		})

		// Private routes
		pr.Group(func(privR chi.Router) {
			transactionController, _ := src.InitializeTransactionController()
			privR.Use(middleware.AuthSessionMiddleware)
			privR.Get("/account-verification", transactionController.GetAccountVerification)
			privR.Post("/disbursements", transactionController.PostDisbursement)
		})

	})

}
