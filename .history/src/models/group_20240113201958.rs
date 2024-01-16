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




fn filter_db_record(group: &Group) -> GroupResponse {
    NoteModelResponse {
        id: note.id.to_owned(),
        title: note.title.to_owned(),
        description: note.content.to_owned(),
        createdAt: note.created_at.unwrap(),
        updatedAt: note.updated_at.unwrap(),
    }
}
