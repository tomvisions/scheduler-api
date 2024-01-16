use std::fmt::Error;
use chrono::prelude::*;
use crate::models::group::Group;
use std::sync::{Arc, Mutex};

pub struct GroupDatabase {
    pub groups: Arc<Mutex<Vec<Group>>>,
}


impl GroupDatabase {
    pub fn create_group(group: Group) -> Result<Group, Error> {
   //     let mut groups = self.groups.lock().unwrap();
        let id = uuid::Uuid::new_v4().to_string();
        let title = group.title;
        let description = group.description;

        let created_at = Utc::now();
        let updated_at = Utc::now();
        let group = Group {
            id: Some(id),
            title: title,
            description: description,
            created_at: Some(created_at),
            updated_at: Some(updated_at),
            ..group
        };
        groups.push(group.clone());
        Ok(group)
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
