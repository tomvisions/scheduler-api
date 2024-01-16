use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Debug)]
pub struct CreateUserSchema {
    pub name: String,
    pub email: String,
    pub description: String,
}
