package display

import (
	"fmt"
	"time"

	"github.com/mayankbansal12/yak/internal/model"
)

func PrintTodosWithOptions(todos []model.Todo, opts Options) {
	if opts.Now.IsZero() {
		opts.Now = time.Now()
	}

	if opts.JSON {
		printJSON(todos, opts)
		return
	}

	if len(todos) == 0 {
		fmt.Println("No todos found.")
		return
	}

	if opts.Compact {
		printCompact(todos, opts)
		return
	}

	printCards(todos, opts)
}
