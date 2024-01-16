use std::sync::{Arc, Mutex};

use crate::models::group::Group;
use crate::models::user::User;
use crate::repository::group::GroupDatabase;

pub struct Database {
    pub group: Arc<Mutex<Vec<Group>>>,
  //  pub user: Arc<Mutex<Vec<User>>>,
}

impl Database {
    pub fn new() -> Self {
        let group = Arc::new(Mutex::new(vec![]));
    //    let user: Arc<Mutex<Vec<_>>> = Arc::new(Mutex::new(vec![]));
        Database { group }
        //Database { user }
    }
}