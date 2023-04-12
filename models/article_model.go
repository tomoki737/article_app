package models

import (
	"app/database"
)

type Article struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Articles struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Comment struct {
	Id        uint64
	ArticleId uint64
	UserId    uint64
	Text      string
}

func (a *Article) CreateArticle() error {
	db := database.GetDB()
	stmt, err := db.Prepare("INSERT INTO articles(title, body) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(a.Title, a.Body)
	if err != nil {
		return err
	}
	return nil
}

func (a *Article) EditArticle() error {
	db := database.GetDB()
	stmt, err := db.Prepare("UPDATE articles SET title = ?, body = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(a.Title, a.Body, a.Id)
	if err != nil {
		return err
	}
	return nil
}

func (a *Article) DeleteArticle() error {
	db := database.GetDB()
	stmt, err := db.Prepare("DELETE FROM articles WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(a.Id)
	if err != nil {
		return err
	}
	return nil
}

func GetAllArticles() ([]Article, error) {
	db := database.GetDB()
	var articles []Article

	rows, err := db.Query("SELECT id, title, body FROM articles")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		article := &Article{}
		err := rows.Scan(&article.Id, &article.Title, &article.Body)
		if err != nil {
			return nil, err
		}

		articles = append(articles, Article{
			Id:    article.Id,
			Title: article.Title,
			Body:  article.Body,
		})
	}
	return articles, nil
}

func (a *Article) GetSingleArticle() error {
	db := database.GetDB()
	row := db.QueryRow("SELECT title, body FROM articles where id = ?", a.Id)

	err := row.Scan(&a.Title, &a.Body)

	if err != nil {
		return err
	}
	return nil
}

func SearchArticles(title, body string) ([]Article, error) {
	db := database.GetDB()
	var articles []Article
	query := "SELECT id, title, body FROM articles WHERE 1 = 1"
	if title != "" {
		query += " AND title Like '%" + title + "%'"
	}
	if body != "" {
		query += " AND body Like '%" + body + "%'"
	}

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		article := &Article{}
		err := rows.Scan(&article.Id, &article.Title, &article.Body)
		if err != nil {
			return nil, err
		}

		articles = append(articles, Article{
			Id:    article.Id,
			Title: article.Title,
			Body:  article.Body,
		})
	}
	return articles, nil
}

func (c *Comment) CommentSave() error {
	db := database.GetDB()
	stmt, err := db.Prepare("INSERT INTO comments(article_id, user_id, text) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(c.ArticleId, c.UserId, c.Text)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	c.Id = uint64(id)

	return nil
}
