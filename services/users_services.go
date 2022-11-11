package services

import (
	"../domain"
	"../utils"
)

func CreateUser(user *domain.User) (*domain.User, *utils.RestErr) {
	user, restErr := domain.Create(user)
	if restErr != nil {
		return nil, restErr
	}
	return user, nil
}

func CreateLinkData(linkData *domain.LinkData) (*domain.LinkData, *utils.RestErr) {
	linkData, restErr := domain.CreateLinkData(linkData)
	if restErr != nil {
		return nil, restErr
	}
	return linkData, nil
}

func FindUser(email string) (*domain.User, *utils.RestErr) {
	user, restErr := domain.Find(email)
	if restErr != nil {
		return nil, restErr
	}
	user.Password = ""
	return user, nil
}

func UrlClicked(url string, userPhone string) (*domain.UserLinkData, *utils.RestErr) {
	userLinkData, restErr := domain.UrlClicked(url, userPhone)
	if restErr != nil {
		return nil, restErr
	}
	return userLinkData, nil
}

func FindLinkData(shortUrl string) (*domain.LinkData, *utils.RestErr) {
	linkData, restErr := domain.FindLinkData(shortUrl)
	if restErr != nil {
		return nil, restErr
	}
	return linkData, nil
}

func UpdateUser(email string, field string, value string) (*domain.User, *utils.RestErr) {
	user, restErr := domain.Update(email, field, value)
	if restErr != nil {
		return nil, restErr
	}
	user.Password = ""
	return user, nil
}
