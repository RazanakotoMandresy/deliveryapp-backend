package handler

import (
	"fmt"

	services "github.com/RazanakotoMandresy/hotels-backend/internal/service"
	"golang.org/x/crypto/bcrypt"

	"net/http"
)


// first paramaeter the hashed passord from the db second parameter from the body
func IsTruePassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return fmt.Errorf("you given a false passewords %v", err)
	}
	return nil
}

func (s service) Register() http.HandlerFunc {
	type registerReq struct {
		Name      string `json:"name"`
		Passwords string `json:"passwords"`
		Mail      string `json:"mail"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(registerReq)
		if err := s.decode(r, req); err != nil {
			s.respond(w, errorResponse{err.Error() + " Decode's problems"}, http.StatusBadRequest)
			return
		}
		res, err := s.services.Register(r.Context(), services.RegisterParams{
			Name:     req.Name,
			Password: req.Passwords,
			Mail:     req.Mail,
		})
		if err != nil {
			s.respond(w, errorResponse{err.Error() + " Services register error"}, http.StatusInternalServerError)
			return
		}
		s.respond(w, responseUsers{*res}, http.StatusOK)
	}
}
