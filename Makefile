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
	docker compose run --rm -T aoc2025 sh -c "go build -o /tmp/app ./$(DAY) && /tmp/app"

clean:
	docker compose rm -sv aoc2025

