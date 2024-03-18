package testdata

import "github.com/KakinokiKanta/go-intermediate/models"

var ArticleTestData = []models.Article{
	models.Article{
		// 記事IDが1番の記事内容
		ID: 1,
		Title: "firstPost",
		Contents: "This is my first blog",
		UserName: "saki",
		NiceNum: 2,
	},
	{
		// 記事IDが2番の記事内容
		ID: 2,
		Title: "2nd",
		Contents: "Second blog post",
		UserName: "saki",
		NiceNum: 4,
	},
}
