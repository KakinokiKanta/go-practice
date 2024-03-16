package repositories

import (
	"database/sql"
	"fmt"

	"github.com/KakinokiKanta/go-intermediate/models"
)

const articleNumPerPage = 5

/*
	新規投稿をデータベースにinsertする関数
*/
func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	// クエリの定義
	const sqlStr = `
		insert into articles (title, contents, username, nice, created_at) values (?, ?, ?, 0, now());
	`
	// sql.DB型のメソッドExecを用いて、クエリを実行
	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	if err != nil {
		fmt.Println(err)
		return models.Article{}, err
	}

	// 返り値用のmodels.Article型を宣言
	newArticle := models.Article {
		Title: article.Title,
		Contents: article.Contents,
		UserName: article.UserName,
	}

	// 投稿した記事に割り振られたIDを取得
	id, _ := result.LastInsertId()
	newArticle.ID = int(id)

	return newArticle, nil
}

/*
	変数pageで指定されたページに表示する投稿一覧をデータベースから取得する関数
*/
func SelectedArticleList (db *sql.DB, page int) ([]models.Article, error) {
	// クエリの定義
	const sqlStr = `
		selected article_id, title, contents, username, nice from articles limit ? offset ?;
	`

	// sql.DB型のQueryメソッドを用いてクエリを実行し、得られたデータをrowsに格納
	rows, err := db.Query(sqlStr, articleNumPerPage, (page - 1) * articleNumPerPage)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// models.Articleのスライスを用意
	articleArray := make([]models.Article, 0)
	// rowsにレコードがある間は繰り返し
	for rows.Next() {
		// 一旦、rowsのレコードをarticle変数に格納
		var article models.Article
		rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum)

		// article変数に格納したデータをarticleArrayに追加
		articleArray = append(articleArray, article)
	}

	return articleArray, nil
}

// いいねの数をupdateする関数
// あえて2つのクエリを送る実装にして、トランザクション処理の練習としています
func UpdageNiceNum(db *sql.DB, articleID int) error {
	// クエリの定義
	const sqlGetNice = `
		select nice from articles where article_id = ?;
	`
	const sqlUpdateNice = `
		update articles set nice = ? where article_id = ?
	`

	// トランザクション処理の開始
	tx, err:= db.Begin()
	if err != nil {
		fmt.Println(err)
		return err
	}

	// 1レコードだけ抽出するQueryRowメソッドで、指定した記事IDのいいね数を取得
	row := tx.QueryRow(sqlGetNice, articleID)
	if err := row.Err(); err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	// sql.Row型のメソッドScanでrowからいいね数を抽出
	var nicenum int
	err = row.Scan(&nicenum)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	// Execメソッドでupdateを実行
	_, err = tx.Exec(sqlUpdateNice, nicenum + 1, articleID)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	// Commitして、これまでの一連の処理を確定
	if err := tx.Commit(); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
