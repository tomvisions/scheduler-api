use chrono::prelude::{DateTime, Utc};
use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Debug, Clone,  sqlx::FromRow)]
#[allow(non_snake_case)]
pub struct Group {
    pub id: String,
    pub title: String,
    pub description: String,
    pub created_at: Option<DateTime<Utc>>,
    pub updated_at: Option<DateTime<Utc>>,
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
    pub createdAt: Option<DateTime<Utc>>,
    pub updatedAt: Option<DateTime<Utc>>,
}

