package main

import (
	"context"

	"github.com/vartanbeno/go-reddit/v2/reddit"
)

func findPost(postType string, subreddit string, num int) ([]*reddit.Post, error) {

	var ctx = context.Background()

	switch postType {
	case "New":
		posts, err := getNewPost(ctx, subreddit, num)
		return posts, err
	case "Hot":
		posts, err := getHotPost(ctx, subreddit, num)
		return posts, err
	case "Top":
		posts, err := getTopPost(ctx, subreddit, num)
		return posts, err
	case "Rising":
		posts, err := getRisingPost(ctx, subreddit, num)
		return posts, err
	}

	return nil, nil
}

func getTopPost(ctx context.Context, subreddit string, num int) ([]*reddit.Post, error) {

	posts, _, err := reddit.DefaultClient().Subreddit.TopPosts(ctx, subreddit, &reddit.ListPostOptions{
		ListOptions: reddit.ListOptions{
			Limit: num,
		},
		Time: "all",
	})
	if err != nil {
		return nil, err
	}

	return posts, nil

}

func getHotPost(ctx context.Context, subreddit string, num int) ([]*reddit.Post, error) {

	posts, _, err := reddit.DefaultClient().Subreddit.TopPosts(ctx, subreddit, &reddit.ListPostOptions{
		ListOptions: reddit.ListOptions{
			Limit: num,
		},
		Time: "all",
	})
	if err != nil {
		return nil, err
	}

	return posts, nil

}

func getRisingPost(ctx context.Context, subreddit string, num int) ([]*reddit.Post, error) {

	posts, _, err := reddit.DefaultClient().Subreddit.TopPosts(ctx, subreddit, &reddit.ListPostOptions{
		ListOptions: reddit.ListOptions{
			Limit: num,
		},
		Time: "all",
	})
	if err != nil {
		return nil, err
	}

	return posts, nil

}

func getNewPost(ctx context.Context, subreddit string, num int) ([]*reddit.Post, error) {

	posts, _, err := reddit.DefaultClient().Subreddit.TopPosts(ctx, subreddit, &reddit.ListPostOptions{
		ListOptions: reddit.ListOptions{
			Limit: num,
		},
		Time: "all",
	})
	if err != nil {
		return nil, err
	}

	return posts, nil

}
