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
        web::scope("/api")
            .service(create_group)
          //  .service(get_group)
           // .service(get_group_by_id)
    );
}
