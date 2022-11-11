package domain

import (
	"context"
	"fmt"
	"time"

	"../utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Create(user *User) (*User, *utils.RestErr) {
	usersC := db.Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*20)
	result, err := usersC.InsertOne(ctx, bson.M{
		"name":     user.Name,
		"email":    user.Email,
		"password": user.Password,
	})
	fmt.Println(result)
	fmt.Println(err)
	if err != nil {
		restErr := utils.InternalErr("can't insert user to the database.")
		return nil, restErr
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	user.Password = ""
	return user, nil
}

func CreateLinkData(linkData *LinkData) (*LinkData, *utils.RestErr) {
	usersC := db.Collection("LinkData")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*20)
	result, err := usersC.InsertOne(ctx, bson.M{
		"url":      linkData.Url,
		"shortUrl": linkData.ShortUrl,
		"useCase":  linkData.UseCase,
	})
	fmt.Println(err)
	if err != nil {
		restErr := utils.InternalErr("can't insert linkData to the database.")
		return nil, restErr
	}
	linkData.ID = result.InsertedID.(primitive.ObjectID)
	return linkData, nil
}

func CreateUserLinkData(url string, userPhone string) (*UserLinkData, *utils.RestErr) {
	userLinkDatasC := db.Collection("UserLinkData")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*20)
	result, err := userLinkDatasC.InsertOne(ctx, bson.M{
		"url":        url,
		"userPhone":  userPhone,
		"clickCount": 1,
	})
	if err != nil {
		restErr := utils.InternalErr("can't insert userLinkData to the database.")
		return nil, restErr
	}
	var userLinkData UserLinkData
	errFind := userLinkDatasC.FindOne(ctx, bson.M{"url": url, "userPhone": userPhone}).Decode(&userLinkData)
	if errFind != nil {
		restErr := utils.NotFound("userLinkData not found.")
		return nil, restErr
	}
	userLinkData.ID = result.InsertedID.(primitive.ObjectID)
	return &userLinkData, nil
}

func Find(email string) (*User, *utils.RestErr) {
	var user User
	usersC := db.Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	err := usersC.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		restErr := utils.NotFound("user not found.")
		return nil, restErr
	}
	return &user, nil
}

func UrlClicked(url string, userPhone string) (*UserLinkData, *utils.RestErr) {
	var userLinkData UserLinkData
	userLinkDatasC := db.Collection("UserLinkData")

	fmt.Println(url)
	fmt.Println(userPhone)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	err := userLinkDatasC.FindOne(ctx, bson.M{"url": url, "userPhone": userPhone}).Decode(&userLinkData)

	fmt.Println(userLinkData)
	if err != nil {
		CreateUserLinkData(url, userPhone)
	} else {
		UpdateUserLinkData(&userLinkData)
	}
	return &userLinkData, nil
}

func FindLinkData(shortUrl string) (*LinkData, *utils.RestErr) {
	var linkData LinkData
	linkDatasC := db.Collection("LinkData")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	err := linkDatasC.FindOne(ctx, bson.M{"shortUrl": shortUrl}).Decode(&linkData)
	if err != nil {
		restErr := utils.NotFound("LinkData not found.")
		return nil, restErr
	}
	// fmt.Println(shortUrl)
	// fmt.Println(linkData)
	return &linkData, nil
}

func FindUserLinkData(url string, userPhone string) (*UserLinkData, *utils.RestErr) {
	var userLinkData UserLinkData
	userLinkDatasC := db.Collection("UserLinkData")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	err := userLinkDatasC.FindOne(ctx, bson.M{"url": url, "userPhone": userPhone}).Decode(&userLinkData)
	if err != nil {
		restErr := utils.NotFound("UserLinkData not found.")
		return nil, restErr
	}
	return &userLinkData, nil
}

func UpdateUserLinkData(userLinkData *UserLinkData) (*UserLinkData, *utils.RestErr) {
	UserLinkDatasC := db.Collection("UserLinkData")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	result, err := UserLinkDatasC.UpdateOne(ctx, bson.M{"url": userLinkData.Url, "userPhone": userLinkData.UserPhone}, bson.M{"$set": bson.M{"clickCount": userLinkData.ClickCount + 1}})
	if err != nil {
		restErr := utils.InternalErr("can not update.")
		return nil, restErr
	}
	if result.MatchedCount == 0 {
		restErr := utils.NotFound("userLinkData not found.")
		return nil, restErr
	}
	if result.ModifiedCount == 0 {
		restErr := utils.BadRequest("no such field")
		return nil, restErr
	}
	userLinkData, restErr := FindUserLinkData(userLinkData.Url, userLinkData.UserPhone)
	if restErr != nil {
		return nil, restErr
	}
	return userLinkData, restErr
}

func Update(email string, field string, value string) (*User, *utils.RestErr) {
	usersC := db.Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	result, err := usersC.UpdateOne(ctx, bson.M{"email": email}, bson.M{"$set": bson.M{field: value}})
	if err != nil {
		restErr := utils.InternalErr("can not update.")
		return nil, restErr
	}
	if result.MatchedCount == 0 {
		restErr := utils.NotFound("user not found.")
		return nil, restErr
	}
	if result.ModifiedCount == 0 {
		restErr := utils.BadRequest("no such field")
		return nil, restErr
	}
	user, restErr := Find(email)
	if restErr != nil {
		return nil, restErr
	}
	return user, restErr
}
