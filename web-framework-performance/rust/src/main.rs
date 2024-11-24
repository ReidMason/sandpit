use std::net::SocketAddr;

use axum::{http::StatusCode, routing::get, Router};

#[tokio::main]
async fn main() {
    let app = Router::new().route("/health", get(root));

    let addr = SocketAddr::from(([127, 0, 0, 1], 8080));
    axum::Server::bind(&addr)
        .serve(app.into_make_service())
        .await
        .unwrap();
}

async fn root() -> (StatusCode, String) {
    return (StatusCode::OK, "Ok".to_string());
}
