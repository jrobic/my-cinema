use std::{net::SocketAddr, sync::Arc};

use axum::{routing::get, Router};
use domain::repository::movie_repo::MovieRepository;
use std::fs;

use tokio::signal;
use tower_http::trace::{self, TraceLayer};
use tracing::Level;
use tracing_subscriber::{layer::SubscriberExt, util::SubscriberInitExt};
use utoipa::OpenApi;
use utoipa_swagger_ui::SwaggerUi;

use crate::infra::repository;

mod domain;
mod infra;
mod usecase;

#[tokio::main]
async fn main() {
	let doc = infra::api_doc::ApiDoc::openapi();
	fs::write("openapi.json", doc.to_pretty_json().unwrap()).expect("Unable to write file");

	tracing_subscriber::registry()
		.with(
			tracing_subscriber::EnvFilter::try_from_default_env().unwrap_or_else(|_| {
				"movies_db=debug,tower_http=debug,axum::rejection=trace".into()
			}),
		)
		.with(tracing_subscriber::fmt::layer().compact())
		.init();

	let app_state = AppState {
		movie_repo: Arc::new(repository::movie_inmemory_repo::MovieInmemoryRepository::new()),
	};

	let movie_router = Router::new().route(
		"/movies",
		get(infra::controller::movies_ctrl::get_all_movies),
	);

	let app = Router::new()
		.merge(SwaggerUi::new("/swagger-ui").url("/api-docs/openapi.json", doc))
		.merge(movie_router)
		.route("/health", get(infra::controller::common_ctrl::health))
		.fallback(infra::controller::catchers_ctrl::not_found_ctrl)
		.with_state(app_state)
		.layer(
			TraceLayer::new_for_http()
				.make_span_with(trace::DefaultMakeSpan::new().level(Level::INFO))
				.on_response(trace::DefaultOnResponse::new().level(Level::INFO)),
		);

	let addr = SocketAddr::from(([127, 0, 0, 1], 3000));
	tracing::info!("Listening on {}", addr);

	axum::Server::bind(&addr)
		.serve(app.into_make_service())
		.with_graceful_shutdown(shutdown_signal())
		.await
		.unwrap();
}

#[derive(Clone)]
pub struct AppState {
	pub movie_repo: Arc<dyn MovieRepository + Send + Sync>,
}

async fn shutdown_signal() {
	let ctrl_c = async {
		signal::ctrl_c().await.expect("failed to install Ctrl+C handler");
	};

	#[cfg(unix)]
	let terminate = async {
		signal::unix::signal(signal::unix::SignalKind::terminate())
			.expect("failed to install signal handler")
			.recv()
			.await;
	};

	#[cfg(not(unix))]
	let terminate = std::future::pending::<()>();

	tokio::select! {
		_ = ctrl_c => {},
		_ = terminate => {},
	}

	tracing::info!("signal received, starting graceful shutdown");
}
