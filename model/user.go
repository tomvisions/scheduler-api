package model

import (
	"fmt"
	"log"
	"scheduler-api/db"
	e "scheduler-api/entity"
	"strings"

	sg "github.com/Masterminds/squirrel"
)

func AddUser(user *e.User) error {

	fmt.Printf("name valueL: %s\n", user.Name)
	//now := time.Now()
	id := strings.ToLower(user.Name)

	if strings.Contains(id, ",") {
		temp := strings.Split(id, ",")
		fmt.Printf("temp valueL: %s\n", temp)
		id = fmt.Sprintf("%s %s", temp[1], temp[0])
		fmt.Printf("id valueL: %s\n", id)

	}
	id = strings.Trim(id, " ")
	id = strings.Replace(id, " ", "-", -1)
	fmt.Printf("id2 valueL: %s\n", id)

	const query = `INSERT INTO user (id, name,email, description, phone) VALUES (?, ?, ?, ?, ?)`
	tx, err := db.DB.Begin()

	if err != nil {
		return err
	}

	_, err = tx.Exec(query, id, user.Name, user.Email, user.Description, user.Phone)
	fmt.Printf("err valueL: %s\n", err)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func GetUsers(pageIndex uint64, pageSize uint64, field string, order string) ([]e.User, error) {

	var userList []e.User
	user := e.User{}

	offset := ((pageIndex - 1) * pageSize)

	orderBy := fmt.Sprintf("%s %s", field, order)

	userListSQL, args, err := sg.Select("user.id, user.name, user.email, user.phone").
		From("user").
		OrderBy(orderBy).
		Limit(pageSize).
		Offset(offset).
		ToSql()

	fmt.Println(userListSQL)
	fmt.Println(args)

	//	fmt.Println("here")
	rows, err := db.DB.Queryx(userListSQL)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.StructScan(&user)

		if err != nil {
			log.Fatalln(err)
		}

		userList = append(userList, user)
	}
	///	test := fmt.Sprintf("{data: %s}", userList)

	err = rows.Err()

	return userList, err
}

func GetUserById(id string) (e.User, error) {
	var user e.User

	userSQL, args, err := sg.Select("user.id, user.name, user.description, user.email, user.email, (SELECT CAST(CONCAT('[',GROUP_CONCAT(JSON_OBJECT('Id', `usher_group`.`id`)),']') as JSON) FROM usher_group LEFT JOIN user_usher_group ON user_usher_group.usher_group = usher_group.id where user_usher_group.user = user.id) as usher_group").
		From("user").
		Where(sg.Eq{"user.id": id}).ToSql()

	fmt.Println(userSQL)
	fmt.Println(args)

	rows, err := db.DB.Queryx(userSQL, id)

	if err != nil {
		panic(err)
	}
	fmt.Println(rows)

	defer rows.Close()

	for rows.Next() {
		err := rows.StructScan(&user)

		if err != nil {
			log.Fatalln(err)
		}

	}

	err = rows.Err()
	return user, err
}

func GetUsersByUsherGroupId(category string) ([]e.Gallery, error) {
	var galleryList []e.Gallery
	fmt.Println(category)
	gallery := e.Gallery{}
	gallerySQL, args, err := sg.Select("gallery.id, gallery.image_name, gallery.gallery_name, gallery.image_name, gallery.slug, category.category_name, tag.tag_name").
		From("gallery").
		LeftJoin("category ON gallery.category = category.id").
		LeftJoin("tag ON tag.id = gallery.tag").
		Where(sg.Eq{"gallery.category": category, "gallery.featured": 1}).ToSql()

	fmt.Println(gallerySQL)
	fmt.Println(args)

	rows, err := db.DB.Queryx(gallerySQL, category, 1)

	if err != nil {
		panic(err)
	}
	fmt.Println(rows)

	defer rows.Close()

	for rows.Next() {
		err := rows.StructScan(&gallery)

		if err != nil {
			log.Fatalln(err)
		}

		galleryList = append(galleryList, gallery)
	}

	err = rows.Err()
	return galleryList, err
}
