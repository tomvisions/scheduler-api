use std::error::Error;

use actix_web::web;
use actix_web::{web::{
    Data,
    Json,
}, post, get,  HttpResponse, Responder};
use sqlx::types::chrono::Utc;

use crate::{
    model::user::{UserModel, UserModelResponse},
    schema::user::CreateUserSchema,
    schema::common::FilterOptions,
    AppState,
};

#[get("")]
pub async fn get_user(
    opts: web::Query<FilterOptions>,
    data: web::Data<AppState>,
) -> impl Responder {
    let limit = opts.limit.unwrap_or(10);
    let offset = (opts.page.unwrap_or(1) - 1) * limit;

    let notes: Vec<UserModel> = sqlx::query_as::<_, UserModel>(
        r#"SELECT * FROM user ORDER by id LIMIT ? OFFSET ?"#).bind(limit as i32).bind(offset as i32)
    .fetch_all(&data.db)
    .await
    .unwrap();

    let note_responses = notes
        .into_iter()
        .map(|note| filter_db_record(&note))
        .collect::<Vec<UserModelResponse>>();

    let json_response = serde_json::json!({
        "status": "success",
        "results": note_responses.len(),
        "users": note_responses
    });
    HttpResponse::Ok().json(json_response)
}


#[post("")]
pub async fn create_user(
    body: web::Json<CreateUserSchema>,
    data: web::Data<AppState>,
) -> impl Responder {
    print!("arrive");
    //let mut groups = self.groups.lock().unwrap();
    let id = uuid::Uuid::new_v4().to_string();
    let created_at = Utc::now();
    let updated_at = Utc::now();
    
    print!("hello");
    let query_result = sqlx::query(
        r#"INSERT INTO user (id, name, email,description,created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"#,
    )
    .bind(id.to_string())
    .bind(body.name.to_string())
    .bind(body.email.to_string())
    .bind(body.description.to_string())
    .bind(created_at)
    .bind(updated_at)
    .execute(&data.db)
    .await
    .map_err(|err: sqlx::Error| err.to_string());

    if let Err(err) = query_result {
        if err.contains("Duplicate entry") {
            return HttpResponse::BadRequest().json(
            serde_json::json!({"status": "fail","message": "User with that email already exists"}),
        );
        }

        return HttpResponse::InternalServerError()
            .json(serde_json::json!({"status": "error","message": format!("{:?}", err)}));
    }


    let query_result = sqlx::query_as::<_, UserModel>(
       
        r#"SELECT * FROM user WHERE id = ?"#).bind(id)
    //        let query_result = sqlx::query(r#"SELECT * FROM group WHERE id = ?"#).bind(id)
    .fetch_one(&data.db)
    .await;

    let mut output = String::new();
    match query_result {
        Ok(note) => {
            let note_response = serde_json::json!({"status": "success","data": serde_json::json!({
                "note": filter_db_record(&note)
            })});

            return HttpResponse::Ok().json(note_response);
        }
        Err(e) => {
            return HttpResponse::InternalServerError()
                .json(serde_json::json!({"status": "error","message": format!("{:?}", e)}));
        }
    }
}

pub fn config(conf: &mut web::ServiceConfig) {
  //  print!("userddddconfig");
  //  cfg.service(
          //web::resource("/user").route(route)
    //    web::scope("/user")
      //      .service(create_user)
  //  );

    let scope = web::scope("")
    .service(create_user)
    .service(get_user);
    conf.service(scope);
}

fn filter_db_record(user: &UserModel) -> UserModelResponse {
    UserModelResponse {
        id: user.id.to_owned(),
        email: user.email.to_owned(),
        name: user.name.to_owned(),
        description: user.description.to_owned(),
        createdAt: user.created_at.unwrap(),
        updatedAt: user.updated_at.unwrap(),
    }
}