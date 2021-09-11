// Code generated by entc, DO NOT EDIT.

package http

import (
	"net/http"

	"github.com/mailru/easyjson"
	"github.com/masseelch/elk/internal/fridge/ent"
	"github.com/masseelch/elk/internal/fridge/ent/compartment"
	"github.com/masseelch/elk/internal/fridge/ent/fridge"
	"github.com/masseelch/elk/internal/fridge/ent/item"
	"go.uber.org/zap"
)

// Create creates a new ent.Compartment and stores it in the database.
func (h CompartmentHandler) Create(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Create"))
	// Get the post data.
	var d CompartmentCreateRequest
	if err := easyjson.UnmarshalFromReader(r.Body, &d); err != nil {
		l.Error("error decoding json", zap.Error(err))
		BadRequest(w, "invalid json string")
		return
	}
	// Save the data.
	b := h.client.Compartment.Create()
	if d.Name != nil {
		b.SetName(*d.Name)
	}
	if d.Fridge != nil {
		b.SetFridgeID(*d.Fridge)
	}
	if d.Contents != nil {
		b.AddContentIDs(d.Contents...)
	}
	e, err := b.Save(r.Context())
	if err != nil {
		switch {
		default:
			l.Error("could not create compartment", zap.Error(err))
			InternalServerError(w, nil)
		}
		return
	}
	// Reload entry.
	q := h.client.Compartment.Query().Where(compartment.ID(e.ID))
	e, err = q.Only(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int("id", e.ID))
			NotFound(w, msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.Int("id", e.ID))
			BadRequest(w, msg)
		default:
			l.Error("could not read compartment", zap.Error(err), zap.Int("id", e.ID))
			InternalServerError(w, nil)
		}
		return
	}
	l.Info("compartment rendered", zap.Int("id", e.ID))
	easyjson.MarshalToHTTPResponseWriter(NewCompartment1151786357View(e), w)
}

// Create creates a new ent.Fridge and stores it in the database.
func (h FridgeHandler) Create(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Create"))
	// Get the post data.
	var d FridgeCreateRequest
	if err := easyjson.UnmarshalFromReader(r.Body, &d); err != nil {
		l.Error("error decoding json", zap.Error(err))
		BadRequest(w, "invalid json string")
		return
	}
	// Save the data.
	b := h.client.Fridge.Create()
	if d.Title != nil {
		b.SetTitle(*d.Title)
	}
	if d.Compartments != nil {
		b.AddCompartmentIDs(d.Compartments...)
	}
	e, err := b.Save(r.Context())
	if err != nil {
		switch {
		default:
			l.Error("could not create fridge", zap.Error(err))
			InternalServerError(w, nil)
		}
		return
	}
	// Reload entry.
	q := h.client.Fridge.Query().Where(fridge.ID(e.ID))
	e, err = q.Only(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int("id", e.ID))
			NotFound(w, msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.Int("id", e.ID))
			BadRequest(w, msg)
		default:
			l.Error("could not read fridge", zap.Error(err), zap.Int("id", e.ID))
			InternalServerError(w, nil)
		}
		return
	}
	l.Info("fridge rendered", zap.Int("id", e.ID))
	easyjson.MarshalToHTTPResponseWriter(NewFridge2716213877View(e), w)
}

// Create creates a new ent.Item and stores it in the database.
func (h ItemHandler) Create(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Create"))
	// Get the post data.
	var d ItemCreateRequest
	if err := easyjson.UnmarshalFromReader(r.Body, &d); err != nil {
		l.Error("error decoding json", zap.Error(err))
		BadRequest(w, "invalid json string")
		return
	}
	// Save the data.
	b := h.client.Item.Create()
	if d.Name != nil {
		b.SetName(*d.Name)
	}
	if d.Compartment != nil {
		b.SetCompartmentID(*d.Compartment)
	}
	e, err := b.Save(r.Context())
	if err != nil {
		switch {
		default:
			l.Error("could not create item", zap.Error(err))
			InternalServerError(w, nil)
		}
		return
	}
	// Reload entry.
	q := h.client.Item.Query().Where(item.ID(e.ID))
	e, err = q.Only(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int("id", e.ID))
			NotFound(w, msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.Int("id", e.ID))
			BadRequest(w, msg)
		default:
			l.Error("could not read item", zap.Error(err), zap.Int("id", e.ID))
			InternalServerError(w, nil)
		}
		return
	}
	l.Info("item rendered", zap.Int("id", e.ID))
	easyjson.MarshalToHTTPResponseWriter(NewItem1509516544View(e), w)
}
