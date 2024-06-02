SHELL=/bin/bash

GO_DIR := go
GO_FILE := main.go
GO_OUTPUT := libgo.so

all:
	@make help


.PHONY: help ## Display help commands
help:
	@printf 'Usage:\n'
	@printf '  make <tagert>\n'
	@printf '\n'
	@printf 'Targets:\n'
	@IFS=$$'\n' ; \
    help_lines=(`fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//'`); \
    for help_line in $${help_lines[@]}; do \
        IFS=$$'#' ; \
        help_split=($$help_line) ; \
        help_info=`echo $${help_split[2]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
		IFS=$$':' ; \
		phony_command=($$help_split); \
        help_command=`echo $${phony_command[1]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
		printf "  %-50s %s\n" $$help_command $$help_info ; \
    done

# clippy

.PHONY: clippy ## cargo clippy
clippy:
	SKIP_WASM_BUILD=1 cargo clippy --workspace --all-targets --all-features -- -D warnings

.PHONY: clippyfix ## cargo clippy --fix
clippyfix:
	SKIP_WASM_BUILD=1 cargo clippy --allow-dirty --allow-staged --fix --workspace --all-targets --all-features -- -D warnings

.PHONY: cargofix ## cargo fix
cargofix:
	cargo fix --allow-dirty --allow-staged --workspace --all-targets --all-features


# toml && lock check
.PHONY: tomlcheck ## cargo toml check
tomlcheck:
	cargo audit fix --dry-run

# format

.PHONY: fmtcheck ## cargo fmt check
fmtcheck:
	cargo fmt --all -- --check

.PHONY: taplocheck ## taplo fmt check
#  cargo install taplo-cli --locked
taplocheck:
	taplo fmt --check

.PHONY: fmt ## cargo fmt all && taplo fmt all
fmt:
	cargo fmt --all && taplo fmt


# build
.PHONY: build ## build node
build:
	make gobuild && cargo build --release

.PHONY: gobuild
gobuild:
	cd $(GO_DIR) && go build -o $(GO_OUTPUT) -buildmode=c-shared $(GO_FILE)

.PHONY: run
	cargo run

.PHONY: clean
clean :
	rm -f go/libgo.*
	cargo clean
