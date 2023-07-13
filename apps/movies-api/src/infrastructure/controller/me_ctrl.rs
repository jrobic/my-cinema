use rocket::{get, serde::json::Json};

use crate::{domain::entity::user_entity::User, usecase};

#[utoipa::path(
	responses(
		(status = 200, description = "Return user information", body = User),
	),
	tag = "Me"
)]
#[get("/me")]
pub fn me_ctrl() -> Json<User> {
	let me_usecase = usecase::me_usecase::MeUsecase::new();
	Json(me_usecase.get_me())
}
