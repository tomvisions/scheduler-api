
use actix_web::{get, web, App, HttpResponse, HttpServer, Responder, Result, header};
use actix_web::web;
use actix_web::{web::{
    Data,
    Json,
}, post, get, HttpResponse};
use crate::{models::group::Group, repository::group::GroupDatabase};

pub struct Response {
    pub status: String,
    pub message: String,
}


#[get("/health")]
async fn healthcheck() -> impl Responder {
    let response = Response {
        status: "200".to_string(),
        message: "Everything is working fine".to_string(),
    };
    HttpResponse::Ok().json(response)
}


async fn not_found() -> Result<HttpResponse> {
    let response = Response {
        status: "404".to_string(),
        message: "Resource not found".to_string(),
    };
    Ok(HttpResponse::NotFound().json(response))
}

pub fn config(cfg: &mut web::ServiceConfig) {
    println!("dfdfd");
    cfg.service(
        web::scope("/")
            .service(healthcheck)

          //  .service(get_group)
           // .service(get_group_by_id)
    );
}
