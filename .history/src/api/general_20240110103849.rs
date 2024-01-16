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
