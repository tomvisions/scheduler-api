use std::{fmt::Error, clone};
use actix_web::Responder;
use chrono::prelude::*;
use sqlx::{any, Execute};
use crate::{models::usher_group::{UsherGroupModel, UsherGroupResponse}};
use std::sync::{Arc, Mutex};
use actix_web::{web::{
    Data,
    Json,
}, post, get, web, HttpResponse}
use serde_json::json;
use std::fmt::{self, write};


use crate::AppState;

pub struct GroupDatabase {
//    pub groups: Arc<Mutex<Vec<Group>>>,
      pub create_group: Arc<Mutex<Vec<UsherGroupModel>>>,

}


impl UsherGroupDatabase {
    pub async fn create_group(group: web::Json<UsherGroupModel>, data:web::Data<AppState>) -> Result<UsherGroupResponse, Error> {  //impl Responder {
    
    
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
            .bind(id.clone())
            .bind(group.title.to_string())
            .bind(group.description.to_string())
            .bind(group.created_at)
            .bind(group.updated_at)
            .execute(&data.db)
            .await
            .map_err(|err: sqlx::Error| err.to_string());

        if let Err(err) = query_result {
          //  if err.contains("Duplicate entry") {
                Err(err);
//                return HttpResponse::BadRequest().json(
  //              serde_json::json!({"status": "fail","message": "Group with that title already exists"}),
           // }
        
            
            //return HttpResponse::InternalServerError()
              //  .json(serde_json::json!({"status": "error","message": format!("{:?}", err)}));
        }
        let query_result = sqlx::query_as!(UsherGroupModel, r#"SELECT * FROM usher_group WHERE id = ?"#, id)
//        let query_result = sqlx::query(r#"SELECT * FROM group WHERE id = ?"#).bind(id)
        .fetch_one(&data.db)
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

            let grooup_jeon = serde_json::from_str(group_response);
            return Ok(serde_json::from_str(group_response));

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

}


    fn filter_db_record(group: &Group) -> GroupResponse {
        GroupResponse {
            id: group.id.to_owned(),
            title: group.title.to_owned(),
            description: group.description.to_owned(),
            createdAt: group.created_at.unwrap(),
            updatedAt: group.updated_at.unwrap(),
        }
    }