#[macro_use]
extern crate rocket;

use std::fs;

use rocket::{launch, routes};
use utoipa::OpenApi;
use utoipa_swagger_ui::SwaggerUi;

mod domain;
mod infrastructure;
mod usecase;

#[launch]
async fn rocket() -> _ {
	let doc = infrastructure::api_doc::ApiDoc::openapi();
	fs::write("openapi.json", doc.to_pretty_json().unwrap()).expect("Unable to write file");

	rocket::build()
		.mount(
			"/",
			SwaggerUi::new("/swagger-ui/<_..>").url("/api-docs/openapi.json", doc),
		)
		.mount(
			"/",
			routes![
				infrastructure::controller::common_ctrl::health,
				infrastructure::controller::common_ctrl::headers,
				infrastructure::controller::me_ctrl::me_ctrl
			],
		)
		.register(
			"/",
			catchers![
				infrastructure::controller::catchers_ctrl::not_found_ctrl,
				infrastructure::controller::catchers_ctrl::catch_default_ctrl
			],
		)
}
