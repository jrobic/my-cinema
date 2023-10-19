use axum::{extract::State, http::StatusCode};
use serde::Serialize;

use crate::{
	domain::entity::movie_entity::Movie,
	infra::api_response::{ApiResponse, ApiResponseData},
	usecase::get_all_movies_usecase::{GetAllMoviesUsecase, MovieException},
	AppState,
};

impl<E> From<MovieException> for ApiResponseData<E>
where
	E: Serialize + 'static,
{
	fn from(value: MovieException) -> Self {
		match value {
			MovieException::Unknown => {
				ApiResponseData::status_code(StatusCode::INTERNAL_SERVER_ERROR)
			},
		}
	}
}

#[derive(Serialize)]
pub struct GetAllMoviesResponse {
	movies: Vec<Movie>,
}

pub async fn get_all_movies(
	State(app_state): State<AppState>,
) -> ApiResponse<GetAllMoviesResponse, ()> {
	let gell_all_movies_usecase = GetAllMoviesUsecase::new(&app_state.movie_repo);

	let movies = gell_all_movies_usecase.execute().await?;

	Ok(ApiResponseData::success_with_data(
		GetAllMoviesResponse { movies },
		StatusCode::OK,
	))
}
