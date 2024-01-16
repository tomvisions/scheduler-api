use std::fmt::Error;
use actix_web::Responder;
use chrono::prelude::*;
use sqlx::{any, Execute};
use crate::{models::group::Group, api::group::create_group};
use std::sync::{Arc, Mutex};

pub struct GroupDatabase {
//    pub groups: Arc<Mutex<Vec<Group>>>,
      pub create_group: Arc<Mutex<Vec<Group>>>,

}


impl GroupDatabase {
    pub fn create_group(group: Group) -> Result<Group, Error> {
    
    
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
            .execute(&data.db);


    }

    pub fn get_group(&self) -> Vec<Group> {
        let groups = self.groups.lock().unwrap();
        groups.clone()
    }

    pub fn get_group_by_id(&self, id: &str) -> Option<Group> {
        let groups = self.groups.lock().unwrap();
        groups.iter().find(|todo| todo.id == Some(id.to_string())).cloned()
    }


}
