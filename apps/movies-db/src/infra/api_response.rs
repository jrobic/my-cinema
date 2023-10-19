use axum::{
	http::StatusCode,
	response::{IntoResponse, Response},
	Json,
};
use erased_serde::Serialize as ErasedSerialize;
use serde::Serialize;

pub enum ApiResponseType {
	SuccessWithData,
	#[allow(dead_code)]
	StatusCodeOnly,
	#[allow(dead_code)]
	Error,
}

impl Default for ApiResponseType {
	fn default() -> Self {
		Self::SuccessWithData
	}
}

pub enum ApiResponseData<T: Serialize> {
	Data {
		data: T,
		status: StatusCode,
	},
	#[allow(dead_code)]
	Error {
		error: ApiResponseError,
		status: StatusCode,
	},
	#[allow(dead_code)]
	StatusCode(StatusCode),
}

impl<T> ApiResponseData<T>
where
	T: Serialize + 'static,
{
	pub fn success_with_data(data: T, status: StatusCode) -> Self {
		Self::Data { data, status }
	}

	#[allow(dead_code)]
	pub fn status_code(status: StatusCode) -> Self {
		Self::StatusCode(status)
	}
	#[allow(dead_code)]
	pub fn error(error: Option<T>, message: &'static str, status: StatusCode) -> Self {
		match error {
			Some(err) => Self::Error {
				error: ApiResponseError::complicated_error(message, err),
				status,
			},
			None => Self::Error {
				error: ApiResponseError::simple_error(message),
				status,
			},
		}
	}
}

impl<T> IntoResponse for ApiResponseData<T>
where
	T: Serialize,
{
	fn into_response(self) -> Response {
		match self {
			ApiResponseData::Data { data, status } => (
				status,
				Json(ApiResponseObject::<T> {
					data: Some(data),
					error: None,
				}),
			)
				.into_response(),
			ApiResponseData::Error { error, status } => (
				status,
				Json(ApiResponseObject::<T> {
					data: None,
					error: Some(error.into()),
				}),
			)
				.into_response(),
			ApiResponseData::StatusCode(status) => status.into_response(),
		}
	}
}

#[derive(Serialize)]
pub struct ApiResponseObject<T>
where
	T: Serialize,
{
	data: Option<T>,
	error: Option<ApiResponseErrorObject>,
}

pub type ApiResponse<T, E> = Result<ApiResponseData<T>, ApiResponseData<E>>;

#[derive(Serialize)]
pub struct ApiResponseErrorObject {
	pub message: String,
	pub error: Option<Box<dyn ErasedSerialize>>,
}

pub enum ApiResponseError {
	Simple(String),
	Complicated {
		message: String,
		error: Box<dyn ErasedSerialize>,
	},
}

impl ApiResponseError {
	pub fn simple_error(msg: &'static str) -> Self {
		Self::Simple(msg.into())
	}
	pub fn complicated_error(msg: &'static str, error: impl Serialize + 'static) -> Self {
		Self::Complicated {
			message: msg.into(),
			error: Box::new(error),
		}
	}
}

impl From<ApiResponseError> for ApiResponseErrorObject {
	fn from(val: ApiResponseError) -> Self {
		match val {
			ApiResponseError::Simple(message) => Self {
				message,
				error: None,
			},
			ApiResponseError::Complicated { message, error } => Self {
				message,
				error: Some(error),
			},
		}
	}
}
