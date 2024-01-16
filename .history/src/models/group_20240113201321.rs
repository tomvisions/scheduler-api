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

pub struct GroupResponse {
    pub status: String,
    pub data: String,
}


fn filter_db_record(note: &Group) -> NoteModelResponse {
    NoteModelResponse {
        id: note.id.to_owned(),
        title: note.title.to_owned(),
        content: note.content.to_owned(),
        category: note.category.to_owned().unwrap(),
        published: note.published != 0,
        createdAt: note.created_at.unwrap(),
        updatedAt: note.updated_at.unwrap(),
    }
}
