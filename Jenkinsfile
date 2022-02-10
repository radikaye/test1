pipeline {
    agent any
    stages {
        stage("Checkout code") {
            steps {
                checkout scm
            }
        }
        stage("Done") {
            steps {
                scripts {
                    echo "done"
                }
            }
        }
    }
}
