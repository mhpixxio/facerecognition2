# face_recognition_2_0
face_recognition_2_0

Look at the Powerpoint Presentation for more information.

----------- instructions -----------
build the SQL tables. In the MySQL 8.0 Command Line:
$ use facerecognition
$ source <absolute path to sql file>
build the Docker Image (Docker Dekstop must be running)
$ docker build -t facerecognition2 .
run the docker image as container
$ docker run -d -p 8080:8080 -v C:/Users/MichaelHuber/Desktop/EnvironmentForFaceRecognition/files:/app/files/ --memory 1000m --restart unless-stopped facerecognition2
use the client foun at github.com/mhpixxio/clientfacerecognition2 to represent the Front End