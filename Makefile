.PHONY: help up clean

help:
	@echo ""
	@echo "make run DIR='<service>'   - Run a day e.g. make run DIR='day1'"
	@echo "make clean                 - Stop docker container and clean images"
	@echo ""

run:
	@if [ -z "$(DIR)" ]; then \
		echo "? Error: You must provide DIR='<service>'"; \
		exit 1; \
	fi
	@echo "? Running service: $(DIR)"
	SERVICE_DIR=$(DIR) docker compose up --build

clean:
	docker-compose down --rmi local --volumes --remove-orphans

