# Option
#===============================================================
APP_ENV              :=
TAG_OPTION           :=
PROFILE_OPTION       :=

# Const
#===============================================================
PRODUCT_NAME         := LycleL
STACK_NAME           := MiniApp
ENTRY_POINT          := miniapp
IMAGE_REPO           := $(CDK_DEFAULT_ACCOUNT).dkr.ecr.ap-northeast-1.amazonaws.com/

# include build-go.mk deploy-ecs.mk docker-ecr.mk
include .make/*.mk


# internal task
#===============================================================

## aws cliのprofile設定
.set-profile:
ifneq ($(PROFILE),)
	$(eval PROFILE_OPTION := --profile $(PROFILE))
endif