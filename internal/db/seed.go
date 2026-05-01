package db

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/Verifieddanny/go-social/internal/store"
)

var usernames = []string{
	"tech_guru", "travel_bug", "foodie_vibes", "backend_wizard", "pixel_perfect",
	"nature_lover", "code_ninja", "sunny_dayz", "urban_explorer", "data_junkie",
}

var titles = []string{
	"10 Tips for Better Go Code",
	"My Trip to the Swiss Alps",
	"Why I Switched to Neovim",
	"The Best Ramen in Tokyo",
	"Understanding PostgreSQL Indexes",
	"Building a Fintech App in 2026",
	"Morning Coffee Rituals",
	"Top 5 UI Trends This Year",
	"How to Handle API Rate Limiting",
	"Exploring the Amazon Rainforest",
}

var contents = []string{
	"Writing clean code in Go requires a deep understanding of interfaces and composition...",
	"The mountains were covered in a thick blanket of snow, and the air was crisp...",
	"I realized that my productivity soared once I stopped reaching for the mouse...",
	"Hidden in a small alleyway in Shinjuku, this ramen shop serves the richest broth...",
	"If your queries are slow, the first place you should look is your indexing strategy...",
	"Security is the number one priority when handling user transactions and private keys...",
	"There is something magical about the smell of freshly ground beans at 6:00 AM...",
	"Minimalism is making a huge comeback in mobile app design this quarter...",
	"When your service grows, you'll eventually hit limits. Here is how to scale gracefully...",
	"The biodiversity in this part of the world is truly breathtaking and worth protecting...",
}

var tags = []string{
	"golang", "programming", "adventure", "productivity", "database",
	"fintech", "startup", "ui-design", "backend", "nature",
}

var comments = []string{
	"Great article! I never thought about interfaces that way.",
	"Adding this to my bucket list for next year.",
	"I'm still struggling with the keybindings, any tips?",
	"I've been there! The spicy miso is incredible.",
	"Thanks for the explanation, my query time dropped by 50%.",
	"Very insightful. How do you handle encryption at rest?",
	"Coffee is life. Totally agree.",
	"The dark mode trend is still my favorite part of this year.",
	"Do you recommend Redis for rate limiting or just in-memory?",
	"Amazing photos! What camera did you use?",
}

func Seed(store store.Storage) {
	ctx := context.Background()

	users := generateUsers(100)
	for _, user := range users {
		if err := store.Users.Create(ctx, user); err != nil {
			log.Println("Error creating user:", err)
			return
		}
	}

	posts := generatePosts(200, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error creating post:", err)
			return
		}
	}

	comments := generateComment(500, users, posts)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("Error creating comment:", err)
			return
		}
	}

	log.Println("Seeding complete")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@examples.com",
			Password: "123123",
		}
	}

	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)

	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]

		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: contents[rand.Intn(len(contents))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}

	return posts
}

func generateComment(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	cms := make([]*store.Comment, num)
	for i := 0; i < num; i++ {
		cms[i] = &store.Comment{
			PostID:  posts[rand.Intn(len(posts))].ID,
			UserID:  users[rand.Intn(len(users))].ID,
			Content: contents[rand.Intn(len(comments))],
		}
	}
	return cms
}
