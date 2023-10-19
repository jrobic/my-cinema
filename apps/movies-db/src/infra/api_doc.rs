use utoipa::OpenApi;

use crate::domain::entity::health_entity::Health;

#[derive(OpenApi)]
#[openapi(
	paths(crate::infra::controller::common_ctrl::health),
	components(schemas(Health)),
	security()
)]
pub struct ApiDoc;