
use actix_web::{get, web, App, HttpResponse, HttpServer, Responder, Result};
use serde::{Serialize};

#[derive(Serialize)]

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

pub fn config(cfg: &mut web::ServiceConfig) {
    println!("dfdfd");
    cfg.service(
        web::scope("/")
            .service(healthcheck)

          //  .service(get_group)
           // .service(get_group_by_id)
    );
}
