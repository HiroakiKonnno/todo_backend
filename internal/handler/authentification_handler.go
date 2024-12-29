package handler

import (
	"errors"
	"net/http"

	crypto "todo_backend/internal/libraries"
	"todo_backend/internal/model"
	"todo_backend/internal/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type JsonRequest struct {
	LoginId  string `json:"login_id"`
	Password  string `json:"password"`
	Name string `json:"name"`
}


func RegisterAuthentificationRoutes(r *gin.Engine, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	r.POST("/api/signup", CreateUser(userRepo))
	r.POST("/api/signin", SignInUser(userRepo))
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

