package src

import (
	"encoding/json"
	"fmt"
	"installment_back/controllers"
	"installment_back/errors"
	"installment_back/models"
	"installment_back/repositories"
	"installment_back/services"
	"installment_back/storage"
	"io"
	"log"
	"net/http"
)

type IServer interface {
	Init() error
	Run() error
}

type Server struct {
	config                models.Config
	userController        *controllers.UserController
	cardController        *controllers.CardController
	installmentController *controllers.InstallmentController
	db                    storage.DataBase
}

func (s *Server) Init() error {
	s.config.Url = "0.0.0.0:7777"
	s.config.DBUrl = "mongodb://localhost:27017"
	s.config.DBName = "Installment_Front"

	ur := repositories.NewUserRepository(&s.db)
	cr := repositories.NewCardRepository(&s.db)
	ir := repositories.NewInstallmentRepository(&s.db)
	ps := repositories.NewPaymentRepository(&s.db)
	itr := repositories.NewItemRepository(&s.db)

	cs := services.NewCardService(cr)
	is := services.NewInstallmentService(ir, cs, ps, itr)
	us := services.NewUserService(ur, is, cr)

	s.userController = controllers.NewUserController(us)
	s.cardController = controllers.NewCardController(cs)
	s.installmentController = controllers.NewInstallmentController(is)

	return nil
}

func (s *Server) Run() error {
	if err := s.db.Connect(s.config.DBUrl, s.config.DBName); err != nil {
		return err
	}
	defer s.db.Disconnect()
	fmt.Printf("Server started! %s/EndPoint", s.config.Url)
	http.HandleFunc("/EndPoint", s.handler)
	if err := http.ListenAndServe(s.config.Url, nil); err != nil {
		return err
	}
	return nil
}

func (s *Server) handler(w http.ResponseWriter, r *http.Request) {
	data, _ := io.ReadAll(r.Body)
	var request models.RPCRequest
	var response models.RPCResponse
	if err := json.Unmarshal(data, &request); err != nil {
		log.Printf("Failed to decode JSON request: %v", err)
	}

	switch request.Method {
	case "card.add":
		{
			response = s.cardController.Add(request.Params)
		}
	case "installment.pay":
		{
			response = s.installmentController.Pay(request.Params)
		}
	case "user.get":
		response = s.userController.Get(request.Params)
	default:
		response = errors.Method_not_found

	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode JSON response: %v", err)
	}
}
