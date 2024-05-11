package main

import (
	"context"

	"github.com/bobbyiliev/sample-go-serverless-func/lib"
)

type Event struct {
	Name string `json:"name"`
}

type Response struct {
	Body string `json:"body"`
}

func Main(ctx context.Context, event Event) Response {
	if event.Name == "" {
		event.Name = "stranger"
	}
	greeting := lib.FormatGreeting(event.Name)
	version := ctx.Value("function_version").(string)
	return Response{
		Body: greeting + "! This is function version " + version,
	}
}
