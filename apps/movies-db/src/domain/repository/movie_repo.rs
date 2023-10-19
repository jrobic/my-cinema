use axum::async_trait;

use crate::domain::entity::movie_entity::Movie;

#[derive(Debug)]
pub enum FindAllMoviesError {
	DBInternalError,
}

#[async_trait]
pub trait MovieRepository: Send + Sync {
	async fn find_all_movies(&self) -> Result<Vec<Movie>, FindAllMoviesError>;
}
