package main

import (
	"fmt"
	"golang-graphql-gorm/controllers"
	"golang-graphql-gorm/db"
	"golang-graphql-gorm/models"
	"net/http"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main(){
	db.Run()
	fmt.Println("golang-graphql-gorm")


	carType:= graphql.NewObject(graphql.ObjectConfig{
		Name: "Car",
		Fields : graphql.Fields{
			"id": &graphql.Field{Type: graphql.ID},
			"name" : &graphql.Field{Type: graphql.String},
			"size" : &graphql.Field{Type: graphql.String},
			"created_at": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if car, ok := p.Source.(*models.Car); ok {
						return car.CreatedAt, nil
					}
					return nil, nil
				},
			},
		},
	})

	rootQuery:= graphql.NewObject(graphql.ObjectConfig{

		Name: "RootQuery",
		Fields: graphql.Fields{
			"car" : &graphql.Field{
				Type: graphql.NewList(carType),
				Description: "GET CAR LIST",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, _:= p.Args["id"].(int)
					listCar, err:= controllers.GetCar(id)
					return listCar, err
				},
			},
		},
	})
	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"createCar": &graphql.Field{
				Type:        carType,
				Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"size": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				name, _ := p.Args["name"].(string)
				size, _ := p.Args["size"].(string)
				createdAt := time.Now()

				carData := &models.Car{
					Name: name,
					Size: size,
					CreatedAt: createdAt,
				}

				// newCar := controllers.CreateCar(carData)
				// return carData, newCar.Error
				// return newCar, err

				// newCar := db.DB.Debug().Create(carData).Error
			
				// return newCar
				newCar := db.DB.Debug().Create(carData)
				return carData, newCar.Error
			},
		},
		"updateCar": &graphql.Field{
			Type:        carType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"size": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id, _ := params.Args["id"].(int)
				name, _ := params.Args["name"].(string)
				size, _ := params.Args["size"].(string)

				

				carData := &models.Car{
					ID:    id,
					Name:  name,
					Size:  size,
				}

				updateCar := db.DB.Debug().Save(carData)
			
				return carData, updateCar.Error
			},
		},
		"deleteCar": &graphql.Field{
			Type:        carType,
			Description: "Delete an car",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, _ := p.Args["id"].(int)
				carId, err:= controllers.DeleteCar(id)
					return carId, err
			},
		},
	},
})


	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})
	http.Handle("/graphql", h)
	http.ListenAndServe(":8080", nil)









}