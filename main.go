package main

import (
	"context"
	"log"
	"net/http"

	"github.com/notmyalias/ogent-custom-ent-enum-issue/ent"
	"github.com/notmyalias/ogent-custom-ent-enum-issue/ent/ogent"
)

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	// Run the auto migration tool.
	if err = client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	srv, err := ogent.NewServer(ogent.NewOgentHandler(client))
	if err != nil {
		log.Fatalf("failed creating ogent server: %v", err)
	}

	if err = http.ListenAndServe(":8080", srv); err != nil {
		log.Fatalf("failed to listen and serve: %v", err)
	}
}
