table "cards" {
	schema = schema.public

	column "id" {
		null = false
		type = integer
		comment = "The ID of the card"
	}

	column "user_id" {
		null = false
		type = integer
		comment = "The ID of the owner"
	}

	column "type" {
		null = false
		type = enum.card_type
		comment = "The type of the card"
	}

	column "credit" {
		null = false
		type = integer
		comment = "The credit of the card"
	}

	column "debit" {
		null = false
		type = integer
		comment = "The debit of the card"
	}

	column "expiration_date" {
		null = false
		type = timestamp
		comment = "The expiration date of the card"
	}

	column "status" {
		null = false
		type = enum.card_status
		comment = "The status of the card (requested, active, blocked, ...)"
	}

	column "created_at" {
		null = false
		type = timestamp
		comment = "The time when the card was first created"
	}

	column "updated_at" {
		null = false
		type = timestamp
		comment = "The time when the card was last updated"
	}

	primary_key {
		columns = [column.id]
	}

	foreign_key "fk_user_id" {
		columns = [column.user_id]
		ref_columns = [table.users.column.id]
		on_delete = CASCADE
	}

	index "idx_user_id" {
		columns = [column.user_id]
	}

	index "idx_expiration_date" {
		columns = [column.expiration_date]
	}
}

enum "card_status" {
	schema = schema.public
    values = ["requested", "active", "blocked", "expired", "closed"]
}

enum "card_type" {
	schema = schema.public
	values = ["gold", "diamond", "platinum"]
}
