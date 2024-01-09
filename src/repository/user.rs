use std::fmt::Error;
use chrono::prelude::*;

impl Database for User {

pub fn create_user(&self, todo: Todo) -> Result<User, Error> {
    let mut todos = self.todos.lock().unwrap();
    let id = uuid::Uuid::new_v4().to_string();
    let created_at = Utc::now();
    let updated_at = Utc::now();
    let user = User {
        id: Some(id),
        created_at: Some(created_at),
        updated_at: Some(updated_at),
        ..user
    };
    user.push(user.clone());
    Ok(todo)
}
}
