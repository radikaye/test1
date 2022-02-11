def notify(identifier, message) {

    def colorCode = '#FF0000'

    if (identifier == 'build') {
        colorCode = '#d1d1d1'
    } else if (identifier == 'push') {
        colorCode = '#b0b0b0'
    } else if (identifier == 'remove image') {
        colorCode = '#ba9527'
    } else if (identifier == 'deploy') {
        colorCode = '#cfc104'
    } else if (identifier == 'done') {
        colorCode = '#1eb300'
    } else {
        colorCode = '#e62a00'
    }

    slackSend (
        channel: "#${SLACK_CHANNEL}",
        color: colorCode,
        message: message
    )
}

pipeline {
  agent any
  tools {
        go 'basego'
  }
  environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0 
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
  }
  stages {
    stage('Pre Test') {
      steps {
        echo 'Installing dependencies'
        sh 'go version'
        sh 'go get -u golang.org/x/lint/golint'
        notify('pre test', 'pre test before build app')
      }
    }
    stage('Build App') {
      steps {
        echo 'Compiling and building'
        sh 'go build'
        notify('build', 'building application go')
      }
    }
    stage('Done') {
      steps {
        notify('done', 'awesome!!')
      }
    } 
  }
}
