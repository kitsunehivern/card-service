table "cards" {
	schema = schema.public
	column "id" {
		null = false
		type = serial
		comment = "The ID of the card"
	}

	column "user_id" {
		null = false
		type = varchar(64)
		comment = "The ID of the owner"
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

	column "status" {
		null = false
		type = enum.card_status
		comment = "The status of the card (requested, active, blocked, ...)"
	}

	column "created_at" {
		null = false
		type = timestamptz
		comment = "The last time when the card was created"
	}

	column "updated_at" {
		null = false
		type = timestamptz
		comment = "The last time when the card was updated"
	}

	column "deleted_at" {
		type = timestamptz
		comment = "The last time when the card was deleted"
	}

	primary_key {
		columns = [column.id]
	}

	index "idx_user_id" {
		columns = [column.user_id]
	}
}

enum "card_status" {
	schema = schema.public
    values = ["requested", "active", "blocked", "retired", "closed"]
}
