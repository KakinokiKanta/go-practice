package repositories

import (
	"database/sql"
	"fmt"

	"github.com/KakinokiKanta/go-intermediate/models"
)

/*
	新規投稿をデータベースにinsertする関数
*/
func InsertComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	// クエリの定義
	const sqlStr = `
		insert into comments (article_id, message, created_at) values (?, ?, now());
	`

	// Execメソッドでクエリの実行
	result, err := db.Exec(sqlStr, comment.ArticleID, comment.Message)
	if err != nil {
		fmt.Println(err)
		return models.Comment{}, err
	}

	// 新規投稿した記事用の変数newCommentを用意
	var newComment models.Comment
	newComment.ArticleID = comment.ArticleID
	newComment.Message = comment.Message

	// sql.Result型のメソッドLastInsertIdで、新規投稿した記事のIDを取得
	id, _ := result.LastInsertId()
	newComment.CommentID = int(id)

	return newComment, nil
}

/*
	指定IDの記事についたコメント一覧を取得する関数
*/
func SelectedCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {
	// クエリを定義
	const sqlStr = `
		select * from comments where article_id = ?;
	`

	// sql.DB型のメソッドQueryを用いてクエリを実行し、rowsに結果を格納
	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	// models.Comment型のスライスを用意
	commentArray := make([]models.Comment, 0)
	// rowsの中身がなくなるまで繰り返し
	for rows.Next() {
		// models.Comment型の変数commentを用意
		var comment models.Comment
		// createdTimeがnullだったときのための変数を用意
		var createdTime sql.NullTime
		// rowsの要素をcommentに読み出し
		rows.Scan(&comment.CommentID, &comment.ArticleID, &comment.Message, &createdTime)

		// sql.NullTime型のcreatedTimeのフィールドであるValidがtrueなら、comment.CreatedAtに書き込む
		// Validフィールドは値がnullでなければ、trueとなる
		if createdTime.Valid {
			comment.CreatedAt = createdTime.Time
		}

		// commentArrayにcommentを追加
		commentArray = append(commentArray, comment)
	}

	return commentArray, nil
}
