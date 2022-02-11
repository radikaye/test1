pipeline {
  agent any
  stages {
    stage('build') {
      steps {
        sh 'node --version'
      }
    }
    stage('build-2') {
      steps {
        sh 'go version'
      }
    }
  }
}
