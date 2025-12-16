
docker_registry = "localhost:15444"
docker_namespace = "default"

image_tag = "latest"

api_image_ref = '%s/%s/hello-world-api:%s' % (docker_registry, docker_namespace, image_tag)

local_resource(
    'build-api',
    cmd='docker build -t %s -f ./api/Dockerfile ./api --push' % api_image_ref,

)

db_image_ref = '%s/%s/hello-world-db:%s' % (docker_registry, docker_namespace, image_tag)

local_resource(
    'build-db',
    cmd='docker build -t %s -f ./db/Dockerfile ./db --push' % db_image_ref,
)

web_image_ref = '%s/%s/hello-world-app:%s' % (docker_registry, docker_namespace, image_tag)

local_resource(
    'build-web',
    cmd='docker build -t %s -f ./my-app/Dockerfile ./my-app --push' % web_image_ref,

)

# api svc
k8s_yaml('api/k8s_yaml/deployment.yaml',
    allow_duplicates=True,
)

k8s_yaml('api/k8s_yaml/service.yaml',
    allow_duplicates=True,
)
k8s_yaml('api/k8s_yaml/ingress.yaml',
    allow_duplicates=True,
)

# db svc
k8s_yaml('db/k8s_yaml/deployment.yaml',
    allow_duplicates=True,
)

k8s_yaml('db/k8s_yaml/service.yaml',
    allow_duplicates=True,
)


# web svc
k8s_yaml('my-app/k8s_yaml/deployment.yaml',
    allow_duplicates=True,
)

k8s_yaml('my-app/k8s_yaml/service.yaml',
    allow_duplicates=True,
)

k8s_yaml('my-app/k8s_yaml/ingress.yaml',
    allow_duplicates=True,
)