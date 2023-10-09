/**
    package: Golang-lesson
    filename: main
    author: diogo
    time: 2020/11/17 18:02
**/
package main

import (
	"context"
	"fmt"
)

type User struct {
	name string
}
type MovieService interface {
	 
	Update(ctx context.Context, id string,) error
	Delete(ctx context.Context, id string) error
}

type movieService struct {
	C string
}

func (m movieService) Update(ctx context.Context, id string) error {
	panic("implement me")
}

func (m movieService) Delete(ctx context.Context, id string) error {
	panic("implement me")
}

func main()  {

	var un  = (*User)(nil)
	fmt.Println(un)

	var mo MovieService = (*movieService)(nil)
	
	fmt.Println(mo)
}
