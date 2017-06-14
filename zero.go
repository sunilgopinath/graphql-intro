// Package starwars provides a example schema and resolver based on Star Wars characters.
//
// Source: https://github.com/graphql/graphql.github.io/blob/source/site/_core/swapiSchema.js
package starwars

import graphql "github.com/neelance/graphql-go"

//Schema is the graphql schema
var Schema = `
	schema {
		query: Query
	}
	# The query type, represents all of the entry points into our object graph
	type Query {
		person(id: ID!): Person
	}
	type Person {
		id: ID!
		first_name: String!
		last_name: String!
		username: String!
		email: String!
		friends: [Person]
	}
`

type person struct {
	ID        graphql.ID
	FirstName string
	LastName  string
	Username  string
	Email     string
	Friends   []graphql.ID
}

var persons = []*person{
	{
		ID:        "1000",
		FirstName: "Luke",
		LastName:  "Skywalker",
		Username:  "lskywalker",
		Email:     "lskywalker@facebook.com",
		Friends:   []graphql.ID{"1002", "1003", "2000", "2001"},
	},
	{
		ID:        "1001",
		FirstName: "Darth",
		LastName:  "Vader",
		Username:  "dvader",
		Email:     "dvader@facebook.com",
		Friends:   []graphql.ID{"1004"},
	},
	{
		ID:        "1002",
		FirstName: "Han",
		LastName:  "Solo",
		Username:  "hsolo",
		Email:     "hsolo@facebook.com",
		Friends:   []graphql.ID{"1000", "1003", "2001"},
	},
	{
		ID:        "1002",
		FirstName: "Han",
		LastName:  "Solo",
		Username:  "hsolo",
		Email:     "hsolo@facebook.com",
		Friends:   []graphql.ID{"1000", "1003"},
	},
	{
		ID:        "1003",
		FirstName: "Leia",
		LastName:  "Organa",
		Username:  "lorgana",
		Email:     "lorgana@facebook.com",
		Friends:   []graphql.ID{"1000", "1002"},
	},
	{
		ID:        "1004",
		FirstName: "Wilhuff",
		LastName:  "Tarkin",
		Username:  "wtarkin",
		Email:     "wtarkin@facebook.com",
		Friends:   []graphql.ID{"1001"},
	},
}

var personData = make(map[graphql.ID]*person)

func init() {
	for _, p := range persons {
		personData[p.ID] = p
	}
}

//Resolver maps the schema to go
type Resolver struct{}

//Person represents the Person type
func (r *Resolver) Person(args struct{ ID graphql.ID }) *personResolver {
	if p := personData[args.ID]; p != nil {
		return &personResolver{p}
	}
	return nil
}

type personResolver struct {
	p *person
}

func (r *personResolver) ID() graphql.ID {
	return r.p.ID
}

func (r *personResolver) FirstName() string {
	return r.p.FirstName
}

func (r *personResolver) LastName() string {
	return r.p.LastName
}

func (r *personResolver) Username() string {
	return r.p.Username
}

func (r *personResolver) Email() string {
	return r.p.Email
}

func (r *personResolver) Friends() *[]*personResolver {
	return resolvePersons(r.p.Friends)
}

func resolvePersons(ids []graphql.ID) *[]*personResolver {
	var persons []*personResolver
	for _, id := range ids {
		if c := resolvePerson(id); c != nil {
			persons = append(persons, c)
		}
	}
	return &persons
}

func resolvePerson(id graphql.ID) *personResolver {
	if p, ok := personData[id]; ok {
		return &personResolver{p}
	}
	return nil
}
