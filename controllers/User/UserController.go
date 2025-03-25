package controllers

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/Pedrooaj/Api_rest-Go/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	


)

type User struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nome string `bson:"nome" json:"nome" binding:"required"` 
	Idade int `bson:"idade" json:"idade" binding:"required,min=12"`
	DataCriacao time.Time `bson:"dataCriacao" json:"dataCriacao"`
}



func InsertUser(c *gin.Context){
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Dados inválidos", "details": "São necessarios os seguintes campos Nome e Idade..."})
		return
	}

	user.DataCriacao = time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := database.Collection().InsertOne(ctx, user)
	if err != nil {
		c.JSON(500, gin.H{"error": "Erro ao inserir usuário" })
		return
	}

	c.JSON(200, gin.H{"message": "Usuário inserido com sucesso"})
}


func ListUsers(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	
	cursor, err := database.Collection().Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
		return
	}

	defer cursor.Close(ctx)


	var values []bson.M
	for cursor.Next(ctx) {
		var value bson.M
		if err := cursor.Decode(&value); err != nil {
			c.JSON(500, gin.H{"error": "Erro ao decodificar dados"})
			return
		}
		values = append(values, value)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(500, gin.H{"error": "Erro no cursor"}) // não iterou os documentos corretamente
		return
	}

	c.JSON(200, values)
}


func DeleteUser(c *gin.Context){
	id := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		c.JSON(400, gin.H{"erro": "Id inválido, usuário não existe"})
		return
	}

	result := database.Collection().FindOneAndDelete(ctx, bson.M{"_id": objectId})

	if result.Err() != nil {
		c.JSON(500, gin.H{"error": "Erro ao encontrar usuário"})
		return
	}

	var deletedUser bson.M
	if err := result.Decode(&deletedUser); err != nil{
		c.JSON(500, gin.H{"error": "Erro ao decodificar usuário deletado"})
		return
	}
	c.JSON(200, deletedUser)
}


func GetUser(c *gin.Context){
	id := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		c.JSON(400, gin.H{"erro": "Id inválido"})
		return
	}


	result := database.Collection().FindOne(ctx, bson.M{"_id": objectID})

	if result.Err() != nil{
		c.JSON(500, gin.H{"error": "Erro ao buscar usuário"})
		return
	}

	var user bson.M
	if err := result.Decode(&user); err != nil {
		c.JSON(500, gin.H{"error": "Erro ao decodificar usuário"})
		return
	}

	c.JSON(200, user)
}
