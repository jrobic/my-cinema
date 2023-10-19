use std::sync::Mutex;

use axum::async_trait;
use uuid::Uuid;

use crate::domain::{
	entity::movie_entity::Movie,
	repository::movie_repo::{FindAllMoviesError, MovieRepository},
};

pub struct MovieInmemoryRepository {
	movies: Mutex<Vec<Movie>>,
}

impl MovieInmemoryRepository {
	#![allow(dead_code)]
	pub fn new() -> Self {
		Self {
			movies: Mutex::new(vec![
				Movie {
					id: Uuid::new_v4(),
					tmdb_id: 758323,
					title: "The Pope's Exorcist".to_string(),
					genres: vec!["Horror".to_string(), "Mystery".to_string(), "Thriller".to_string()],
					original_language: "en".to_string(),
					overview: "Father Gabriele Amorth Chief Exorcist of the Vatican investigates a young boy's terrifying possession and ends up uncovering a centuries-old conspiracy the Vatican has desperately tried to keep hidden.".to_string(),
					popularity: 5953.227,
					production_companies: vec!["Screen Gems".to_string(), "2.0 Entertainment".to_string() ,"Jesus & Mary".to_string(), "Worldwide Katz".to_string(), "Loyola Productions".to_string(), "FFILME.RO".to_string()],
					release_date: "2023-04-05".to_string(),
					budget: 18000000,
					revenue: 65675816,
					runtime: 103,
					status: "Released".to_string(),
					tagline: "Inspired by the actual files of Father Gabriele Amorth, Chief Exorcist of the Vatican.".to_string(),
					vote_average: 7.433,
					vote_count: 545,
					backdrop_path: "/9JBEPLTPSm0d1mbEcLxULjJq9Eh.jpg".to_string(),
					poster_path: "/hiHGRbyTcbZoLsYYkO4QiCLYe34.jpg".to_string(),
					credits: "713704-296271-502356-1076605-1084225-1008005-916224-1023313-1033219-980078-842945-943822-816904-804150-638974-649609-603692-849869-809787-776835-1104040".split('-').map(|s| s.to_string()).collect(),
					keywords: "spain-rome italy-vatican-pope-pig-possession-conspiracy-devil-exorcist-skepticism-catholic priest-1980s-supernatural horror".split('-').map(|s| s.to_string()).collect(),
				 },
			]),
		}
	}
}

#[async_trait]
impl MovieRepository for MovieInmemoryRepository {
	async fn find_all_movies(&self) -> Result<Vec<Movie>, FindAllMoviesError> {
		let movies = match self.movies.lock() {
			Ok(movies) => movies,
			Err(_) => return Err(FindAllMoviesError::DBInternalError),
		};

		Ok(movies.to_vec())
	}
}
