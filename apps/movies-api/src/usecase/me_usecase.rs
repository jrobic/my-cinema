use crate::domain::entity::user_entity::User;

pub struct MeUsecase {}

impl MeUsecase {
	pub fn new() -> Self {
		Self {}
	}

	pub fn get_me(&self) -> User {
		User {
			first_name: "John".to_string(),
			last_name: "Doe".to_string(),
			age: 30,
		}
	}
}
