use serde::{Deserialize, Serialize};

#[derive(Default, Debug, Serialize, Clone, Deserialize)]
pub struct Movie {
	pub id: uuid::Uuid,
	pub tmdb_id: i64,
	pub title: String,
	pub overview: String,
	pub original_language: String,
	pub genres: Vec<String>,
	pub release_date: String,
	pub popularity: f64,
	pub production_companies: Vec<String>,
	pub budget: i64,
	pub revenue: i64,
	pub runtime: i64,
	pub status: String,
	pub tagline: String,
	pub vote_average: f64,
	pub vote_count: i64,
	pub credits: Vec<String>,
	pub keywords: Vec<String>,
	pub poster_path: String,
	pub backdrop_path: String,
}
