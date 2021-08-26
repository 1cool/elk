// Code generated by entc, DO NOT EDIT.

package http

import (
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/masseelch/elk/internal/unit/pets/ent"
	"go.uber.org/zap"
)

// handler has some convenience methods used on node-handlers.
type handler struct{}

// Bitmask to configure which routes to register.
type Routes uint32

func (rs Routes) has(r Routes) bool { return rs&r != 0 }

const (
	CategoryCreate Routes = 1 << iota
	CategoryRead
	CategoryUpdate
	CategoryDelete
	CategoryList
	CategoryPets
	CategoryRoutes = 1<<iota - 1
)

// CategoryHandler handles http crud operations on ent.Category.
type CategoryHandler struct {
	handler

	client *ent.Client
	log    *zap.Logger
}

func NewCategoryHandler(c *ent.Client, l *zap.Logger) *CategoryHandler {
	return &CategoryHandler{
		client: c,
		log:    l.With(zap.String("handler", "CategoryHandler")),
	}
}

// RegisterHandlers registers the generated handlers on the given chi router.
func (h *CategoryHandler) Mount(r chi.Router, rs Routes) {
	if rs.has(CategoryCreate) {
		r.Post("/", h.Create)
	}
	if rs.has(CategoryRead) {
		r.Get("/{id}", h.Read)
	}
	if rs.has(CategoryUpdate) {
		r.Patch("/{id}", h.Update)
	}
	if rs.has(CategoryDelete) {
		r.Delete("/{id}", h.Delete)
	}
	if rs.has(CategoryList) {
		r.Get("/", h.List)
	}
	if rs.has(CategoryPets) {
		r.Get("/{id}/pets", h.Pets)
	}
}

const (
	OwnerCreate Routes = 1 << iota
	OwnerRead
	OwnerUpdate
	OwnerDelete
	OwnerList
	OwnerPets
	OwnerRoutes = 1<<iota - 1
)

// OwnerHandler handles http crud operations on ent.Owner.
type OwnerHandler struct {
	handler

	client *ent.Client
	log    *zap.Logger
}

func NewOwnerHandler(c *ent.Client, l *zap.Logger) *OwnerHandler {
	return &OwnerHandler{
		client: c,
		log:    l.With(zap.String("handler", "OwnerHandler")),
	}
}

// RegisterHandlers registers the generated handlers on the given chi router.
func (h *OwnerHandler) Mount(r chi.Router, rs Routes) {
	if rs.has(OwnerCreate) {
		r.Post("/", h.Create)
	}
	if rs.has(OwnerRead) {
		r.Get("/{id}", h.Read)
	}
	if rs.has(OwnerUpdate) {
		r.Patch("/{id}", h.Update)
	}
	if rs.has(OwnerDelete) {
		r.Delete("/{id}", h.Delete)
	}
	if rs.has(OwnerList) {
		r.Get("/", h.List)
	}
	if rs.has(OwnerPets) {
		r.Get("/{id}/pets", h.Pets)
	}
}

const (
	PetCreate Routes = 1 << iota
	PetRead
	PetUpdate
	PetDelete
	PetList
	PetCategory
	PetOwner
	PetFriends
	PetRoutes = 1<<iota - 1
)

// PetHandler handles http crud operations on ent.Pet.
type PetHandler struct {
	handler

	client *ent.Client
	log    *zap.Logger
}

func NewPetHandler(c *ent.Client, l *zap.Logger) *PetHandler {
	return &PetHandler{
		client: c,
		log:    l.With(zap.String("handler", "PetHandler")),
	}
}

// RegisterHandlers registers the generated handlers on the given chi router.
func (h *PetHandler) Mount(r chi.Router, rs Routes) {
	if rs.has(PetCreate) {
		r.Post("/", h.Create)
	}
	if rs.has(PetRead) {
		r.Get("/{id}", h.Read)
	}
	if rs.has(PetUpdate) {
		r.Patch("/{id}", h.Update)
	}
	if rs.has(PetDelete) {
		r.Delete("/{id}", h.Delete)
	}
	if rs.has(PetList) {
		r.Get("/", h.List)
	}
	if rs.has(PetCategory) {
		r.Get("/{id}/category", h.Category)
	}
	if rs.has(PetOwner) {
		r.Get("/{id}/owner", h.Owner)
	}
	if rs.has(PetFriends) {
		r.Get("/{id}/friends", h.Friends)
	}
}

func stripEntError(err error) string {
	return strings.TrimPrefix(err.Error(), "ent: ")
}

func zapFields(errs map[string]string) []zap.Field {
	if errs == nil || len(errs) == 0 {
		return nil
	}
	r := make([]zap.Field, 0)
	for k, v := range errs {
		r = append(r, zap.String(k, v))
	}
	return r
}
