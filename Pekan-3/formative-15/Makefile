build:
	@cd apps && pwd && ./build.sh

run:
	@cd apps && pwd && ./bootstrap &

tfl:
	@cd infra && pwd
	@cd infra && terraform fmt -recursive

tfi:
	@cd infra/environments/01_prod && pwd
	@cd infra/environments/01_prod && terraform init

tfp:
	@cd infra/environments/01_prod && pwd
	@cd infra/environments/01_prod && terraform plan

tfa:
	@cd infra/environments/01_prod && pwd
	@cd infra/environments/01_prod && terraform apply

tfd:
	@cd infra/environments/01_prod && pwd
	@cd infra/environments/01_prod && terraform destroy

tfs:
	@cd infra/environments/01_prod && pwd
	@cd infra/environments/01_prod && terraform state list
