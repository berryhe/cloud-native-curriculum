build:
	docker build -t berryhby/cloud-native-curriculum:week2 .
run: build
	docker run -itd -p 8080:8080 --name week2 berryhby/cloud-native-curriculum:week2
	docker exec -it week2 sh
push:
	docker push berryhby/cloud-native-curriculum:week2	
clean:
	docker stop week2
	docker rm week2
build_nsenter:
	docker run --rm jpetazzo/nsenter cat /nsenter > ./nsenter && chmod +x ./nsenter	
nsenter: build_nsenter
	PID=$(docker inspect --format {{.State.Pid}} week2)
	nsenter -t $PID -n

