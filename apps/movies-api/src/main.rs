use rocket::{get, launch, routes};

#[get("/health")]
fn health() -> &'static str {
	"OK"
}

#[launch]
async fn rocket() -> _ {
	rocket::build().mount("/", routes![health])
}
