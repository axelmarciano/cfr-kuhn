up:
	reflex -r '\.go$$' -s -- sh -c "go run cmd/main.go"

test_app:
	$(MAKE_COVERAGE_CMD)

test_app_watch:
	find . -name '*.go' | entr -n -c $(MAKE) test_app $(DOCKER_FLAG) $(HTML_FLAG)


define MAKE_COVERAGE_CMD
	go test -v -coverprofile=coverage.out ./... && \
	$(call CLEAN_COVERAGE) && \
	$(call GENERATE_HTML)
endef

define CLEAN_COVERAGE
	if [ "$(shell uname -s)" = "Darwin" ]; then \
		sed -i '' -e '/test/d' -e '/cmd/d' coverage.out; \
	else \
		sed -i '/test/d;/cmd/d;' coverage.out; \
	fi
endef

define GENERATE_HTML
	if [ "$(HTML_FLAG)" = "html" ]; then \
		go tool cover -html=coverage.out -o coverage.html && \
		echo 'Coverage report generated: coverage.html'; \
	fi
endef

.PHONY: docker html
