use std::fmt::Error;
use actix_web::Responder;
use chrono::prelude::*;
use sqlx::{any, Execute};
use crate::{models::group::Group, api::group::create_group};
use std::sync::{Arc, Mutex};
use actix_web::{web::{
    Data,
    Json,
}, post, get, HttpResponse, web}
use serde_json::json;

use crate::AppState;

pub struct GroupDatabase {
//    pub groups: Arc<Mutex<Vec<Group>>>,
      pub create_group: Arc<Mutex<Vec<Group>>>,

}


impl GroupDatabase {
    pub async fn create_group(group: web::Json<Group>, data:web::Data<AppState>) -> impl Responder {
    
    
        //let mut groups = self.groups.lock().unwrap();
        let id = uuid::Uuid::new_v4().to_string();
        let title = group.title;
        let description = group.description;

        let created_at = Utc::now();
        let updated_at = Utc::now();
       /*  let group = Group {
            id: Some(id),
            title: title,
            description: description,
            created_at: Some(created_at),
            updated_at: Some(updated_at),
            ..group
        }; */
     //   groups.push(group.clone());
       // Ok(group)

        let query_result =
        sqlx::query(r#"INSERT INTO group (id,title,description,created_at, updated_at) VALUES (?, ?, ?, ?, ?)"#)
            .bind(id.to_string())
            .bind(group.title.to_string())
            .bind(group.description.to_string())
            .bind(group.created_at)
            .bind(group.updated_at)
            .execute(&data.db)
            .await
            .map_err(|err: sqlx::Error| err.to_string());

        if let Err(err) = query_result {
            if err.contains("Duplicate entry") {
                return HttpResponse::BadRequest().json(
                serde_json::json!({"status": "fail","message": "Note with that title already exists"}),
            );
            }
        
            return HttpResponse::InternalServerError()
        .json(serde_json::json!({"status": "error","message": format!("{:?}", err)}));
        }

        let query_result = sqlx::query_as!(Group, r#"SELECT * FROM group WHERE id = ?"#, id)
        .fetch_one(&data.db)
        .await;

         match query_result {
        Ok(group) => {
            let group_response = serde_json::json!({"status": "success","data": serde_json::json!({
                "note": "test"
            })});

            return HttpResponse::Ok().json(group_response);
        }
        Err(e) => {
            return HttpResponse::InternalServerError()
                .json(serde_json::json!({"status": "error","message": format!("{:?}", e)}));
        }
    }



}


/*     pub fn get_group(&self) -> Vec<Group> {
        let groups = self.groups.lock().unwrap();
        groups.clone()
    }

    pub fn get_group_by_id(&self, id: &str) -> Option<Group> {
        let groups = self.groups.lock().unwrap();
        groups.iter().find(|todo| todo.id == Some(id.to_string())).cloned()
    }
    */
}