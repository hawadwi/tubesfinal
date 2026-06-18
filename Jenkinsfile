// pipeline {
//     agent any

//     environment {
//         DOCKER_REGISTRY = "docker.io/dhiyauu"
//         IMAGE_TAG = "${env.BUILD_ID}"
//         KUBECONFIG = "/var/jenkins_home/.kube/config"
//         PATH = "/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
//     }

//     stages {
//         stage('1. Checkout Repo') {
//             steps {
//                 echo 'Checking out source code...'
//                 checkout scm
//             }
//         }

//         stage('2. Unit Tests') {
//             steps {
//                 echo 'Running User Service Unit Tests...'
//                 dir('user-service') {
//                     sh 'go version'
//                     sh 'go test -v ./... -skip Functional'
//                 }
        
//                 echo 'Running Order Service Unit Tests...'
//                 dir('order-service') {
//                     sh 'go version'
//                     sh 'go test -v ./... -skip Functional'
//                 }
        
//                 echo 'Running Tracking Service Unit Tests...'
//                 dir('tracking-service') {
//                     sh 'go version'
//                     sh 'go test -v ./... -skip Functional'
//                 }

//                 echo 'Running Gudang Service Unit Tests...'
//                 dir('gudang-service') {
//                     sh 'go version'
//                     sh 'go test -v ./... -skip Functional'
//                 }

//                 echo 'Running Courier Service Unit Tests...'
//                 dir('courier-service') {
//                     sh 'go version'
//                     sh 'go test -v ./... -skip Functional'
//                 }

//                 echo 'Running Report Service Unit Tests...'
//                 dir('report-service') {
//                     sh 'go version'
//                     sh 'go test -v ./... -skip Functional'
//                 }

//                 echo 'Running Payment Service Unit Tests...'
//                 dir('payment-service') {
//                     sh 'go version'
//                     sh 'go test -v ./... -skip Functional'
//                 }
//             }
//         }

//         stage('3. Lint / Vet') {
//             steps {
//                 dir('user-service') {
//                     sh 'go vet ./...'
//                 }

//                 dir('order-service') {
//                     sh 'go vet ./...'
//                 }

//                 dir('tracking-service') {
//                     sh 'go vet ./...'
//                 }

//                 dir('gudang-service') {
//                     sh 'go vet ./...'
//                 }

//                 dir('courier-service') {
//                     sh 'go vet ./...'
//                 }

//                 dir('report-service') {
//                     sh 'go vet ./...'
//                 }

//                 dir('payment-service') {
//                     sh 'go vet ./...'
//                 }
//             }
//         }

//         stage('4. Build Images') {
//             steps {
//                 echo 'Building User Service Image...'
//                 sh 'docker build -t user-service:latest ./user-service'

//                 echo 'Building Order Service Image...'
//                 sh 'docker build -t order-service:latest ./order-service'

//                 echo 'Building Tracking Service Image...'
//                 sh 'docker build -t tracking-service:latest ./tracking-service'

//                 echo 'Building Gudang Service Image...'
//                 sh 'docker build -t gudang-service:latest ./gudang-service'

//                 echo 'Building Courier Service Image...'
//                 sh 'docker build -t courier-service:latest ./courier-service'

//                 echo 'Building Report Service Image...'
//                 sh 'docker build -t report-service:latest ./report-service'

//                 echo 'Building Payment Service Image...'
//                 sh 'docker build -t payment-service:latest ./payment-service'

//                 // tagging
//                 sh "docker tag user-service:latest ${DOCKER_REGISTRY}/user-service:${IMAGE_TAG}"
//                 sh "docker tag order-service:latest ${DOCKER_REGISTRY}/order-service:${IMAGE_TAG}"
//                 sh "docker tag tracking-service:latest ${DOCKER_REGISTRY}/tracking-service:${IMAGE_TAG}"
//                 sh "docker tag gudang-service:latest ${DOCKER_REGISTRY}/gudang-service:${IMAGE_TAG}"
//                 sh "docker tag courier-service:latest ${DOCKER_REGISTRY}/courier-service:${IMAGE_TAG}"
//                 sh "docker tag report-service:latest ${DOCKER_REGISTRY}/report-service:${IMAGE_TAG}"
//                 sh "docker tag payment-service:latest ${DOCKER_REGISTRY}/payment-service:${IMAGE_TAG}"

