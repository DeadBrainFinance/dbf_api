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
                    if (env.BRACNH_NAME == "main" || env.BRACNH_NAME == "master") {
                        sh(script: "docker compose down")
                        sh(script: "docker compose up -d --build")
                    }
                }
            }
        }

        stage("Upload to Docker Hub") {
            steps {
                sh(script: "docker login -u inkeister -p Ink@0346333767")
                sh(script: "docker tag dbf_api dbf_api:latest")
                sh(script: "docker push dbf_api:latest")
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
    def DB = "dbf_db"
    def HOST = "dbf"
    def DB_USER = "postgres"
    def DB_PASSWORD = "Ink0346333767"
    def PGDATA = "/var/lib/postgresql/data/pgdata"
    def DB_PORT = "5432"
    def API_PORT = "4000"
    def CONNECTION_STRING = "postgresql://${DB_USER}:${DB_PASSWORD}@${HOST}:${DB_PORT}/${DB}"

    sh(script: "sed -i '/^DB_DRIVER=/s\$/${DB_DRIVER}' .env")
    sh(script: "sed -i '/^DB=/s\$/${DB}' .env")
    sh(script: "sed -i '/^HOST=/s\$/${HOST}' .env")
    sh(script: "sed -i '/^DB_USER=/s\$/${DB_USER}' .env")
    sh(script: "sed -i '/^DB_PASSWORD=/s\$/${DB_PASSWORD}' .env")
    sh(script: "sed -i '/^DB_PORT=/s\$/${DB_PORT}' .env")
    sh(script: "sed -i '/^API_PORT=/s\$/${API_PORT}' .env")
    sh(script: "sed -i '/^PGDATA=/s\$/${PGDATA}' .env")
    sh(script: "sed -i '/^CONNECTION_STRING=/s\$/${CONNECTION_STRING}' .env")

    sh(script: "cat .env")
}
