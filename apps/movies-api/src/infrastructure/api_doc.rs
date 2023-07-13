use utoipa::OpenApi;

use crate::domain::entity::user_entity::User;

#[derive(OpenApi)]
#[openapi(
	paths(
		crate::infrastructure::controller::common_ctrl::health,
		crate::infrastructure::controller::me_ctrl::me_ctrl,
	),
	components(schemas(User)),
	security()
)]
pub struct ApiDoc;
