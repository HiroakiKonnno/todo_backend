package handler

import (
	"errors"
	"net/http"

	crypto "todo_backend/internal/libraries"
	"todo_backend/internal/model"
	"todo_backend/internal/repository"

	auth "todo_backend/internal/auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type JsonRequest struct {
	LoginId  string `json:"login_id"`
	Password  string `json:"password"`
	Name string `json:"name"`
}


func RegisterAuthentificationRoutes(r *gin.RouterGroup, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	r.POST("/api/signup", CreateUser(userRepo))
	r.POST("/api/signin", SignInUser(userRepo))
	r.POST("/api/signout", SignOutUser(userRepo))
	r.GET("/api/me", GetCurrentUser(userRepo))
}

// ユーザを作成する
func CreateUser(repo *repository.UserRepositoryImpl) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var json  JsonRequest
		if err := ctx.ShouldBindJSON(&json); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error" : err.Error(),
		})
		return
		}

		_, err := repo.FindByUserId(json.LoginId)

		if !errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusBadRequest, gin.H{
					"message" : "そのUserIdは既に登録されています。",
			})
			return
		}
		encryptPw, err := crypto.PasswordEncrypt(json.Password)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
					"message" : "パスワードの暗号化でエラーが発生しました。",
			})
			return
		}

		user := model.User{LoginId: json.LoginId, Password: encryptPw, Name: json.Name}
		create_err := repo.CreateUser(&user)

		if create_err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "ユーザー作成中にエラーが発生しました。",
			})
			return
	}
		token, err := auth.GenerateJWT(user.Id, user.LoginId)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
				return
			}
			
		http.SetCookie(ctx.Writer, &http.Cookie{
			Name:     "jwt",
			Value:    token,
			Path:     "/",
			MaxAge:   3600 * 24,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteNoneMode, // ← これが重要
		})

		ctx.JSON(http.StatusCreated, model.PublicUser{
			Id: user.Id,
			Name: user.Name,
			LoginId: user.LoginId,
		})
	}
}

func SignInUser(repo *repository.UserRepositoryImpl) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var json JsonRequest
		if err := ctx.ShouldBindJSON(&json); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
					"error" : err.Error(),
			})
			return
		}

		user, err := repo.FindByUserId(json.LoginId)

		if errors.Is(err, gorm.ErrRecordNotFound)  {
			ctx.JSON(http.StatusInternalServerError, gin.H{
					"message" : "ユーザーが存在しません。",
			})
			return
		}

		err_pw := crypto.CompareHashAndPassword(user.Password, json.Password)
		if err_pw != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
					"message" : "パスワードが一致しません。",
			})
			return
		}

		token, err := auth.GenerateJWT(user.Id, user.LoginId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		http.SetCookie(ctx.Writer, &http.Cookie{
			Name:     "jwt",
			Value:    token,
			Path:     "/",
			MaxAge:   3600 * 24,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteNoneMode,
		})

		ctx.JSON(http.StatusOK, gin.H{
			"message" : "ログイン成功",
			"user" : model.PublicUser{
				Id: user.Id,
				Name: user.Name,
				LoginId: user.LoginId,
			},
		})
	}

}

func SignOutUser(repo *repository.UserRepositoryImpl) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		http.SetCookie(ctx.Writer, &http.Cookie{
			Name:     "jwt",
			Value:    "",
			Path:     "/",
			MaxAge:   -1,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteNoneMode, // ← これが重要
		})
	
		// レスポンスを返す
		ctx.JSON(200, gin.H{
			"message": "Logged out successfully",
		})
	}
}

func GetCurrentUser(repo *repository.UserRepositoryImpl) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, ok := auth.ExtractAndVerifyToken(ctx)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		loginId := claims["loginId"].(string)

		user, err := repo.FindByUserId(loginId)

		if errors.Is(err, gorm.ErrRecordNotFound)  {
			ctx.JSON(http.StatusInternalServerError, gin.H{
					"message" : "ユーザーが存在しません。",
			})
			return
		}
		ctx.JSON(http.StatusCreated, model.PublicUser{
			Id: user.Id,
			Name: user.Name,
			LoginId: user.LoginId,
		})
	}
}

