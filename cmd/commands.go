package main

import (
	"encoding/json"
	"github.com/alleswebdev/go-command-executor/internal/command"
	"github.com/alleswebdev/go-command-executor/internal/config"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type CommandsResource struct {
	cfg         config.Config
	commandsMap map[string]command.Command
}

type Commander interface {
	GetName() string
	Start() (string, error)
	Stop() (string, error)
	Restart() (string, error)
}

func New(cfg config.Config, cm map[string]command.Command) CommandsResource {
	return CommandsResource{cfg: cfg, commandsMap: cm}
}

func (rs CommandsResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/all", rs.List) // GET /command - read a list of command
	r.Post("/", rs.Create) // POST /command - create a new command and persist it
	r.Put("/", rs.Delete)

	r.Route("/{name}", func(r chi.Router) {
		r.Get("/", rs.Get)       // GET /command/{id} - read a single command by :id
		r.Get("/exec", rs.Exec)  // GET /command/{id} - read a single command by :id
		r.Put("/", rs.Update)    // PUT /command/{id} - update a single command by :id
		r.Delete("/", rs.Delete) // DELETE /command/{id} - delete a single command by :id
	})

	return r
}

func (rs CommandsResource) List(w http.ResponseWriter, r *http.Request) {
	cg, err := json.Marshal(rs.cfg.Commands)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(200)
	w.Write(cg)
}

func (rs CommandsResource) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("command create"))
}

func (rs CommandsResource) Exec(w http.ResponseWriter, r *http.Request) {
	if name := chi.URLParam(r, "name"); name != "" {
		if cmd, ok := rs.commandsMap[name]; ok {
			result, err := cmd.Start()
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}

			w.Write([]byte(result))
			return
		}
	}
	w.Write([]byte("command not found"))
}

func (rs CommandsResource) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("command get"))
}

func (rs CommandsResource) Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("command update"))
}

func (rs CommandsResource) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("command delete"))
}