//               // tagging latest
//               sh "docker tag user-service:latest ${DOCKER_REGISTRY}/user-service:latest"
//               sh "docker tag order-service:latest ${DOCKER_REGISTRY}/order-service:latest"
//               sh "docker tag tracking-service:latest ${DOCKER_REGISTRY}/tracking-service:latest"
//               sh "docker tag gudang-service:latest ${DOCKER_REGISTRY}/gudang-service:latest"
//               sh "docker tag courier-service:latest ${DOCKER_REGISTRY}/courier-service:latest"
//               sh "docker tag report-service:latest ${DOCKER_REGISTRY}/report-service:latest"
//               sh "docker tag payment-service:latest ${DOCKER_REGISTRY}/payment-service:latest"
//             }
//         }

//         stage('5. Functional Tests') {
//             steps {
        
//                 sh 'docker compose up -d'
        
//                 sleep time: 40, unit: 'SECONDS'
        
//                 echo 'Running User Functional Tests...'
//                 dir('user-service') {
//                     sh '''
//                     DB_HOST=host.docker.internal \
//                     DB_PORT=3306 \
//                     DB_USER=root \
//                     DB_PASSWORD=root \
//                     DB_NAME=userdb \
//                     go test -tags=functional -v -run Functional ./...
//                     '''
//                 }
        
//                 echo 'Running Order Functional Tests...'
//                 dir('order-service') {
//                     sh '''
//                     DB_HOST=host.docker.internal \
//                     DB_PORT=3306 \
//                     DB_USER=root \
//                     DB_PASSWORD=root \
//                     DB_NAME=orderdb \
//                     go test -tags=functional -v -run Functional ./...
//                     '''
//                 }
        
//                 echo 'Running Tracking Functional Tests...'
//                 dir('tracking-service') {
//                     sh '''
//                     DB_HOST=host.docker.internal \
//                     DB_PORT=3306 \
//                     DB_USER=root \
//                     DB_PASSWORD=root \
//                     DB_NAME=trackingdb \
//                     go test -tags=functional -v -run Functional ./...
//                     '''
//                 }
        
//                 echo 'Running Gudang Functional Tests...'
//                 dir('gudang-service') {
//                     sh '''
//                     DB_HOST=host.docker.internal \
//                     DB_PORT=3306 \
//                     DB_USER=root \
//                     DB_PASSWORD=root \
//                     DB_NAME=gudangdb \
//                     go test -tags=functional -v -run Functional ./...
//                     '''
//                 }
        
//                 echo 'Running Courier Functional Tests...'
//                 dir('courier-service') {
//                     sh '''
//                     DB_HOST=host.docker.internal \
//                     DB_PORT=3306 \
//                     DB_USER=root \
//                     DB_PASSWORD=root \
//                     DB_NAME=courierdb \
//                     go test -tags=functional -v -run Functional ./...
//                     '''
//                 }
        
//                 echo 'Running Report Functional Tests...'
//                 dir('report-service') {
//                     sh '''
//                     DB_HOST=host.docker.internal \
//                     DB_PORT=3306 \
//                     DB_USER=root \
//                     DB_PASSWORD=root \
//                     DB_NAME=reportdb \
//                     go test -tags=functional -v -run Functional ./...
//                     '''
//                 }

//                 echo 'Running Payment Functional Tests...'
//                 dir('payment-service') {
//                     sh '''
//                     DB_HOST=host.docker.internal \
//                     DB_PORT=3306 \
//                     DB_USER=root \
//                     DB_PASSWORD=root \
//                     DB_NAME=paymentdb \
//                     go test -tags=functional -v -run Functional ./...
//                     '''
//                 }
//             }
        
//             post {
//                 always {
//                     sh 'docker compose down'
//                 }
//             }
//         }

//         stage('6. Push Images') {
//             steps {
//                 sh "docker push ${DOCKER_REGISTRY}/user-service:${IMAGE_TAG}"
//                 sh "docker push ${DOCKER_REGISTRY}/order-service:${IMAGE_TAG}"
//                 sh "docker push ${DOCKER_REGISTRY}/tracking-service:${IMAGE_TAG}"
//                 sh "docker push ${DOCKER_REGISTRY}/gudang-service:${IMAGE_TAG}"
//                 sh "docker push ${DOCKER_REGISTRY}/courier-service:${IMAGE_TAG}"
//                 sh "docker push ${DOCKER_REGISTRY}/report-service:${IMAGE_TAG}"
//                 sh "docker push ${DOCKER_REGISTRY}/payment-service:${IMAGE_TAG}"

//                 sh "docker push ${DOCKER_REGISTRY}/user-service:latest"
//                 sh "docker push ${DOCKER_REGISTRY}/order-service:latest"
//                 sh "docker push ${DOCKER_REGISTRY}/tracking-service:latest"
//                 sh "docker push ${DOCKER_REGISTRY}/gudang-service:latest"
//                 sh "docker push ${DOCKER_REGISTRY}/courier-service:latest"
//                 sh "docker push ${DOCKER_REGISTRY}/report-service:latest"
//                 sh "docker push ${DOCKER_REGISTRY}/payment-service:latest"
//             }
//         }

