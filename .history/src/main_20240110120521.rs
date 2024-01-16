use actix_web::{web, App, HttpResponse, HttpServer, Result, http::header};
use serde::{Serialize};

use sqlx::mysql::{MySqlPool, MySqlPoolOptions};
use actix_cors::Cors;
use dotenv::dotenv;

mod api;
mod models;
mod repository;

#[derive(Serialize)]
pub struct Response {
    pub status: String,
    pub message: String,
}

pub struct AppState {
    db: MySqlPool,
}

async fn not_found() -> Result<HttpResponse> {
    let response = Response {
        status: "404".to_string(),
        message: "Resource not found".to_string(),
    };
    Ok(HttpResponse::NotFound().json(response))
}


#[actix_web::main]
async fn main() -> std::io::Result<()> {
    dotenv().ok();
    std::env::set_var("RUST_LOG", "debug");
    env_logger::init();

    let database_url = std::env::var("DATABASE_URL").expect("DATABASE_URL must be set");
    let pool = match MySqlPoolOptions::new()
        .max_connections(10)
        .connect(&database_url)
        .await
    {
        Ok(pool) => {
            println!("✅Connection to the database is successful!");
            pool
        }
        Err(err) => {
            println!("🔥 Failed to connect to the database: {:?}", err);
            std::process::exit(1);
        }
    };

    HttpServer::new(move || {
        let cors = Cors::default()
       // .allowed_origin("http://localhost:3000")
        .allowed_methods(vec!["GET", "POST", "PATCH", "DELETE"])
        .allowed_headers(vec![
            header::CONTENT_TYPE,
            header::AUTHORIZATION,
            header::ACCEPT,
        ])
        .supports_credentials();
         App::new()
            .app_data(web::Data::new(AppState {db: pool.clone()} ))
            .configure(api::group::config)
            .configure(api::user::config)
            .configure(api::general::config)
            .wrap(cors)
            .wrap(actix_web::middleware::Logger::default())
            .default_service(web::route().to(not_found))
        })
        .bind(("127.0.0.1", 8080))?
        .run()
        .await
}       