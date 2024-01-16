//use actix_web::web;
use actix_web::{web::{
    Data,
    Json,
}, post, get, HttpResponse, web};

use crate::{models::group::Group, repository::group::GroupDatabase, AppState,schema::{CreateNoteSchema, FilterOptions, UpdateNoteSchema}, };
//use actix_web::{delete, get, patch, post, web, HttpResponse, Responder};


#[post("/group")]
pub async fn create_group( body: web::Json<CreateNoteSchema>, data: web::Data<AppState>) -> HttpResponse {
    println!("hello");
    println!("{}", data.title);

    let group = db.create_group(data.into_inner());
    match group {
        Ok(group) => HttpResponse::Ok().json(group),
        Err(err) => HttpResponse::InternalServerError().body(err.to_string()),
    }
}

#[get("/group")]
pub async fn get_group(db: web::Data<GroupDatabase>) -> HttpResponse {
    let todos = db.get_group();
    HttpResponse::Ok().json(todos)
}

#[get("/group/{id}")]
pub async fn get_group_by_id(db: web::Data<GroupDatabase>, id: web::Path<String>) -> HttpResponse {
    let todo = db.get_group_by_id(&id);
    match todo {
        Some(todo) => HttpResponse::Ok().json(todo),
        None => HttpResponse::NotFound().body("Todo not found"),
    }
}

pub fn config(cfg: &mut web::ServiceConfig) {
    cfg.service(
        web::scope("/api")
            .service(create_group)
          //  .service(get_group)
           // .service(get_group_by_id)
    );
}
