package model

import (
	"fmt"
	"log"
	"scheduler-api/db"
	e "scheduler-api/entity"

	sg "github.com/Masterminds/squirrel"
)

func GetUsherGroups() ([]e.Gallery, error) {

	var galleryList []e.Gallery
	gallery := e.Gallery{}
	gallerySQL, args, err := sg.Select("gallery.id, gallery.image_name, gallery.gallery_name, gallery.image_name, gallery.slug, category.category_name, tag.tag_name").
		From("gallery").
		LeftJoin("category ON gallery.category = category.id").
		LeftJoin("tag ON tag.id = gallery.tag").
		Where("gallery.main_featured = 1").
		ToSql()

	//	fmt.Println(gallerySQL)
	fmt.Println(args)

	//	fmt.Println("here")
	rows, err := db.DB.Queryx(gallerySQL)

	if err != nil {
		panic(err)
	}

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

func GetUsherGroupById(category string) ([]e.Gallery, error) {
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
