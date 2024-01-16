use std::{fmt::Error, clone};
use actix_web::Responder;
use chrono::prelude::*;
use sqlx::{any, Execute};
use crate::models::usher_group::{UsherGroupModelResponse, CreateUsherGroupSchema, ResponseCall, UsherGroupModel};
use std::sync::{Arc, Mutex};
use actix_web::{web::{
    Data,
    Json,
}, post, get, web, HttpResponse}
use serde_json::json;
use std::fmt::{self, write};

pub struct UsherGroupDatabase {
 //  pub groups: Arc<Mutex<Vec<Group>>>,
   //   pub create_group: Arc<Mutex<Vec<UsherGroupModel>>>,

}


impl UsherGroupDatabase {
    pub async fn create_group(body: web::Json<CreateUsherGroupSchema>, data:web::Data<AppState>) -> Result<UsherGroupModelResponse, Error> {  //impl Responder {
    
    
        //let mut groups = self.groups.lock().unwrap();
        let id = uuid::Uuid::new_v4().to_string();
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
        let data2 = AppState.db;

        let query_result =
        sqlx::query(r#"INSERT INTO group (id,title,description,created_at, updated_at) VALUES (?, ?, ?, ?, ?)"#)
            .bind(id.to_string())
            .bind(body.title.to_string())
            .bind(body.description.to_string())
            .bind(created_at)
            .bind(updated_at)
            .execute(&data.db)
            .await
            .map_err(|err: sqlx::Error| err.to_string());

     /*    if let Err(err) = query_result {
          //  if err.contains("Duplicate entry") {
                Err(err);
//                return HttpResponse::BadRequest().json(
  //              serde_json::json!({"status": "fail","message": "Group with that title already exists"}),
           // }
                
            
            //return HttpResponse::InternalServerError()
              //  .json( *serde_json::json!({"status": "error","message": format!("{:?}", err)}));
       // } */
        let query_result = sqlx::query_as!(UsherGroupModel, r#"SELECT * FROM usher_group WHERE id = ?"#, id.to_string())
//        let query_result = sqlx::query(r#"SELECT * FROM group WHERE id = ?"#).bind(id)
        .fetch_one(&mut data.db)
        .await;

        let mut output = String::new();
        match query_result {
        Ok(group) => {
        //    Ok(group);

            let group_response = serde_json::json!({"status": "success","data": serde_json::json!({
                "group": filter_db_record(&group)
            })});

       //  let group_response = r#"{"status": "success","data": serde_json::json!({
         //       "note": "test"
          //    });
         //     let test = r#"{"hello":"sdfaf"}"#;

            let group_json: ResponseCall  = serde_json::from_value(group_response);
            return Ok(group_json);

           // return HttpResponse::Ok().json(group_response);
        }
        Err(e) => {
            write(&mut output, format_args!("Hello {}!", e)) 

        }
       //   //  Err(e); 
 //         let group_error = serde_json::json!({"status": "error","data": serde_json::json!({
   //         "note": "test"
    //   })});
      //      Error(group_error);
    //    Err(group_error);
    }
        
    //    Err(serde_json::json!({"status": "error","message": format!("{:?}", e))
     //       return HttpResponse::InternalServerError()
       //         .json(serde_json::json!({"status": "error","message": format!("{:?}", e)}));
      //  }
}

#[post("/notes/")]
async fn create_note_handler(
    body: web::Json<CreateNoteSchema>,
    data: web::Data<AppState>,
) -> impl Responder {
    let user_id = uuid::Uuid::new_v4().to_string();
    let query_result =
        sqlx::query(r#"INSERT INTO notes (id,title,content,category) VALUES (?, ?, ?, ?)"#)
            .bind(user_id.clone())
            .bind(body.title.to_string())
            .bind(body.content.to_string())
            .bind(body.category.to_owned().unwrap_or_default())
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

    let query_result = sqlx::query_as!(NoteModel, r#"SELECT * FROM notes WHERE id = ?"#, user_id)
        .fetch_one(&data.db)
        .await;

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