ROOT_DIR=$(shell pwd)
OUT_DIR=$(ROOT_DIR)/out
LOGS_DIR=$(ROOT_DIR)/logs

#
MODULE_NAME=github.com/ThoPham02/research_management
SERVICE=service
SERVICE_DIR=$(ROOT_DIR)/$(SERVICE)
API_DIR=$(ROOT_DIR)/api
MIGRATION_DIR=$(ROOT_DIR)/migration
DATASOURCE=postgres://thopham:Lk8UZozxGk0BfPfAdfjsjUsptsYtMcKu@dpg-cgkogb8rddleudvr9ee0-a.singapore-postgres.render.com/nckhdb

# table account api
USER=user_tbl

# table topic api
DEPARTMENT=department_tbl
FACULTY=faculty_tbl
TOPIC=topic_tbl

#table progress api
NOTIFICATION=notification_tbl
LIBRARY=library_tbl
TOPIC_REPORT=topic_report_tbl
EVENT=event_tbl
STAGE=stage_tbl

# table result api
SUBCOMMITTEE=subcommittee_tbl
GROUP=group_tbl
TOPIC_MARK=topic_mark_tbl

# account service
ACCOUNT_DIR=$(SERVICE_DIR)/account
ACCOUNT_API_DIR=$(ACCOUNT_DIR)/api
ACCOUNT_MODEL_DIR=$(ACCOUNT_DIR)/model

# topic service
TOPIC_DIR=$(SERVICE_DIR)/topic
TOPIC_API_DIR=$(TOPIC_DIR)/api
TOPIC_MODEL_DIR=$(TOPIC_DIR)/model

# result service
RESULT_DIR=$(SERVICE_DIR)/result
RESULT_API_DIR=$(RESULT_DIR)/api
RESULT_MODEL_DIR=$(RESULT_DIR)/model

# progress service
PROGRESS_DIR=$(SERVICE_DIR)/progress
PROGRESS_API_DIR=$(PROGRESS_DIR)/api
PROGRESS_MODEL_DIR=$(PROGRESS_DIR)/model

dep-init:
	go mod init $(MODULE_NAME)
	go mod tidy


dep:
	go mod tidy


#gen api code
gen-account-service:
	goctl api go -api $(API_DIR)/account.api -dir $(ACCOUNT_API_DIR)

gen-result-service:
	goctl api go -api $(API_DIR)/result.api -dir $(RESULT_API_DIR)

gen-progress-service:
	goctl api go -api $(API_DIR)/progress.api -dir $(PROGRESS_API_DIR)

gen-topic-service:
	goctl api go -api $(API_DIR)/topic.api -dir $(TOPIC_API_DIR)

gen-services: gen-account-service gen-result-service gen-progress-service gen-topic-service

# migrate 
migrate:
	migrate create -ext sql -dir ${MIGRATION_DIR} -seq update_student_group
migrateup:
	migrate -path ${MIGRATION_DIR} -database "${DATASOURCE}" -verbose up
migratedown:
	migrate -path ${MIGRATION_DIR} -database "${DATASOURCE}" -verbose down

# gen db model
gen-account-model: 
	goctl model pg datasource -url="${DATASOURCE}" -table="${USER}" -dir="$(ACCOUNT_MODEL_DIR)"

gen-result-model: 
	goctl model pg datasource -url="${DATASOURCE}" -table="${TOPIC_MARK}" -dir="$(RESULT_MODEL_DIR)"
	goctl model pg datasource -url="${DATASOURCE}" -table="${GROUP}" -dir="$(RESULT_MODEL_DIR)"
	goctl model pg datasource -url="${DATASOURCE}" -table="${SUBCOMMITTEE}" -dir="$(RESULT_MODEL_DIR)"

gen-progress-model: 
	goctl model pg datasource -url="${DATASOURCE}" -table="${TOPIC_REPORT}" -dir="$(PROGRESS_MODEL_DIR)"
	goctl model pg datasource -url="${DATASOURCE}" -table="${NOTIFICATION}" -dir="$(PROGRESS_MODEL_DIR)"
	goctl model pg datasource -url="${DATASOURCE}" -table="${LIBRARY}" -dir="$(PROGRESS_MODEL_DIR)"
	goctl model pg datasource -url="${DATASOURCE}" -table="${EVENT}" -dir="$(PROGRESS_MODEL_DIR)"
	goctl model pg datasource -url="${DATASOURCE}" -table="${STAGE}" -dir="$(PROGRESS_MODEL_DIR)"

gen-topic-model: 
	goctl model pg datasource -url="${DATASOURCE}" -table="${TOPIC}" -dir="$(TOPIC_MODEL_DIR)"
	goctl model pg datasource -url="${DATASOURCE}" -table="${DEPARTMENT}" -dir="$(TOPIC_MODEL_DIR)"
	goctl model pg datasource -url="${DATASOURCE}" -table="${FACULTY}" -dir="$(TOPIC_MODEL_DIR)"

gen-models: gen-account-model gen-result-model gen-progress-model gen-topic-model

runs:
	go run main.go -f etc/server.yaml