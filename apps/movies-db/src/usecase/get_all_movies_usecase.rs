use std::sync::Arc;

use crate::domain::{entity::movie_entity::Movie, repository::movie_repo::MovieRepository};

pub enum MovieException {
	Unknown,
}

pub struct GetAllMoviesUsecase<'a> {
	pub movie_repo: &'a Arc<dyn MovieRepository + Send + Sync>,
}

impl<'a> GetAllMoviesUsecase<'a> {
	pub fn new(movie_repo: &'a Arc<dyn MovieRepository + Send + Sync>) -> Self {
		Self { movie_repo }
	}

	pub async fn execute(&self) -> Result<Vec<Movie>, MovieException> {
		let movies = match self.movie_repo.find_all_movies().await {
			Ok(movies) => movies,
			Err(_) => return Err(MovieException::Unknown),
		};

		Ok(movies)
	}
}
