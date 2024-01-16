use actix_web::web;
use actix_web::{web::{
    Data,
    Json,
}, post, get,  HttpResponse};
use crate::repository::user::UserDatabase;


#[post("/user")]
pub async fn create_user(db: Data<UserDatabase>, new_todo: Json< User>) -> HttpResponse {
    let todo = db.create_user(new_todo.into_inner());
    match todo {
        Ok(todo) => HttpResponse::Ok().json(todo),
        Err(err) => HttpResponse::InternalServerError().body(err.to_string()),
    }
}
                      
#[get("/user")]
pub async fn get_user(db: Data<UserDatabase>, new_todo: Json<User>) -> HttpResponse {
    let todo = db.get_user(new_todo.into_inner());
    match todo {
        Ok(todo) => HttpResponse::Ok().json(todo),
        Err(err) => HttpResponse::InternalServerError().body(err.to_string()),
    }
}

#[get("/user/{id}")]
pub async fn get_user_by_id(db: web::Data<UserDatabase>, id: web::Path<String>) -> HttpResponse {
    let todo = db.get_user_by_id(&id);
    match todo {
        Some(todo) => HttpResponse::Ok().json(todo),  
        None => HttpResponse::NotFound().body("User not found"),
    }
}

pub fn config(cfg: &mut web::ServiceConfig) {
    cfg.service(
        web::scope("/api")
            .service(create_user)
            .service(get_user)
            .service(get_user_by_id)
    );
}
