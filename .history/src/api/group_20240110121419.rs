use actix_web::web;
use actix_web::{web::{
    Data,
    Json,
}, post, get, HttpResponse};
use crate::{models::group::Group, repository::group::GroupDatabase};


#[post("/group")]
pub async fn create_group(db: Data<GroupDatabase>, data: web::Data<AppState>) -> HttpResponse {
    println!("hello");
    println!("{}", new_group.title);

    let group = db.create_group(new_group.into_inner());
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
