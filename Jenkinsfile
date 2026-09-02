pipeline {
    agent {
        label 'build-golang-stable'
    }
    options {
        skipDefaultCheckout true
    }
    stages {
        stage('Checkout') {
            steps {
                dir('/root/workspace/go/src/github.com/derekpedersen/go-utils') {
                    checkout scm
                }
            }
        }
        stage('build') {
            steps {
                dir('/root/workspace/go/src/github.com/derekpedersen/go-utils') {
                    sh 'make build'
                }
            }
        }
        stage('test') {
            steps {
                dir('/root/workspace/go/src/github.com/derekpedersen/go-utils') {
                    sh 'make test'
                }
            }
        }
    }
}
