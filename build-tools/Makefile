DOCKER_REGISTRY ?= yiannisgkoufas
DOCKER_REGISTRY_URL ?= docker.io
#this is only the case with docker.io
#if you have any other repo:

#DOCKER_REGISTRY ?= my-repo.com
#DOCKER_REGISTRY_URL ?= my-repo.com

USERNAME ?=
PASSWORD ?=

ARCH ?=

EXTERNAL_ATTACHER_TAG ?= v3.0.0-rc1
EXTERNAL_PROVISIONER_TAG ?= v2.0.0-rc2
NODE_DRIVER_REGISTRAR_TAG ?= v1.3.0

include helpers.mk

build-sidecars:
	$(call install_sidecar,https://github.com/kubernetes-csi/external-attacher.git,$(EXTERNAL_ATTACHER_TAG),external-attacher,$(EXTERNAL_ATTACHER_IMAGE),$(BASE_EXTERNAL_ATTACHER_IMAGE))
	$(call install_sidecar,https://github.com/kubernetes-csi/external-provisioner.git,$(EXTERNAL_PROVISIONER_TAG),external-provisioner,$(EXTERNAL_PROVISIONER_IMAGE),$(BASE_EXTERNAL_PROVISIONER_IMAGE))
	$(call install_sidecar,https://github.com/kubernetes-csi/node-driver-registrar.git,$(NODE_DRIVER_REGISTRAR_TAG),node-driver-registrar,$(NODE_DRIVER_REGISTRAR_IMAGE),$(BASE_NODE_DRIVER_REGISTRAR_IMAGE))

build-components:
	$(call install_local,dataset-operator,$(DATASET_OPERATOR_IMAGE))
	$(call install_local,generate-keys,$(GENERAGE_KEYS_IMAGE))

build-csi-plugins:
	$(call install_local,csi-s3,$(CSI_S3_IMAGE))
	$(call install_local,csi-driver-nfs,$(CSI_NFS_IMAGE))

push-sidecars:
	@docker login $(DOCKER_REGISTRY_URL) --username $(USERNAME) --password $(PASSWORD) ;\
	docker push $(EXTERNAL_PROVISIONER_IMAGE)-$(ARCH) ;\
	docker push $(EXTERNAL_ATTACHER_IMAGE)-$(ARCH) ;\
	docker push $(NODE_DRIVER_REGISTRAR_IMAGE)-$(ARCH)

push-components:
	@docker login $(DOCKER_REGISTRY_URL) --username $(USERNAME) --password $(PASSWORD) ;\
	docker push $(DATASET_OPERATOR_IMAGE)-$(ARCH) ;\
	docker push $(GENERAGE_KEYS_IMAGE)-$(ARCH)

push-csi-plugins:
	@docker login $(DOCKER_REGISTRY_URL) --username $(USERNAME) --password $(PASSWORD) ;\
	docker push $(CSI_S3_IMAGE)-$(ARCH) ;\
	docker push $(CSI_NFS_IMAGE)-$(ARCH)