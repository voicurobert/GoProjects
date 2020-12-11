package repository

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/voicurobert/GoProjects/rest_api/entity"
	iterator "google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type repo struct {
}

//NewFirestoreRepository constructor
func NewFirestoreRepository() PostRespository {
	return &repo{}
}

const (
	projectID      = "rest-api-98457"
	collectionName = "posts"
	certPath       = "/Users/robert/go/src/github.com/voicurobert/GoProjects/rest_api/rest-api-98457-firebase-adminsdk-jlz5d-666b0ecde5.json"
)

//Save a Post struct
func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID, option.WithCredentialsFile(certPath))
	if err != nil {
		log.Fatalf("Failed to create a Firestore client: %v", err)
		return nil, err
	}
	defer client.Close()
	/*
		_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
			"ID":    post.ID,
			"Title": post.Title,
			"Text":  post.Text,
		})
	*/
	_, _, err = client.Collection(collectionName).Add(ctx, post)

	if err != nil {
		log.Fatalf("Failed to add a new post: %v", err)
		return nil, err
	}

	return post, nil
}

func (*repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID, option.WithCredentialsFile(certPath))
	if err != nil {
		log.Fatalf("Failed to create a Firestore client: %v", err)
		return nil, err
	}
	defer client.Close()
	var posts []entity.Post
	iter := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate the list of Posts: %v", err)
			return nil, err
		}

		post := entity.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}
	return posts, nil
}
