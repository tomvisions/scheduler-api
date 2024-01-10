use std::sync::Arc;

//use sqlx::postgres::PgPoolOptions;
///use sqlx::PgPool;
use sqlx::mysql::MySqlPoolOptions;
use sqlx::mysql;
use sqlx::MySqlPool;
//use crate::SETTINGS;

/// Storing the state of the application
/// Can be accessed using the AppData extractor.
#[derive(Clone)]
pub struct AppState {
    pub creds: argon2_creds::Config,
    // sqlx
    pub db: MySqlPool,
}

impl AppState {
    pub async fn new() -> Arc<Self> {
        let creds = argon2_creds::ConfigBuilder::default()
            .username_case_mapped(true)
            .profanity(true)
            .blacklist(true)
            .password_policy(argon2_creds::PasswordPolicy::default())
            .build()
            .unwrap();

        let c = creds.clone();

        #[cfg_attr(debug_assertions, allow(unused_variables))]
        let init = std::thread::spawn(move || {
            log::info!("Initializing credential manager");
            c.init();
            log::info!("Initialized credential manager");
        });

        
        let db = MySqlPoolOptions::new()
            .max_connections(15)
        //    .connect(&SETTINGS.database.url())
            .connect("mysql://root:net123DBMaster@103-98-215-187.cloud-xip.com:3306/scheduler")
            .await
            .expect("Unable to form database pool");
        #[cfg(not(debug_assertions))]
        init.join().unwrap();
        Arc::new(AppState { creds, db })
    }
}