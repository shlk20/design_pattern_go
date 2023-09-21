package iterator

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	user1 := &User{
		name: "a",
		age:  20,
	}
	user2 := &User{
		name: "b",
		age:  25,
	}

	userCollection := &UserCollection{
		users: []*User{user1, user2},
	}

	iterator := userCollection.createIterator()

	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("User is %+v\n", user)
	}
}
