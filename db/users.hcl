table "users" {
	schema = schema.public

	column "id" {
		null = false
		type = integer
		comment = "The ID of the user"
	}

	column "name" {
		null = false
		type = varchar(64)
		comment = "The full name of the user"
	}

	column "phone_number" {
		null = false
		type = varchar(16)
		comment = "The phone number of the user"
	}

	column "password_hash" {
		null = false
		type = varchar(64)
		comment = "The hash value (+ salt value) of the password"
	}

	column "created_at" {
		null = false
		type = timestamp
		comment = "The time when the user was first created"
	}

	column "updated_at" {
		null = false
		type = timestamp
		comment = "The time when the user was last updated"
	}

	primary_key {
		columns = [column.id]
	}

	index "idx_phone_number" {
		columns = [column.phone_number]
	}
}
