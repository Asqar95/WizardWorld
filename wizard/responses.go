package wizard

import (
	"fmt"
)

type ElixirsResponse struct {
	Elixir []Elixir
}

type ElixirResponse struct {
	Elixir Elixir
}

type Ingredients struct {
	id   string `json:"id"`
	name string `json:"name"`
}

type Inventors struct {
	id        string `json:"id"`
	firstName string `json:"firstName"`
	lastName  string `json:"lastName"`
}

type Elixir struct {
	id              string        `json:"id"`
	name            string        `json:"name"`
	effect          string        `json:"effect"`
	sideEffects     string        `json:"sideEffects"`
	characteristics string        `json:"characteristics"`
	time            string        `json:"time"`
	difficulty      string        `json:"difficulty"`
	ingredients     []Ingredients `json:"ingredients"`
	inventors       []Inventors   `json:"inventors"`
	manufacturer    string        `json:"manufacturer"`
}

func (e Elixir) Info() string {
	return fmt.Sprintf("[ID] %s | [NAME] %s | [EFFECT] %s | [SIDEEFFECT] %s | [characteristics] %s", e.id, e.name, e.effect, e.sideEffects, e.characteristics)
}
