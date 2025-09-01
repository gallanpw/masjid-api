package routes

import (
	"masjid-api/controllers"
	"masjid-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		// Rute untuk user login, tidak memerlukan otentikasi
		// api.POST("/users/login", controllers.Login)

		// Rute Jadwal Sholat (tanpa autentikasi)
		api.GET("/jadwal-sholat", controllers.GetJadwalSholat)

		// Rute untuk otentikasi
		users := api.Group("/users")
		{
			users.POST("/login", controllers.LoginUser)
			users.POST("/logout", controllers.LogoutUser)
		}

		// Rute yang dilindungi
		// protected := api.Group("/")
		// protected.Use(middleware.JWTAuthMiddleware())
		// {
		// Contoh rute yang memerlukan otentikasi
		// protected.GET("/prayer_schedules/all", controllers.GetPrayerSchedules)
		// Tambahkan rute protected lainnya di sini
		// }

		// Rute Ustadz
		ustadzRoutes := api.Group("/ustadz")
		{
			// GET /api/ustadz: Endpoint publik untuk melihat data (tanpa autentikasi)
			ustadzRoutes.GET("/", controllers.GetAllUstadz)
			ustadzRoutes.GET("/:id", controllers.GetUstadzByID)

			// Rute yang dilindungi oleh JWT middleware
			protected := ustadzRoutes.Group("/")
			protected.Use(middleware.JWTAuthMiddleware())
			{
				protected.POST("/", controllers.CreateUstadz)
				protected.PUT("/:id", controllers.UpdateUstadz)
				protected.DELETE("/:id", controllers.DeleteUstadz)
			}
		}

		// Rute Kategori Kajian
		kategoriKajianRoutes := api.Group("/kategori-kajian")
		{
			// GET /api/kategori-kajian: Endpoint publik untuk melihat data (tanpa autentikasi)
			kategoriKajianRoutes.GET("/", controllers.GetAllKategoriKajian)
			kategoriKajianRoutes.GET("/:id", controllers.GetKategoriKajianByID)

			// Rute yang dilindungi oleh JWT middleware
			protected := kategoriKajianRoutes.Group("/")
			protected.Use(middleware.JWTAuthMiddleware())
			{
				protected.POST("/", controllers.CreateKategoriKajian)
				protected.PUT("/:id", controllers.UpdateKategoriKajian)
				protected.DELETE("/:id", controllers.DeleteKategoriKajian)
			}
		}

		// Rute Kajian
		kajianRoutes := api.Group("/kajian")
		{
			// GET /api/kajian: Endpoint publik untuk melihat data (tanpa autentikasi)
			kajianRoutes.GET("/", controllers.GetAllKajian)
			kajianRoutes.GET("/:id", controllers.GetKajianByID)

			// Rute yang dilindungi oleh JWT middleware
			protected := kajianRoutes.Group("/")
			protected.Use(middleware.JWTAuthMiddleware())
			{
				protected.POST("/", controllers.CreateKajian)
				protected.PUT("/:id", controllers.UpdateKajian)
				protected.DELETE("/:id", controllers.DeleteKajian)
			}
		}

		// Rute Finance
		financeRoutes := api.Group("/finance")
		{
			// GET /api/finance: Endpoint publik untuk melihat data (tanpa autentikasi)
			financeRoutes.GET("/", controllers.GetAllFinance)

			// Rute yang dilindungi oleh JWT middleware
			protected := financeRoutes.Group("/")
			protected.Use(middleware.JWTAuthMiddleware())
			{
				protected.POST("/donations", controllers.CreateDonation)
				protected.POST("/expenses", controllers.CreateExpense)

			}
		}

		// Rute Role
		roleRoutes := api.Group("/roles")
		{
			// Rute yang dilindungi oleh JWT middleware
			protected := roleRoutes.Group("/")
			protected.Use(middleware.JWTAuthMiddleware())
			{
				protected.GET("/", controllers.GetAllRoles)
				protected.GET("/:id", controllers.GetRoleByID)
				protected.POST("/", controllers.CreateRole)
				protected.PUT("/:id", controllers.UpdateRole)
				protected.DELETE("/:id", controllers.DeleteRole)
			}
		}

		// Rute User
		userRoutes := api.Group("/users")
		{
			// Rute yang dilindungi oleh JWT middleware
			protected := userRoutes.Group("/")
			protected.Use(middleware.JWTAuthMiddleware())
			{
				protected.GET("/", controllers.GetAllUsers)
				protected.GET("/:id", controllers.GetUserByID)
				protected.POST("/", controllers.CreateUser)
				protected.PUT("/:id", controllers.UpdateUser)
				protected.DELETE("/:id", controllers.DeleteUser)
			}
		}
	}
}
