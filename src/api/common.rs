
use crate::api::user::config as user_config;
use crate::api::usher_group::config as usher_group_config;
use actix_web::{web, App, HttpResponse, HttpServer};
use actix_web::{web::{
    Data,
    Json,
}, post, get};


pub fn config(conf: &mut web::ServiceConfig) {

   // conf.service(
      //  web::resource("/app")
     //       .route(web::get().to(|| HttpResponse::Ok().body("app")))
    //        .route(web::head().to(|| HttpResponse::MethodNotAllowed())),
   // )
   // .service(web::scope("/api").configure(scoped_config))
    //.route("/", web::get().to(|| HttpResponse::Ok().body("/")));

    


    let scope = web::scope("/api")
    .service(web::scope("/user").configure(user_config))
    .service(web::scope("/usher-group").configure(usher_group_config));

    conf.service(scope);
}
