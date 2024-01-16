use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Debug)]
pub struct CreateUsherGroupSchema {
    pub title: String,
    pub description: String,
}