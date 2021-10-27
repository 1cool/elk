// Code generated by entc, DO NOT EDIT.

package http

import (
	"net/http"
	"strconv"

	"github.com/mailru/easyjson"
	"go.uber.org/zap"
)

// Read fetches the ent.Category identified by a given url-parameter from the
// database and returns it to the client.
func (h *CategoryHandler) List(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "List"))
	q := h.client.Category.Query()
	var err error
	page := 1
	if d := r.URL.Query().Get("page"); d != "" {
		page, err = strconv.Atoi(d)
		if err != nil {
			l.Info("error parsing query parameter 'page'", zap.String("page", d), zap.Error(err))
			BadRequest(w, "page must be an integer greater zero")
			return
		}
	}
	itemsPerPage := 30
	if d := r.URL.Query().Get("itemsPerPage"); d != "" {
		itemsPerPage, err = strconv.Atoi(d)
		if err != nil {
			l.Info("error parsing query parameter 'itemsPerPage'", zap.String("itemsPerPage", d), zap.Error(err))
			BadRequest(w, "itemsPerPage must be an integer greater zero")
			return
		}
	}
	es, err := q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage).All(r.Context())
	if err != nil {
		l.Error("error fetching categories from db", zap.Error(err))
		InternalServerError(w, nil)
		return
	}
	l.Info("categories rendered", zap.Int("amount", len(es)))
	easyjson.MarshalToHTTPResponseWriter(NewCategory4094953247Views(es), w)
}

// Read fetches the ent.Collar identified by a given url-parameter from the
// database and returns it to the client.
func (h *CollarHandler) List(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "List"))
	q := h.client.Collar.Query()
	var err error
	page := 1
	if d := r.URL.Query().Get("page"); d != "" {
		page, err = strconv.Atoi(d)
		if err != nil {
			l.Info("error parsing query parameter 'page'", zap.String("page", d), zap.Error(err))
			BadRequest(w, "page must be an integer greater zero")
			return
		}
	}
	itemsPerPage := 30
	if d := r.URL.Query().Get("itemsPerPage"); d != "" {
		itemsPerPage, err = strconv.Atoi(d)
		if err != nil {
			l.Info("error parsing query parameter 'itemsPerPage'", zap.String("itemsPerPage", d), zap.Error(err))
			BadRequest(w, "itemsPerPage must be an integer greater zero")
			return
		}
	}
	es, err := q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage).All(r.Context())
	if err != nil {
		l.Error("error fetching collars from db", zap.Error(err))
		InternalServerError(w, nil)
		return
	}
	l.Info("collars rendered", zap.Int("amount", len(es)))
	easyjson.MarshalToHTTPResponseWriter(NewCollar1522160880Views(es), w)
}

// Read fetches the ent.Owner identified by a given url-parameter from the
// database and returns it to the client.
func (h *OwnerHandler) List(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "List"))
	q := h.client.Owner.Query()
	var err error
	page := 1
	if d := r.URL.Query().Get("page"); d != "" {
		page, err = strconv.Atoi(d)
		if err != nil {
			l.Info("error parsing query parameter 'page'", zap.String("page", d), zap.Error(err))
			BadRequest(w, "page must be an integer greater zero")
			return
		}
	}
	itemsPerPage := 30
	if d := r.URL.Query().Get("itemsPerPage"); d != "" {
		itemsPerPage, err = strconv.Atoi(d)
		if err != nil {
			l.Info("error parsing query parameter 'itemsPerPage'", zap.String("itemsPerPage", d), zap.Error(err))
			BadRequest(w, "itemsPerPage must be an integer greater zero")
			return
		}
	}
	es, err := q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage).All(r.Context())
	if err != nil {
		l.Error("error fetching owners from db", zap.Error(err))
		InternalServerError(w, nil)
		return
	}
	l.Info("owners rendered", zap.Int("amount", len(es)))
	easyjson.MarshalToHTTPResponseWriter(NewOwner139708381Views(es), w)
}

// Read fetches the ent.Pet identified by a given url-parameter from the
// database and returns it to the client.
func (h *PetHandler) List(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "List"))
	q := h.client.Pet.Query()
	var err error
	page := 1
	if d := r.URL.Query().Get("page"); d != "" {
		page, err = strconv.Atoi(d)
		if err != nil {
			l.Info("error parsing query parameter 'page'", zap.String("page", d), zap.Error(err))
			BadRequest(w, "page must be an integer greater zero")
			return
		}
	}
	itemsPerPage := 30
	if d := r.URL.Query().Get("itemsPerPage"); d != "" {
		itemsPerPage, err = strconv.Atoi(d)
		if err != nil {
			l.Info("error parsing query parameter 'itemsPerPage'", zap.String("itemsPerPage", d), zap.Error(err))
			BadRequest(w, "itemsPerPage must be an integer greater zero")
			return
		}
	}
	es, err := q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage).All(r.Context())
	if err != nil {
		l.Error("error fetching pets from db", zap.Error(err))
		InternalServerError(w, nil)
		return
	}
	l.Info("pets rendered", zap.Int("amount", len(es)))
	easyjson.MarshalToHTTPResponseWriter(NewPet359800019Views(es), w)
}