use std::error::Error;

use crate::{
    model::usher_group::{UsherGroupModel, UsherGroupModelResponse},
    schema::common::FilterOptions,
    schema::usher_group::CreateUsherGroupSchema,
    AppState,
};
use sqlx::types::chrono::Utc;
use actix_web::{delete, get, patch, post, web, HttpResponse, Responder};
use serde_json::json;

#[get("")]
pub async fn get_usher_group(
    opts: web::Query<FilterOptions>,
    data: web::Data<AppState>,
) -> impl Responder {
    let limit = opts.limit.unwrap_or(10);
    let offset = (opts.page.unwrap_or(1) - 1) * limit;

    let notes: Vec<UsherGroupModel> = sqlx::query_as::<_, UsherGroupModel>(
        r#"SELECT * FROM usher_group ORDER by id LIMIT ? OFFSET ?"#).bind(limit as i32).bind(offset as i32)
    .fetch_all(&data.db)
    .await
    .unwrap();

    let note_responses = notes
        .into_iter()
        .map(|note| filter_db_record(&note))
        .collect::<Vec<UsherGroupModelResponse>>();

    let json_response = serde_json::json!({
        "status": "success",
        "results": note_responses.len(),
        "usher_groups": note_responses
    });
    HttpResponse::Ok().json(json_response)
}


#[post("")]
pub async fn create_group(
    body: web::Json<CreateUsherGroupSchema>,
    data: web::Data<AppState>,
) -> impl Responder {
    //let mut groups = self.groups.lock().unwrap();
    let id = uuid::Uuid::new_v4().to_string();
    let created_at = Utc::now();
    let updated_at = Utc::now();
    
    print!("hello");
    let query_result = sqlx::query(
        r#"INSERT INTO usher_group (id,title,description,created_at, updated_at) VALUES (?, ?, ?, ?, ?)"#,
    )
    .bind(id.to_string())
    .bind(body.title.to_string())
    .bind(body.description.to_string())
    .bind(created_at)
    .bind(updated_at)
    .execute(&data.db)
    .await
    .map_err(|err: sqlx::Error| err.to_string());

    if let Err(err) = query_result {
        if err.contains("Duplicate entry") {
            return HttpResponse::BadRequest().json(
            serde_json::json!({"status": "fail","message": "Group with that title already exists"}),
        );
        }

        return HttpResponse::InternalServerError()
            .json(serde_json::json!({"status": "error","message": format!("{:?}", err)}));
    }


    let query_result = sqlx::query_as::<_, UsherGroupModel>(
       
        r#"SELECT * FROM usher_group WHERE id = ?"#).bind(id)
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

fn filter_db_record(group: &UsherGroupModel) -> UsherGroupModelResponse {
    UsherGroupModelResponse {
        id: group.id.to_owned(),
        title: group.title.to_owned(),
        description: group.description.to_owned(),
        createdAt: group.created_at.unwrap(),
        updatedAt: group.updated_at.unwrap(),
    }
}

pub fn config(conf: &mut web::ServiceConfig) {
    print!("start ocnfig");
    let scope = web::scope("")
    .service(create_group)
    .service(get_usher_group);
    conf.service(scope);
}
