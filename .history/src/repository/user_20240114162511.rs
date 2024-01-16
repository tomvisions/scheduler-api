use std::fmt::Error;
use chrono::prelude::*;
use crate::models::user::UserModel;
use std::sync::{Arc, Mutex};

pub struct UserDatabase {
    pub users: Arc<Mutex<Vec<UserModel>>>,
}


impl UserDatabase {

pub fn create_user(&self, user: UserModel) -> Result<UserModel, Error> {
    let mut users = self.users.lock().unwrap();
    let id = uuid::Uuid::new_v4().to_string();
    let created_at = Utc::now();
    let updated_at = Utc::now();
    let user = User {
        id: Some(id),
        created_at: Some(created_at),
        updated_at: Some(updated_at),
        ..user
    };
    users.push(user.clone());
    Ok(user)
}

pub fn get_user(&self, user: UserModel) -> Result<UserModel, Error> {
    let mut users = self.users.lock().unwrap();
    let id = uuid::Uuid::new_v4().to_string();
    let created_at = Utc::now();
    let updated_at = Utc::now();
    let user = User {
        id: Some(id),
        created_at: Some(created_at),
        updated_at: Some(updated_at),
        ..user
    };
    users.push(user.clone());
    Ok(user)
}

pub fn get_user_by_id(&self, id: &str) -> Option<User> {
    let users = self.users.lock().unwrap();
    users.iter().find(|user| user.id == Some(id.to_string())).cloned()
}





}
