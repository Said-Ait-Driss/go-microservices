pipeline {
    agent any

    environment {
        GO_VERSION = '1.21.0'
        GOROOT = '/usr/local/go'
        PATH = "${GOROOT}/bin:${env.PATH}"
    }

    stages {
        stage('Setup') {
            steps {
                script {
                    //Debug step to print the PATH
                    sh 'echo $PATH'
                    // Verify Go installation
                    sh 'go version'
                }
            }
        }

        stage('Checkout') {
            steps {
                git branch: 'main', url: 'https://github.com/Said-Ait-Driss/go-microservices'
            }
        }

        stage('Build and Test') {
            parallel {
                stage('product-service') {
                    steps {
                        dir('product-service') {
                            sh 'go build -o product-service main.go'
                        }
                    }
                }

                stage('review-service') {
                    steps {
                        dir('review-service') {
                            sh 'go build -o review-service main.go'
                        }
                    }
                }
                stage('store-service') {
                    steps {
                        dir('store-service') {
                            sh 'go build -o store-service main.go'
                        }
                    }
                }
            }
        }
    }

    post {
        always {
            cleanWs()
        }
        success {
            echo 'Build passed!'
        }
        failure {
            echo 'Build failed!'
        }
    }
}
