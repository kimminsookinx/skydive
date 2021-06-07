OVN_VERSION ?= branch-20.12
SCHEMA_REMOTE_FILE = https://raw.githubusercontent.com/ovn-org/ovn/$(OVN_VERSION)/ovn-nb.ovsschema
MODEL_DIR = topology/probes/ovn/ovnmodel/
GEN_DIR = topology/probes/ovn/ovnmetagen
GEN_BIN = $(GEN_DIR)/ovnmetagen
SCHEMA_FILE = $(MODEL_DIR)/ovn-nb.ovsschema


$(SCHEMA_FILE):
	curl -o $@ $(SCHEMA_REMOTE_FILE)

$(GEN_BIN):
	go build -o $(GEN_DIR) ./$(GEN_DIR)

.PHONY: .ovnmodel
.ovnmodel: $(SCHEMA_FILE) $(GEN_BIN)
ifeq ($(WITH_OVN), true)
	@go generate -tags "${BUILD_TAGS}" topology/probes/ovn/ovnmodel/gen.go
endif

.PHONY: .ovnmodel.clean
.ovnmodel.clean:
ifeq ($(WITH_OVN), true)
	@find $(MODEL_DIR) -name "*.go" | xargs grep 'generated by ovnmetagen'  | cut -d: -f 1 | xargs rm -f
	@rm -f $(SCHEMA_FILE)
	@rm -f $(GEN_BIN)
endif
