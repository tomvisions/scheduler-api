use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Debug, Clone,  sqlx::FromRow)]
#[allow(non_snake_case)]
pub struct UsherGroupModel {
    pub id: String,
    pub title: String,
    pub description: String,
    pub created_at: Option<chrono::DateTime<chrono::Utc>>,
    pub updated_at: Option<chrono::DateTime<chrono::Utc>>,
}

#[allow(non_snake_case)]
#[derive(Debug, Deserialize, Serialize)]
pub struct UsherGroupModelResponse {
    pub id: String,
    pub title: String,
    pub description: String,
    pub createdAt: chrono::DateTime<chrono::Utc>,
    pub updatedAt: chrono::DateTime<chrono::Utc>,
}
