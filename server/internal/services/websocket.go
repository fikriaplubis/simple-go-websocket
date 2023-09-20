package services

import (
	"net/http"

	// "github.com/CloudyKit/jet/v6"
	"github.com/gorilla/websocket"
)

type Service interface {
	UpgradeConnection(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error)
	//RenderPage(w http.ResponseWriter, tmpl string, data jet.VarMap) error
}

type service struct {
	u *websocket.Upgrader
	// v *jet.Set
}

// func NewService(u *websocket.Upgrader, v *jet.Set) *service {
// 	return &service{u, v}
// }

func NewService(u *websocket.Upgrader) *service {
	return &service{u}
}

func (s *service) UpgradeConnection(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	ws, err := s.u.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}

	return ws, nil
}

// func (s *service) RenderPage(w http.ResponseWriter, tmpl string, data jet.VarMap) error {
// 	view, err := s.v.GetTemplate(tmpl)
// 	if err != nil {
// 		return err
// 	}

// 	err = view.Execute(w, data, nil)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
