package handler

import (
	"github.com/gin-gonic/gin"
	"main/internal/api/service"
	e "main/internal/core/error"
	"main/internal/core/model"
	"main/internal/core/model/request"
	"net/http"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{
		service: s,
	}
}

// Register godoc
// @Summary      Register a user
// @Description  이메일, 비밀번호로 유저 회원가입
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        user  body  request.RegisterRequest  true  "A user information"
// @Success      201
// @Failure      400  {object}  httpError  "Bad request error"
// @Failure      401  {object}  httpError  "Unauthorized error"
// @Failure      404  {object}  httpError  "Not found error"
// @Failure      500  {object}  httpError  "Internal server error"
// @Router       /register [post]
func (h *UserHandler) Register(ctx *gin.Context) {
	var req request.RegisterRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		sendJsonParsingErr(ctx, err)
		return
	}

	if err := h.service.Register(&req); err != nil {
		sendErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, nil)
}

// GetAllUsers godoc
// @Summary      Return all users
// @Description  이메일, 비밀번호로 유저 회원가입
// @Tags         user
// @Accept       json
// @Produce      json
// @Param                    Authorization            header  string  true  "Bearer {access_token}"
// @Param        body  body  request.RegisterRequest  true    "A user information"
// @Success      201
// @Failure      400  {object}  httpError  "Bad request error"
// @Failure      401  {object}  httpError  "Unauthorized error"
// @Failure      404  {object}  httpError  "Not found error"
// @Failure      500  {object}  httpError  "Internal server error"
// @Router       /users [get]
func (h *UserHandler) GetAllUsers(ctx *gin.Context) {
	option := model.NewQuery()

	if err := ctx.BindQuery(option); err != nil {
		sendErr(ctx, err)
		return
	}

	users, err := h.service.GetUsers(option)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, users)
}

// GetUser returns a user with id
func (h *UserHandler) GetUser(ctx *gin.Context) {
	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	userID, err := extractUserId(ctx)
	if err != nil {
		userID = user.ID
	}

	fields := getFieldsFromContext(ctx)

	resp, err := h.service.GetUser(userID, fields)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// DeleteUser deletes a user
func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	err = h.service.DeleteUser(user.ID)
	if err != nil {

		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// EditUserProfile edit a user profile
func (h *UserHandler) EditUserProfile(ctx *gin.Context) {
	var req model.Profile

	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		sendJsonParsingErr(ctx, err)
		return
	}

	err = h.service.UpdateProfile(user.ID, &req)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// UploadUserProfileImage edit user's profile image
func (h *UserHandler) UploadUserProfileImage(ctx *gin.Context) {
	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	file, header, err := ctx.Request.FormFile("profile_image")
	if err != nil {
		sendErr(ctx, e.ErrInvalidFile)
		return
	}
	defer file.Close()

	err = h.service.UploadProfileImage(user.ID, file, header)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// GetFollowings returns list of followings
func (h *UserHandler) GetFollowings(ctx *gin.Context) {
	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	followings, err := h.service.GetFollowings(user.ID)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, followings)
}

// GetFollowers returns list of followers
func (h *UserHandler) GetFollowers(ctx *gin.Context) {
	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	followings, err := h.service.GetFollowers(user.ID)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, followings)
}

// FollowUser refreshes access token
func (h *UserHandler) FollowUser(ctx *gin.Context) {
	var data struct {
		FollowerID uint `json:"follower_id"`
	}

	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	err = ctx.ShouldBindJSON(&data)
	if err != nil {
		sendJsonParsingErr(ctx, err)
		return
	}

	err = h.service.Follow(user.ID, data.FollowerID)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// UnfollowUser refreshes access token
func (h *UserHandler) UnfollowUser(ctx *gin.Context) {
	var data struct {
		FollowingID uint `json:"following_id"`
	}

	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	err = ctx.ShouldBindJSON(&data)
	if err != nil {
		sendJsonParsingErr(ctx, err)
		return
	}

	err = h.service.UnFollow(user.ID, data.FollowingID)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// GetFavoriteQuants returns a favorite quant list of user
func (h *UserHandler) GetFavoriteQuants(ctx *gin.Context) {
	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	quants, err := h.service.GetFavoriteQuants(user.ID)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"count":  len(quants),
		"quants": quants,
	})
}

// AddToFavoriteQuants add a quant to favorite list
func (h *UserHandler) AddToFavoriteQuants(ctx *gin.Context) {
	var data struct {
		QuantID uint `json:"quant_id"`
	}

	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	err = ctx.ShouldBindJSON(&data)
	if err != nil {
		sendJsonParsingErr(ctx, err)
		return
	}

	err = h.service.AddToFavoriteQuants(user.ID, data.QuantID)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, nil)
}

// DeleteFromFavoriteQuants add a quant to favorite list
func (h *UserHandler) DeleteFromFavoriteQuants(ctx *gin.Context) {
	var data struct {
		QuantID uint `json:"quant_id"`
	}

	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	err = ctx.ShouldBindJSON(&data)
	if err != nil {
		sendJsonParsingErr(ctx, err)
		return
	}

	err = h.service.DeleteFromFavoriteQuants(user.ID, data.QuantID)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
