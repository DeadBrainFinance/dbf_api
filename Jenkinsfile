pipeline {
    agent {
        label "slave"
    }

    stages {
        stage("Setup environment") {
            steps {
                createEnvFile(".env.example")
                echo "File .env created"
                sh(script: "ls -a")
            }
        }

        stage("Unit test") {
            steps {
                println("Running unittest...")
            }
        }

        stage("Integration test") {
            steps {
                println("Running unittest...")
            }
        }

        stage("Docker compose") {
            steps {
                script {
                    sh(script: "docker compose down")
                    sh(script: "docker compose up -d --build")
                }
            }
        }

        stage("Upload to Docker Hub") {
            steps {
                sh(script: "docker login -u inkeister -p Ink@0346333767")
                sh(script: "docker tag dbf_api inkeister/dbf_api:latest")
                sh(script: "docker push inkeister/dbf_api:latest")
            }
        }
    }
}

def createEnvFile(sampleFile) {
    def isExisted = fileExists ".env"
    println(isExisted)
    if (isExisted) {
        sh(script: "rm .env")
        sh(script: "cp ${sampleFile} .env")
    } else {
        sh(script: "cp ${sampleFile} .env")
    }

    def DB_DRIVER = "postgres"
    def HOST = "dbf_db"
    def DB = "dbf"
    def DB_USER = "postgres"
    def DB_PASSWORD = "Ink0346333767"
    def DB_PORT = "5432"
    def API_PORT = "4000"
    def CONNECTION_STRING = "postgresql://${DB_USER}:${DB_PASSWORD}@${HOST}:${DB_PORT}/${DB}"

    sh(script: "sed -i '1 i DB_DRIVER=${DB_DRIVER}' .env")
    sh(script: "sed -i '2d' .env")
    sh(script: "sed -i '2 i DB_USER=${DB_USER}' .env")
    sh(script: "sed -i '3d' .env")
    sh(script: "sed -i '3 i DB_PASSWORD=${DB_PASSWORD}' .env")
    sh(script: "sed -i '5d' .env")
    sh(script: "sed -i '6 i DB_PORT=${DB_PORT}' .env")
    sh(script: "sed -i '7d' .env")
    sh(script: "sed -i '/^API_PORT/s/\$/${API_PORT}/' .env")
    sh(script: "sed -i '8 i CONNECTION_STRING=${CONNECTION_STRING}' .env")
    sh(script: "sed -it '9d' .env")
    sh(script: "sed -i '4d' .env")

    sh(script: "cat .env")
}
