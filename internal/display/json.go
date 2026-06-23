package display

import (
	"encoding/json"
	"os"
	"time"

	"github.com/mayankbansal12/yak/internal/model"
)

type listEnvelope struct {
	Todos       []model.Todo `json:"todos"`
	Count       int          `json:"count"`
	GeneratedAt string       `json:"generated_at"`
}

type singleEnvelope struct {
	Todo model.Todo `json:"todo"`
	ID   int        `json:"id"`
}

func printJSON(todos []model.Todo, opts Options) {
	enc := json.NewEncoder(os.Stdout)
	_ = enc.Encode(listEnvelope{
		Todos:       todos,
		Count:       len(todos),
		GeneratedAt: opts.Now.UTC().Format(time.RFC3339),
	})
}

func PrintTodoJSON(t model.Todo, now time.Time) {
	enc := json.NewEncoder(os.Stdout)
	_ = enc.Encode(singleEnvelope{Todo: t, ID: t.ID})
}
