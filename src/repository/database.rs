use std::fmt::Error;
use chrono::prelude::*;
use std::sync::{Arc, Mutex};

use crate::models::group::Group;
use crate::models::user::User;


pub struct Database {
    pub group: Arc<Mutex<Vec<Group>>>,
    pub user: Arc<Mutex<Vec<User>>>,
}

impl Database {
    pub fn new() -> Self {
        let group = Arc::new(Mutex::new(vec![]));
        let user = Arc::new(Mutex::new(vec![]));
        Database { group, user }
        //Database { user }
    }
}