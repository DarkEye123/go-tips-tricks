package main

import "fmt"

type User struct {
	name          string
	subscriptions []*Subscription
}

type Pool struct {
	users []*User
}

type Subscription struct {
	name string
}

func (u *User) GetSubscription(name string) *Subscription {
	if u == nil {
		return nil // due to chain u can be nil
	}
	for _, sub := range u.subscriptions {
		if sub.name == name {
			return sub
		}
	}
	return nil
}

func (p *Pool) FindUser(name string) *User {
	for _, user := range p.users {
		if user.name == name {
			return user
		}
	}
	return nil
}

func main() {
	u := &User{name: "Adam", subscriptions: []*Subscription{&Subscription{name: "pingdom"}, &Subscription{name: "appoptics"}}}
	pool := &Pool{users: []*User{u}}
	fmt.Println(pool.FindUser("Adam").GetSubscription("pingdom").name)
}
