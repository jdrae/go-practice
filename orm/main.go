package main

import (
	"context"
	"log"

	"orm/ent"
	"orm/ent/user"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// create a new ent.Client
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil{
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	//auto migration
	if err := client.Schema.Create(context.Background()); err != nil{
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error){
	u, err := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		Save(ctx)
	if err != nil{
		return nil, fmt.Errorf("failed creating user: %v", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func QueryUser(ctx context.Context, client *ent.Client) (*ent.User, error){
	u, err := client.User.
		Query().
		Where(user.NameEQ("a8m")).
		Only(ctx) // fails if no user found or more than 1 user returned
	if err != nil{
		return nil, fmt.Errorf("failed querying user: %v", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}

