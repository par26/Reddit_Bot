package main

func GetPost(postType string, subreddit string, num int) ([]string, error) {
	posts, err := findPost(postType, subreddit, num)

	var output []string

	if err != nil {
		return nil, err
	}

	for _, post := range posts {
		output = append(output, post.Title)
	}

	return output, nil
}
