use actix_web::{get, web, App, HttpResponse, HttpServer, Responder, Result, header};
use serde::{Serialize};
//use crate::config::AppConfig;
use lazy_static::lazy_static;

use sqlx::mysql::{MySqlPool, MySqlPoolOptions};


mod api;
mod models;
mod repository;

#[derive(Serialize)]
pub struct Response {
    pub status: String,
    pub message: String,
}

lazy_static! {
   // pub static ref SETTINGS: AppConfig = {
    //    let cli_options = cli::Options::new();
      //  AppConfig::new(cli_options.config_dir.as_ref()).unwrap()
//    };
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

pub struct AppState {
    db: MySqlPool,
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    std::env::set_var("RUST_LOG", "debug");
    env_logger::init();

    //let app_state = app_state::AppState::new().await;
//let app_data = actix_web::web::Data::new(new_data);
   // let app_data = web::Data::new(the_db);
    HttpServer::new(move || {
        let cors = Cors::default()
        .allowed_origin("http://localhost:3000")
        .allowed_methods(vec!["GET", "POST", "PATCH", "DELETE"])
        .allowed_headers(vec![
            header::CONTENT_TYPE,
            header::AUTHORIZATION,
            header::ACCEPT,
        ])
         App::new()
            .app_data(web::Data::new(AppState {db: pool.clone()}))
            .configure(api::group::config)
            .service(healthcheck)
            .default_service(web::route().to(not_found)).
            wrap(actix_web::middleware::Logger::default())
})
        .bind(("127.0.0.1", 8080))?
        .run()
        .await
}       