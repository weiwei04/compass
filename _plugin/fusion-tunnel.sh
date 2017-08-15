#!/bin/bash
export COMPASS_ADDR=$TILLER_HOST
$HELM_PLUGIN_DIR/fusion $@
