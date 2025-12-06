.PHONY: help up clean

help:
	@echo ""
	@echo "make run DAY='<day>'   - Run a day e.g. make run DAY='day1'"
	@echo "make clean             - Stop docker container and clean images"
	@echo ""

run:
	@if [ -z "$(DAY)" ]; then \
		echo "? Error: You must provide DAY='<day>'"; \
		exit 1; \
	fi
	@echo "? Running service: $(DAY)"
	docker compose build $(DAY)
	docker compose run -T --rm $(DAY)

clean:
	docker compose down --rmi local --volumes --remove-orphans

