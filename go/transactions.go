package main

import (
	"errors"
	"fmt"
)

type Subscription struct {
	ID               string `json:"id"`
	SubscriptionPlan string `json:"subscription_plan"`
}

type User struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	Subscription string `json:"subscription"`
}

type SubscriptionPlan struct {
	ID     string `json:"id"`
	Amount int32  `json:"amount"`
	Name   string `json:"name"`
}

func (u *User) getById(id string) error {
	return nil
}

func (u *User) update() error {
	return nil
}

func (sp *SubscriptionPlan) getById(id string) error {
	return nil
}

func (s *Subscription) getById(id string) error {
	return nil
}

func (s *Subscription) update() error {
	return nil
}

func (s *Subscription) createNew(sp SubscriptionPlan) error {
	s.ID = "012412"
	s.SubscriptionPlan = sp.ID
	return nil
}

func purchasePlan(userId, planId string) error {
	// make sure there is a connection between database and the application.
	// make sure the app has enough permissions to write on the database.
	if err := verifyConnAndPerm("admin"); err != nil {
		return err
	}

	// get the user
	user, plan := User{}, SubscriptionPlan{}
	if err := user.getById(userId); err != nil {
		// if the user is not found then quit the transaction.
		return err
	}
	// get the subscription plan
	if err := plan.getById(planId); err != nil {
		// if the plan is not found then quit the transaction.
		return err
	}
	// create a new Subscription
	userSub := Subscription{}
	if err := userSub.createNew(plan); err != nil {
		return err
	}
	// update the user
	user.Subscription = userSub.ID
	if err := user.update(); err != nil {
		return err
	}
	return nil
}

func verifyConnAndPerm(permLevel string) error {
	permissions := []string{"admin", "read", "write", "delete"}
	if includes(permissions, permLevel) {
		return nil
	}
	return errors.New(fmt.Sprintf("you do not have '%v' permission on the database.", permLevel))
}

func includes[T comparable](arr []T, item T) bool {
	for _, v := range arr {
		if v == item {
			return true
		}
	}
	return false
}
