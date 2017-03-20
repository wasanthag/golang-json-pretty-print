node("slave1") {
    docker.withRegistry('https://hub.docker.com', 'fa66db69-6a9d-4813-9903-22f36929608c') {
    
        git url: "https://github.com/wasanthag/golang-json-pretty-print.git", credentialsId: 'b71cbc52-dcc5-46a0-8bde-7a825ef3fa5b'
    
        sh "git rev-parse HEAD > .git/commit-id"
        def commit_id = readFile('.git/commit-id').trim()
        println commit_id
    
        stage "build"
        def app = docker.build "jenkins-test"
    
        stage "publish"
        app.push 'master'
        app.push "${commit_id}"
    }
}
