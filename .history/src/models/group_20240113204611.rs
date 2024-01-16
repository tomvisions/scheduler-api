use chrono::prelude::{DateTime, Utc};
use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Debug, Clone,  sqlx::FromRow)]
#[allow(non_snake_case)]
pub struct Group {
    pub id: String,
    pub title: String,
    pub description: String,
    pub created_at: Option<chrono::DateTime<chrono::Utc>>,
    pub updated_at: Option<chrono::DateTime<chrono::Utc>>,
}

pub struct Response {
    pub status: String,
    pub data: String,
}

#[allow(non_snake_case)]
pub struct GroupResponse {
    pub id: String,
    pub title: String,
    pub description: String,
    pub createdAt: chrono::DateTime<chrono::Utc>,
    pub updatedAt: chrono::DateTime<chrono::Utc>,
}

