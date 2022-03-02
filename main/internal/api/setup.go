package api

import (
	"main/internal/api/handler"
	"main/internal/api/repo"
	"main/internal/api/route"
	"main/internal/api/service"
	"main/internal/conf/aws"
	"main/internal/core/model"
	"main/internal/core/model/quant"
)

func (r *Router) setAll() {
	r.migrate()
	r.setupAuth()
	r.setupUser()
	r.setupQuant()
	r.setupComment()
	r.setupReply()
}

func (r *Router) migrate() {
	ms := []interface{}{
		&model.User{},
		&model.Profile{},
		&model.Quant{},
		&model.QuantOption{},
		&quant.MainSector{},
		&model.Reply{},
		&model.Comment{},
		&model.Reply{},
	}
	r.mysqlDB.Migrate(ms)
}

func (r *Router) setupAuth() {
	rp := repo.NewAuthRepo(r.mysqlDB.DB)
	sv := service.NewAuthService(rp)
	hdr := handler.NewAuthHandler(sv)
	route.SetAuth(r.getGroup(), hdr)
}

func (r *Router) setupUser() {
	a := aws.New()
	rp := repo.NewUser(r.mysqlDB.DB, a)
	sv := service.NewUserService(rp, a)
	hdr := handler.NewUserHandler(sv)
	route.SetUser(r.getGroup(), hdr, authMid)
}

func (r *Router) setupQuant() {
	rp := repo.NewQuantRepo(r.mysqlDB.DB, r.mongoDB.DB)
	sv := service.NewQuantService(rp)
	hdr := handler.NewQuantHandler(sv)
	route.SetQuant(r.getGroupWithAuth(), hdr)
}

func (r *Router) setupComment() {
	rp := repo.NewCommentRepo(r.mysqlDB.DB)
	sv := service.NewCommentService(rp)
	hdr := handler.NewCommentHandler(sv)
	route.SetComment(r.getGroupWithAuth(), hdr)
}

func (r *Router) setupReply() {
	rp := repo.NewReplyRepository(r.mysqlDB.DB)
	sv := service.NewReplyService(rp)
	hdr := handler.NewReplyHandler(sv)
	route.SetReply(r.getGroupWithAuth(), hdr)
}
