use ::serde::Serialize;
use utoipa::ToSchema;

#[derive(ToSchema, Serialize)]
#[serde(rename_all = "camelCase", crate = "rocket::serde")]
pub struct User {
	pub first_name: String,
	pub last_name: String,
	pub age: u8,
}
