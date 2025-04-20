module github.com/PlutusCLI

go 1.23.2

replace input => ./user_inputs

require (
	input v0.0.0-00010101000000-000000000000
	save_load v0.0.0-00010101000000-000000000000
)

replace save_load => ./saving_loading
