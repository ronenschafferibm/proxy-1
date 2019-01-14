# Developing locally istio filters and cilium filters

This section describes how to build envoy locally (in a docker image) with a local repo of istio filters and a local repo of cilium filters (current repo)


1. Create a local repo of istio filters in directory istio_proxy
    ```
    git submodule add git@github.com:istio/proxy.git istio_proxy
    ```

1. Checkout the correct revision
    ```
    (cd istio_proxy && git checkout 67a0375be569f9158b361e8f5c2a76a0c1b0a02e)
    ```

1. Change istio_proxy from an `http_archive` to a `local_repository` in `WORKSPACE`:

    Comment out:
    ```
    http_archive(
        name = "istio_proxy",
        url = "https://github.com/istio/proxy/archive/" + ISTIO_PROXY_SHA + ".zip",
        sha256 = ISTIO_PROXY_SHA256,
        strip_prefix = "proxy-" + ISTIO_PROXY_SHA,
    )
    ```
    and add:
    ```
    local_repository(
        name="istio_proxy",
        path="./istio_proxy/"
    )
    ```

1. Build docker image builder
    ```
    make docker-image-envoy-local-dev
    ```

1. Run the docker image builder from the following script
    ```
    ./run_envoy_docker.sh
    ```
    it runs the docker image builder and maps the current directory into the container
    
1. From inside the container, build envoy:
    ```
    make cilium-envoy
    ```

1. From inside the container or outside of it, run envoy:
    ```
    ./cilium-envoy
    ```
    
1. Now you can make changes to the source code of istio filters and cilium filters, and build cilium-envoy with the above container


## Quick summary of make commands to build docker images:

| make command    | docker file | tag | purpose |
| --------------- | ----------- | --- |---------|
| docker-image-builder | Dockerfile.builder | cilium-envoy-builder | build envoy dependencies and artifacts (might take a long time) |
| docker-image-envoy | Dockerfile | cilium-envoy | builds cilium filters and istio filters on top of cilium-envoy-builder image |
| docker-istio-proxy | Dockerfile_istio_proxy | istio-proxy | builds an istio-proxy image. it basically uses the current istio-proxy docker image and replaces its envoy executable with the cilium-envoy version | 
| docker-image-envoy-local-dev | Dockerfile.local_dev | docker-image-envoy-local-dev | **This is my addition** builds envoy with the local repo of istio filters |