//         stage('7. Deploy Kubernetes') {
//             steps {
//                 sh 'kubectl apply -f k8s/user-deployment.yaml'
//                 sh 'kubectl apply -f k8s/order-deployment.yaml'
//                 sh 'kubectl apply -f k8s/tracking-deployment.yaml'
//                 sh 'kubectl apply -f k8s/gudang-deployment.yaml'
//                 sh 'kubectl apply -f k8s/courier-deployment.yaml'
//                 sh 'kubectl apply -f k8s/report-deployment.yaml'
//                 sh 'kubectl apply -f k8s/payment-deployment.yaml'

//                 sh 'kubectl apply -f k8s/user-service.yaml'
//                 sh 'kubectl apply -f k8s/order-service.yaml'
//                 sh 'kubectl apply -f k8s/tracking-service.yaml'
//                 sh 'kubectl apply -f k8s/gudang-service.yaml'
//                 sh 'kubectl apply -f k8s/courier-service.yaml'
//                 sh 'kubectl apply -f k8s/report-service.yaml'
//                 sh 'kubectl apply -f k8s/payment-service.yaml'
//             }
//         }

//         stage('8. Verify') {
//             steps {
        
//                 sh 'kubectl rollout restart deployment/user-service'
//                 sh 'kubectl rollout restart deployment/order-service'
//                 sh 'kubectl rollout restart deployment/tracking-service'
//                 sh 'kubectl rollout restart deployment/gudang-service'
//                 sh 'kubectl rollout restart deployment/courier-service'
//                 sh 'kubectl rollout restart deployment/report-service'
//                 sh 'kubectl rollout restart deployment/payment-service'
        
//                 sleep time: 20, unit: 'SECONDS'
        
//                 sh 'kubectl get pods'
//                 sh 'kubectl get svc'
        
//                 sh 'kubectl rollout status deployment/user-service --timeout=300s'
//                 sh 'kubectl rollout status deployment/order-service --timeout=300s'
//                 sh 'kubectl rollout status deployment/tracking-service --timeout=300s'
//                 sh 'kubectl rollout status deployment/gudang-service --timeout=300s'
//                 sh 'kubectl rollout status deployment/courier-service --timeout=300s'
//                 sh 'kubectl rollout status deployment/report-service --timeout=300s'
//                 sh 'kubectl rollout status deployment/payment-service --timeout=300s'
//             }
//         }
//     }

//     post {
//         success {
//             echo 'Pipeline successfully executed!'
//         }
//         failure {
//             echo 'Pipeline failed. Check the logs above.'
//         }
//         unstable {
//             echo 'Pipeline is unstable (Functional test failed as expected).'
//         }
//     }
// }







pipeline {
agent any

```
environment {
    DOCKER_USERNAME = "umarx"
    IMAGE_TAG = "${BUILD_NUMBER}"
}

stages {

    stage('Checkout') {
        steps {
            checkout scm
        }
    }

    stage('User Service Test') {
        steps {
            dir('user-service') {
                sh 'go test ./...'
            }
        }
    }

    stage('Order Service Test') {
        steps {
            dir('order-service') {
                sh 'go test ./...'
            }
        }
    }

    stage('Build User Service Image') {
        steps {
            sh "docker build -t ${DOCKER_USERNAME}/user-service:latest ./user-service"
            sh "docker tag ${DOCKER_USERNAME}/user-service:latest ${DOCKER_USERNAME}/user-service:${IMAGE_TAG}"
        }
    }

    stage('Build Order Service Image') {
        steps {
            sh "docker build -t ${DOCKER_USERNAME}/order-service:latest ./order-service"
            sh "docker tag ${DOCKER_USERNAME}/order-service:latest ${DOCKER_USERNAME}/order-service:${IMAGE_TAG}"
        }
    }

    stage('Push Images') {
        steps {
            withCredentials([
                usernamePassword(
                    credentialsId: 'dockerhub',
                    usernameVariable: 'DOCKER_USER',
                    passwordVariable: 'DOCKER_PASS'
                )
            ]) {

                sh 'echo $DOCKER_PASS | docker login -u $DOCKER_USER --password-stdin'

                sh "docker push ${DOCKER_USERNAME}/user-service:latest"
                sh "docker push ${DOCKER_USERNAME}/user-service:${IMAGE_TAG}"

                sh "docker push ${DOCKER_USERNAME}/order-service:latest"
                sh "docker push ${DOCKER_USERNAME}/order-service:${IMAGE_TAG}"
            }
        }
    }
}

post {
    success {
        echo 'Pipeline Success'
    }

    failure {
        echo 'Pipeline Failed'
    }
}
```

}
